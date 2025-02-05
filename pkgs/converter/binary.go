package converter

// BinaryConverter gère la conversion d'une valeur binaire représentée sous forme de slice de bytes.
type BinaryConverter struct {
	Value []byte
}

// ToBinary retourne directement la valeur binaire telle quelle.
func (c *BinaryConverter) ToBinary() []byte {
	return c.Value
}

// ToDecimal convertit la valeur binaire en un nombre décimal.
func (c *BinaryConverter) ToDecimal() (int64, error) {
	var n int64
	for _, bit := range c.Value {
		n = (n << 1) | int64(bit) // Décale de 1 bit à gauche et ajoute le bit actuel
	}
	return n, nil
}

// ToText convertit la valeur binaire en texte (Non utilisé dans le programme parent).
func (c *BinaryConverter) ToText() string {
	return string(c.Value)
}
