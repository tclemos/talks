package biblioteca

// Algo é uma estrutura fornecida pela biblioteca
type Algo struct {
	visivel bool
}

// CriarAlgo é uma função fornecida pela biblioteca
// para criar uma instancia de Algo
func CriarAlgo(visivel bool) *Algo {
	return &Algo{visivel: visivel}
}

// Esconder é uma função fornecida pela estrutura Algo
// para que Algo possa se marcar como não visível
func (a *Algo) Esconder() {
	a.visivel = false
}
