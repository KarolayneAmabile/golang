package main

import (
	"fmt"
	"io"
	"os"
	"time" // para usar time.Duration, por exemplo
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct { // configura interface sleeper
	duration time.Duration       // duração da pausa
	sleep    func(time.Duration) // sleep realiza a pausa
}

// implementamos a interface sleep e definimos a sua duração
// Sleep agora é um método definido pelas variáveis da estrutura ConfigurableSleeper
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
