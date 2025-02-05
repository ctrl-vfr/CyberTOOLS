package converter

import "math/big"

// DecimalConverter gère la conversion d'une valeur décimale.
type DecimalConverter struct {
	Value int
}

// ToBinary convertit le nombre décimal en sa représentation binaire sous forme de slice de bytes.
func (c *DecimalConverter) ToBinary() []byte {
	// On convertit le nombre en binaire sous forme de chaîne de caractères.
	bigN := big.NewInt(int64(c.Value))
	return bigN.Bytes() // Retourne la représentation binaire brute
}

// ToDecimal retourne directement la valeur décimale sous forme d'int64.
func (c *DecimalConverter) ToDecimal() (int64, error) {
	return int64(c.Value), nil
}

// ToText convertit le nombre décimal en texte en l'interprétant comme un code Unicode (rune).
func (c *DecimalConverter) ToText() string {
	return string(rune(c.Value))
}
