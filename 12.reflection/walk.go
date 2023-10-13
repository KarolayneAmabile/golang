package main

import (
	"fmt"
	"reflect"
)

// leva uma interfase qualquer e passa todas os valores strings dela em uma função
func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walkValue(v)
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}

func imprimir(valor string) {
	fmt.Printf("O valor armazenado é: %v \n", valor)
}

func main() {

	type Profile struct {
		Age  int
		City string
	}

	type Person struct {
		Name    string
		Profile Profile
	}

	karol := &Person{
		Name: "Karol",
		Profile: Profile{
			Age:  20,
			City: "Uruaçu",
		},
	}

	house := []Profile{
		{
			Age:  1,
			City: "Maracanã",
		},
		{
			Age:  1,
			City: "Universitária",
		},
	}

	bike := [2]Profile{
		{
			Age:  20,
			City: "Uruaçu",
		},
		{
			Age:  7,
			City: "Unknown",
		},
	}

	pets := make(map[string]string)
	pets["hanzel"] = "beautiful fat cat"
	pets["astrid"] = "stupid ass cat"

	walk(karol, imprimir)
	walk(house, imprimir)
	walk(bike, imprimir)
	walk(pets, imprimir)
}
