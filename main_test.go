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
