package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"chat_bot/bot"

	"github.com/common-nighthawk/go-figure"
	"github.com/nsf/termbox-go"
)

var options = []string{
	"Convertir texte en binaire",
	"Convertir binaire en texte",
	"Convertir binaire en décimal",
	"Convertir décimal en binaire",
	"Générer un mot de passe",
	"Quitter",
}

func main() {
	// On initialise termbox pour afficher le menu
	err := termbox.Init()
	if err != nil {
		fmt.Println("Error initializing termbox:", err)
		os.Exit(1)
	}
	// On s'assure que termbox sera fermé à la fin du programme
	defer termbox.Close()

	// On affiche le menu
	selected := 0
	for {
		// on utilise selected pour savoir quelle option est sélectionnée
		displayMenu(selected)
		// ev contient l'événement de la touche pressée
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowUp:
				if selected > 0 {
					selected--
				}
			case termbox.KeyArrowDown:
				if selected < len(options)-1 {
					selected++
				}
			case termbox.KeyEnter:
				if selected == len(options)-1 {
					termbox.Close()
					os.Exit(0)
				}
				handleSelection(selected)
			case termbox.KeyEsc:
				termbox.Close()
				os.Exit(0)
			}
		}
	}
}

// displayMenu affiche le menu avec l'option sélectionnée en surbrillance
func displayMenu(selected int) {
	//On efface l'écran à chaque fois pour éviter que le texte se superpose
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	printTitle()
	// On affiche chaque option du menu
	for i, option := range options {
		fg := termbox.ColorDefault
		bg := termbox.ColorDefault
		if i == selected {
			fg = termbox.ColorBlack
			bg = termbox.ColorWhite
		}
		printLine(2, 8+i, option, fg, bg)
	}
	// Applique les changements à l'écran
	termbox.Flush()
}

func printTitle() {
	// On utilise la bibliothèque go-figure pour afficher le titre en ASCII
	title := figure.NewFigure("Cyber TOOLS", "", true)
	titleStr := title.String()
	printLine(2, 1, titleStr, termbox.ColorCyan, termbox.ColorDefault)
}

// x = position horizontale et y = position verticale
// s = texte à afficher, fg = couleur du texte, bg = couleur de fond
func printLine(x, y int, s string, fg, bg termbox.Attribute) {
	for _, c := range s {
		// Si \n est rencontré, on passe à la ligne suivante avec y++ et on réinitialise x à 2
		if c == '\n' {
			y++
			x = 2
			continue
		}
		// c est le caractère à afficher à la position x, y
		termbox.SetCell(x, y, c, fg, bg)
		// On incrémente x pour passer au caractère suivant
		x++
	}
}

// handleSelection gère l'option sélectionnée par l'utilisateur
func handleSelection(selected int) {
	termbox.Close()
	// On utilise bufio pour lire l'entrée de l'utilisateur
	scanner := bufio.NewScanner(os.Stdin)
	// On affiche un message en fonction de l'option sélectionnée
	switch selected {
	case 0:
		fmt.Print("Entrez le texte : ")
		// On lit l'entrée de l'utilisateur
		scanner.Scan()
		// On stocke le texte dans la variable text
		text := scanner.Text()
		// On appelle la fonction Convert du package bot pour convertir le texte en binaire
		fmt.Println(bot.Convert("text:" + text))
	case 1:
		fmt.Print("Entrez le binaire : ")
		scanner.Scan()
		binary := scanner.Text()
		fmt.Println(bot.Convert("bin:" + binary))
	case 2:
		fmt.Print("Entrez le binaire : ")
		scanner.Scan()
		binary := scanner.Text()
		fmt.Println(bot.Convert("0b" + binary))
	case 3:
		fmt.Print("Entrez le nombre décimal : ")
		scanner.Scan()
		decimal := scanner.Text()
		fmt.Println(bot.Convert(decimal))
	case 4:
		fmt.Print("Entrez la longueur du mot de passe : ")
		scanner.Scan()
		// On lit la longueur du mot de passe
		lengthStr := scanner.Text()
		// On convertit la longueur en entier
		length, err := strconv.Atoi(lengthStr)
		if err != nil {
			fmt.Println("Erreur : longueur invalide.")
			break
		}
		// On appelle la fonction GeneratePassword du package bot pour générer un mot de passe
		password, err := bot.GeneratePassword(length)
		if err != nil {
			fmt.Println("Erreur :", err)
			break
		}
		fmt.Println("Mot de passe généré :", password)
	}
	// On attend que l'utilisateur appuie sur une touche pour continuer
	fmt.Println("Appuyez sur une touche pour continuer...")
	scanner.Scan()
	// On réinitialise termbox pour afficher le menu
	termbox.Init()
}
