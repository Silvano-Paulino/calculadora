package calc

type Calculadora string

func New() Calculadora {
	return "0"
}

func (c Calculadora) Ecra() string {
	return string(c)
}

func (c *Calculadora) Zero() {
	c.pressionar("0")
}

func (c *Calculadora) Um() {
	c.pressionar("1")
}

func (c *Calculadora) Dois() {
	c.pressionar("2")
}

func (c *Calculadora) Tres() {
	c.pressionar("3")
}

func (c *Calculadora) Quatro() {
	c.pressionar("4")
}

func (c *Calculadora) Cinco() {
	c.pressionar("5")
}

func (c *Calculadora) Seis() {
	c.pressionar("6")
}

func (c *Calculadora) Sete() {
	c.pressionar("7")
}

func (c *Calculadora) Oito() {
	c.pressionar("8")
}

func (c *Calculadora) Nove() {
	c.pressionar("9")
}

func (c *Calculadora) Mais() {
	c.pressionar("+")
}

func (c *Calculadora) Menos() {
	c.pressionar("-")
}

func (c *Calculadora) pressionar(s string) {
	if *c == "0" {
		*c = Calculadora(s)
	} else {
		*c = *c + Calculadora(s)
	}
}
