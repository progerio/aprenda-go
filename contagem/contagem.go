package contagem

import (
	"fmt"
	"io"
	"time"
)

const ultimaPalavra = "Vai!"
const inicioContagem = 3

func Contagem(saida io.Writer, sleeper Sleeper) {
	for i := inicioContagem; i > 0; i-- {
		sleeper.Pausa()
		fmt.Fprintln(saida, i)
	}

	sleeper.Pausa()
	fmt.Fprint(saida, ultimaPalavra)
}

type Sleeper interface {
	Sleep()
	Pausa()
}

type SleeperPadrao struct{}

func (s *SleeperPadrao) Sleep() {
	time.Sleep(1 * time.Second)
}

type SleeperConfiguravel struct {
	duracao time.Duration
	pausa   func(time.Duration)
}

func (s *SleeperConfiguravel) Pausa() {
	s.pausa(s.duracao)
}
