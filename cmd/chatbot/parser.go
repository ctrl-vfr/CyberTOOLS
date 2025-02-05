package chatbot

// Parser pourrait être implémenté comme un module séparé dans un package internal/parser (car pkg est plus pour les packages publics).

import (
	"errors"
	"math/big"
	"strconv"
	"strings"
)

// Parser définit une interface pour les différents types de parseurs.
type Parser interface {
	Parse(string) (interface{}, error)
}

type ParseText struct{}
type ParseBinary struct{}
type ParseUint struct{}
type ParseInt struct{}

// Parse convertit une chaîne de caractères en une valeur de type interface{}.
func (p *ParseInt) Parse(s string) (interface{}, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("invalid integer")
	}
	return i, nil
}

// Parse convertit une chaîne de caractères en une valeur de type interface{}.
func (p *ParseText) Parse(s string) (interface{}, error) {
	return s, nil
}

// Parse convertit une chaîne de caractères en une valeur de type interface{}.
func (p *ParseBinary) Parse(s string) (interface{}, error) {
	str := strings.ReplaceAll(s, " ", "")
	for _, c := range str {
		if c != '0' && c != '1' {
			return nil, errors.New("invalid binary")
		}
	}
	n := new(big.Int)
	n.SetString(str, 2) // Base 2 (binaire)
	return n.Bytes(), nil
}

// Parse convertit une chaîne de caractères en une valeur de type interface{}.
func (p *ParseUint) Parse(s string) (interface{}, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("invalid integer")
	}
	return uint(i), nil
}
