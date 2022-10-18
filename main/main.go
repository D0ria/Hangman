package main

import (
	"bufio"
	"fmt"
	"hangman"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // fonction random
	nbr_words := rand.Intn(85) + 1   // choisit un mot au hasard dans la liste words.txt

	file_words, _ := os.Open("words.txt")
	scanner_words := bufio.NewScanner(file_words)
	input := bufio.NewScanner(os.Stdin)

	var tab_mots []string
	for scanner_words.Scan() {
		tab_mots = append(tab_mots, scanner_words.Text()) //tableau de string créé pour le mot random choisit
	}

	vie := 10 //nombre de vies

	trouver := tab_mots[nbr_words]                     //mot à trouver
	trouver_en_rune := hangman.CastStringRune(trouver) //Cast du mot à trouver pour le print

	var ok int                          //dira si la lettre est trouvée
	var deja_trouver []rune             //deja_trouver est le mot caché par des "_" et se modifie en printant la lettre trouvée et ce, jusqu'à ce que l'on gagne
	for i := 0; i < len(trouver); i++ { //initialisation de deja_trouver
		deja_trouver = append(deja_trouver, '_')
	}

	hangman.Reponse(deja_trouver, len(trouver))           //fonction qui s'occupe de l'affichage durant le jeu
	fmt.Println("Bonne chance, vous avez 10 tentatives.") //début du jeu
	for {
		ok = 0
		fmt.Println()
		fmt.Println("Lettre choisie: ")
		input.Scan()
		input_text := input.Text()

		for j := 0; j < len(trouver); j++ { //print l'état du mot cherché
			if input_text == (string(trouver[j])) {
				ok = 1
				deja_trouver[j] = (rune(input_text[0]))
			}
		}

		if ok == 0 {
			vie--
			hangman.Etape(vie)
			if vie >= 1 { //condition tant qu'il nous reste plus d'une vie
				fmt.Print("Cette lettre n'est pas dans le mot, il vous reste exactement ", vie, " .\n")
				fmt.Println(string(deja_trouver)) //print l'état du mot cherché
			}

		} else if ok == 1 {
			hangman.Reponse(deja_trouver, len(trouver))
		}

		if vie <= 0 {
			fmt.Println("Hah cheh. I am not José Morigno.")          //message qui apparaît si la partie est perdue
			fmt.Println("La réponse était", string(trouver_en_rune)) //donne la réponse
			break
		}

		if hangman.Compare(trouver_en_rune, deja_trouver) { //si le mot est trouvé alors message Bingo
			fmt.Println("Bingo! I am José Morigno.")
			break
		}
	}

}
