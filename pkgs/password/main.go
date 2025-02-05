package password

import (
	"crypto/rand"
	"math/big"
)

const passwordChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+"

func New(length uint) (string, error) {
	// Créer un tableau de bytes de la longueur spécifiée
	password := make([]byte, length)
	// Itérer sur chaque élément du tableau
	for i := range password {
		// Générer un index aléatoire pour choisir un caractère dans la chaîne de caractères
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(passwordChars))))
		if err != nil {
			return "", err
		}
		// Affecter le caractère à la position i du tableau
		password[i] = passwordChars[randomIndex.Int64()]
	}

	return string(password), nil
}
