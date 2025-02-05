package converter

import (
	"fmt"
)

// Converter définit une interface commune pour convertir une valeur entre différentes représentations.
type Converter interface {
	// ToBinary retourne la représentation binaire de la valeur sous forme de slice de bytes.
	ToBinary() []byte
	// ToDecimal retourne la représentation décimale de la valeur.
	ToDecimal() (int64, error)
	// ToText retourne la représentation textuelle de la valeur.
	ToText() string
	// ToHex retourne la représentation hexadécimale de la valeur ? (nouvelle méthode)
}

// New est une fonction "factory" qui retourne un Converter adapté selon le type de la valeur passée.
func New(value interface{}) (Converter, error) {
	if v, ok := value.(string); ok {
		return &TextConverter{Value: v}, nil
	}
	if v, ok := value.(int); ok {
		return &DecimalConverter{Value: v}, nil
	}
	if v, ok := value.([]byte); ok {
		return &BinaryConverter{Value: v}, nil
	}
	return nil, fmt.Errorf("type invalide")
}
