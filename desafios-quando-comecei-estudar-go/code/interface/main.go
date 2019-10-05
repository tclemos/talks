package example

type PeçaDeXadrez interface {
	Mover(x, y int)
}

type Cavalo struct{}

func (c Cavalo) Mover(x, y int) {}

type Peão struct{}

func (p Peão) Mover(x, y int) {}

type Torre struct{}

func (t Torre) Mover(x, y int) {}

type Game struct{}

func (g Game) MoverPeça(peça PeçaDeXadrez) {}
