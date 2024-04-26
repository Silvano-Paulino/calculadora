package calc_test

import (
	calc "calculadora"
	"testing"
)

func verifica(t *testing.T, esperado, actual string) {
	if esperado != actual {
		t.Logf("%s != %s", esperado, actual)
		t.Fail()
	}
}

func TestCria_0(t *testing.T) {
	// Arrange
	// Act
	c := calc.New()

	// Assert
	verifica(t, "0", c.Ecra())
}
func TestPressiona5DepoisDeLigar_5(t *testing.T) {
	// Arrange
	c := calc.New()

	// Act
	c.Cinco()

	// Assert
	verifica(t, "5", c.Ecra())

}

func TestPressionaMaisDepoisDe7_7Mais(t *testing.T) {
	// Arrange
	c := calc.New()
	c.Sete()

	// Act
	c.Mais()

	//Assert
	verifica(t, "7+", c.Ecra())

}

func TestPressionaMaisDepoisDe2_9_7__297Mais(t *testing.T) {
	// Arrange
	c := calc.New()
	c.Dois()
	c.Nove()
	c.Sete()

	// Act
	c.Mais()

	// Assert
	verifica(t, "297+", c.Ecra())
}

func TestPressionaMenosDepoisDe1_1Menos(t *testing.T) {
	// Arrange
	c := calc.New()
	c.Um()

	// Act
	c.Menos()

	// Assert
	verifica(t, "1-", c.Ecra())
}

func TestPressiona1AntesDeMenosE2DepoisDeMenos_1Menos2(t *testing.T) {
	// Arrange
	c := calc.New()
	c.Um()
	c.Menos()

	// Act
	c.Dois()

	// Assert
	verifica(t, "1-2", c.Ecra())

}

func TestPressionaMaisDepoisDe1_0__10Mais(t *testing.T) {
	// Arrange
	c := calc.New()
	c.Um()
	c.Zero()

	// Act
	c.Mais()

	// Assert
	verifica(t, "10+", c.Ecra())
}

func TestPressiona34DepoisDeMenos_Menos34(t *testing.T) {
	// Arrange
	c := calc.New()
	c.Tres()
	c.Quatro()

	// Act
	c.Menos()

	// Assert
	verifica(t, "-34", c.Ecra())
}

func TestPressiona68AntesDeMais_68Mais(t *testing.T) {
	// Arrange
	c := calc.New()
	c.Seis()
	c.Oito()

	// Act
	c.Mais()

	// Assert
	verifica(t, "68+", c.Ecra())
}
