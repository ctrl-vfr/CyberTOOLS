package chatbot

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ctrl-vfr/CyberTOOLS/pkgs/converter" // Ces modules peuvent être importés pour d'autre projets via exemple go get github.com/ctrl-vfr/CyberTOOLS/pkgs/converter
	"github.com/ctrl-vfr/CyberTOOLS/pkgs/password"
)

// handleSelection retourne l'option sélectionnée par l'utilisateur
func handleSelection(selected int) (Option, error) {
	if selected < 0 || selected >= len(options) {
		return Option{}, fmt.Errorf("Option invalide")
	}
	return options[selected], nil
}

// handleSelection gère l'option sélectionnée par l'utilisateur
func handleInput(option Option) error {
	// On utilise bufio pour lire l'entrée de l'utilisateur
	scanner := bufio.NewScanner(os.Stdin)
	// On affiche un message en fonction de l'option sélectionnée
	fmt.Println(option.Prompt)
	scanner.Scan()
	text := scanner.Text()

	switch option.Package {
	case CONVERTER:
		// On valide l'entrée de l'utilisateur en fonction de l'option sélectionnée
		val, err := option.Parser.Parse(text)
		if err != nil {
			return err
		}
		// On initialise la variable conv pour stocker le convertisseur
		conv, err := converter.New(val)
		if err != nil {
			return err
		}
		// On utilise un switch pour appeler la fonction appropriée en fonction de l'option sélectionnée
		switch option.Call {
		case TO_BINARY:
			b := conv.ToBinary()
			fmt.Println(option.Message)
			for _, v := range b {
				fmt.Printf("%08b ", v)
			}
			fmt.Println()
		case TO_TEXT:
			fmt.Println(option.Message)
			fmt.Println(conv.ToText())
		case TO_DECIMAL:
			fmt.Println(option.Message)
			d, err := conv.ToDecimal()
			if err != nil {
				return err
			}
			fmt.Println(d)
			return nil
		}

	case PASSWORD:
		switch option.Call {
		case GENERATE_PASSWORD:
			val, err := option.Parser.Parse(text)
			if err != nil {
				fmt.Println("Erreur :", err)
			} else {
				p, err := password.New(val.(uint))
				if err != nil {
					return err
				}
				fmt.Println(option.Message)
				fmt.Println(p)
			}
		default:
			return fmt.Errorf("invalid option")
		}

	case SYSTEM:
		switch option.Call {
		case QUIT:
			os.Exit(0)
		default:
			return fmt.Errorf("invalid option")
		}

	default:
		return fmt.Errorf("invalid package")
	}

	return nil
}

func handleErr(err error) string {
	switch err.Error() {
	case "invalid integer":
		return "❌ Entrez un nombre entier valide"
	case "invalid binary":
		return "❌ Entrez un nombre entier uniquement composé de 0 et 1"
	default:
		return fmt.Sprintf("❌ %s", err)
	}
}
