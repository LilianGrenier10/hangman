package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type HangmanData struct {
	mot_brouillé  string
	mot_a_trouver string
	nombre_essais int
	coordonnées   []int
}

func hangmanreader(ent1 int, ent2 int) {
	readFile, err := os.Open("hangman.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lignes []string
	for fileScanner.Scan() {
		lignes = append(lignes, fileScanner.Text())
	}
	readFile.Close()
	for i := ent1; i <= ent2; i++ {
		fmt.Printf(lignes[i] + "\n")
	}
}

func main() {
	fmt.Scan()
	pendu := HangmanData{}
	var a []int
	fautes := 0
	for i := 0; i <= 9; i++ {
		if i == 0 {
			a = append(a, 0)
			a = append(a, 8)
		} else {
			a = append(a, i*8)
			a = append(a, i*8+7)
		}
	}
	pendu.coordonnées = a
	if fautes == 0 {
		hangmanreader(pendu.coordonnées[0], pendu.coordonnées[1])
	} else {
		hangmanreader(pendu.coordonnées[fautes*2], pendu.coordonnées[fautes*2+1])
	}

}
