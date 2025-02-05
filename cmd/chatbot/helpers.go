package chatbot

import (
	"bufio"
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/nsf/termbox-go"
)

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
		printLine(2, 8+i, option.Name, fg, bg)
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

// waitForKey attend que l'utilisateur appuie sur une touche pour continuer
func waitForKey() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Appuyez sur une touche pour continuer...")
	scanner.Scan()
}
