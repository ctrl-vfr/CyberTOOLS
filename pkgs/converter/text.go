package converter

import "strconv"

// TextConverter gère la conversion d'une valeur textuelle.
type TextConverter struct {
	Value string
}

// ToBinary convertit le texte en slice de bytes.
func (c *TextConverter) ToBinary() []byte {
	return []byte(c.Value)
}

// ToDecimal tente de convertir le texte en nombre décimal.
func (c *TextConverter) ToDecimal() (int64, error) {
	return strconv.ParseInt(c.Value, 10, 64)
}

// ToText retourne simplement le texte initial.
func (c *TextConverter) ToText() string {
	return c.Value
}
