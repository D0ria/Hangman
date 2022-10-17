package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

func Reponse(tabrune []rune, tailletab int) { //fonction qui convertie la réponse (originalement en string) en runes
	for i := 0; i < tailletab; i++ {
		fmt.Print(string(tabrune[i]))
	}
	fmt.Print("\n")
}

func CastStringRune(motcache string) []rune { //fonction qui transforme un string en un tableau de runes pour append chaque lettre dans ce tableau
	var result []rune
	for i := 0; i < len(motcache); i++ {
		result = append(result, (rune(motcache[i])))
	}
	return result
}

func Compare(lettre, reponsetab []rune) bool { //fonction qui compare la lettre entrée par rapport aux lettres dans le mot
	if len(lettre) != len(reponsetab) {
		return false
	}
	for i := 0; i < len(lettre); i++ {
		if lettre[i] != reponsetab[i] {
			return false
		}
	}
	return true
}

func Etape(vie int) { //fonction qui renvoie chaque étape du pendu si la vie décroît
	if vie == 9 {
		content, err := ioutil.ReadFile("pendu1.txt")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(content))
	}

	if vie == 8 {
		content, err := ioutil.ReadFile("pendu2.txt")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(content))
	}

	if vie == 7 {
		content, err := ioutil.ReadFile("pendu3.txt")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(content))
	}

	if vie == 6 {
		content, err := ioutil.ReadFile("pendu4.txt")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(content))
	}

	if vie == 5 {
		content, err := ioutil.ReadFile("pendu5.txt")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(content))
	}

	if vie == 4 {
		content, err := ioutil.ReadFile("pendu6.txt")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(content))
	}

	if vie == 3 {
		content, err := ioutil.ReadFile("pendu7.txt")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(content))
	}

	if vie == 2 {
		content, err := ioutil.ReadFile("pendu8.txt")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(content))
	}

	if vie == 1 {
		content, err := ioutil.ReadFile("pendu9.txt")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(content))
	}

	if vie == 0 {
		content, err := ioutil.ReadFile("pendu10.txt")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(content))
	}
}

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

	trouver := tab_mots[nbr_words]             //mot à trouver
	trouver_en_rune := CastStringRune(trouver) //Cast du mot à trouver pour le print

	var ok int                          //dira si la lettre est trouvée
	var deja_trouver []rune             //deja_trouver est le mot caché par des "_" et se modifie en printant la lettre trouvée et ce, jusqu'à ce que l'on gagne
	for i := 0; i < len(trouver); i++ { //initialisation de deja_trouver
		deja_trouver = append(deja_trouver, '_')
	}

	Reponse(deja_trouver, len(trouver))                   //fonction qui s'occupe de l'affichage durant le jeu
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
			Etape(vie)
			if vie >= 1 { //condition tant qu'il nous reste plus d'une vie
				fmt.Print("Cette lettre n'est pas dans le mot, il vous reste exactement ", vie, " .\n")
				fmt.Println(string(deja_trouver)) //print l'état du mot cherché
			}

		} else if ok == 1 {
			Reponse(deja_trouver, len(trouver))
		}

		if vie <= 0 {
			fmt.Println("Hah cheh. I am not José Morigno.")          //message qui apparaît si la partie est perdue
			fmt.Println("La réponse était", string(trouver_en_rune)) //donne la réponse
			break
		}

		if Compare(trouver_en_rune, deja_trouver) { //si le mot est trouvé alors message Bingo
			fmt.Println("Bingo! I am José Morigno.")
			break
		}
	}

}
