package chatbot

import (
	"fmt"
	"os"

	"github.com/nsf/termbox-go"
)

func Run() {
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
				termbox.Close()
				// On récupère l'option sélectionnée
				option, err := handleSelection(selected)
				if err != nil {
					fmt.Println(handleErr(err))
				} else {
					err := handleInput(option)
					if err != nil {
						fmt.Println(handleErr(err))
					}
				}
				// On attend que l'utilisateur appuie sur une touche pour continuer
				waitForKey()
				termbox.Init()
			case termbox.KeyEsc:
				termbox.Close()
				os.Exit(0)
			}
		}
	}
}
