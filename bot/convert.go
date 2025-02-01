package bot

import (
	"fmt"
	"strconv"
	"strings"
)

// Conversion du texte en binaire
// Il Prend une chaîne de caractères en entrée et convertit chaque caractère en sa représentation binaire
// Il itère sur chaque caractère de la chaîne et utilise la fonction fmt.Sprintf pour convertir chaque caractère en sa représentation binaire
func TextToBinary(text string) string {
	var binaryStrings []string
	for _, char := range text {
		// Convertit chaque caractères en sa représentation binaire sur 8 bits
		binaryStrings = append(binaryStrings, fmt.Sprintf("%08b", char))
	}
	// Pour terminer il faut joindre les chaînes binaires ensemble avec un espace
	return strings.Join(binaryStrings, " ")
}

// Conversion binaire en texte
func BinaryToText(binary string) (string, error) {
	// Diviser la chaîne binaire en morceaux de 8 bits
	bits := strings.Split(binary, " ")
	// Initialiser une chaîne vide pour stocker le résultat
	var result string
	// Itérer sur chaque morceau de 8 bits
	for _, b := range bits {
		// Ici, on utilise parseInt pour convertir une chaîne binaire en un entier
		// b = la chaîne binaire à convertir
		// 2 = la base du nombre à convertir (binaire)
		// 64 = la taille de l'entier de sortie (int64)
		charCode, err := strconv.ParseInt(b, 2, 64)
		if err != nil {
			return "", fmt.Errorf("invalid binary sequence: %s", b)
		}
		// charcode est l'entier représentant le caractère ASCII
		// rune convertit l'entier en un caractère Unicode
		// Envolopper les deux avec string pour obtenir une chaîne de caractères
		// result stocke le résultat final
		result += string(rune(charCode))
	}
	return result, nil
}

// Convertir un nombre décimal en binaire
func ToBinary(decimal int) string {
	// strconv.FormatInt convertit un entier en une chaîne de caractères
	// 2 = la base du nombre à convertir (binaire)
	// int64(decimal) = l'entier à convertir avec une conversion de type int64
	return strconv.FormatInt(int64(decimal), 2)
}

// Convertir un nombre binaire en décimal
func ToDecimal(binary string) (int64, error) {
	// On fait l'inverse mais on utilise ParseInt pour convertir une chaîne binaire en un entier
	return strconv.ParseInt(binary, 2, 64)
}

// Les commandes pour convertir le texte en binaire, le binaire en texte et les nombres décimaux en binaire
func Convert(input string) string {
	// Ici on vérifie si l'entrée commence par "text:" pour convertir le texte en binaire
	if strings.HasPrefix(input, "text:") {
		// On utilise TrimPrefix pour supprimer le préfixe "text:"
		text := strings.TrimPrefix(input, "text:")
		if text == "" {
			return "❌ Erreur : texte vide. Utilisation : text [votre texte]"
		}
		// On appelle la fonction TextToBinary pour convertir le texte en binaire
		// On utilise %s pour insérer la chaîne de caractères dans le message de retour
		return fmt.Sprintf("✅ Texte en binaire : %s", TextToBinary(text))
	}

	// Si l'entrée commence par "bin:" alors ok pour convertir le binaire en texte
	if strings.HasPrefix(input, "bin:") {
		binary := strings.TrimPrefix(input, "bin:")
		if binary == "" {
			return "❌ Erreur : binaire vide. Utilisation : bin [votre binaire]"
		}
		// On appelle BinaryToText pour convertir le binaire en texte
		text, err := BinaryToText(binary)
		// On vérifie si une erreur s'est produite lors de la conversion on utilise %v pour insérer l'erreur dans le message de retour
		if err != nil {
			return fmt.Sprintf("❌ Erreur : %v", err)
		}
		return fmt.Sprintf("✅ Binaire en texte : %s", text)
	}

	// Ici on vérifie si l'entrée commence par "0b" ou si elle est binaire
	if strings.HasPrefix(input, "0b") || isBinary(input) {
		input = strings.TrimPrefix(input, "0b")
		// On appelle ToDecimal pour convertir le binaire en décimal
		decimal, err := ToDecimal(input)
		if err != nil {
			return "❌ Erreur : format binaire invalide."
		}
		// On utilise %d pour insérer l'entier dans le message de retour
		return fmt.Sprintf("✅ %s en décimal : %d", input, decimal)
	}

	// Si l'entrée n'est pas binaire, on suppose qu'elle est décimale
	// On utilise Atoi pour convertir l'entrée en un entier
	decimal, err := strconv.Atoi(input)
	if err != nil {
		return "❌ Erreur : nombre invalide. Entrez un nombre décimal, binaire, texte ou commande valide."
	}
	// On utilise %d pour insérer l'entier dans le message de retour et ToBinary pour convertir le décimal en binaire
	return fmt.Sprintf("✅ %d en binaire : 0b%s", decimal, ToBinary(decimal))
}

// Vérifier si une chaîne est binaire
// On itère sur chaque caractère de la chaîne et on vérifie si le caractère est '0' ou '1'
// On utilise bool car la fonction retourne vrai ou faux
// s = la chaîne à vérifier string = le type de la chaîne
func isBinary(s string) bool {
	// On itère sur chaque caractère c de la chaîne s
	// la variable _ est utilisée pour ignorer l'index qui est la position du caractère dans la chaîne
	for _, c := range s {
		// On vérifie si le caractère c est différent de '0' et '1'
		if c != '0' && c != '1' {
			return false
		}
	}
	return true
}
