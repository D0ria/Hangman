package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func PrintMyTabRune(tab_rune []rune, taille_tab int) { //print la decouverte du mot à chaque tour
	for i := 0; i < taille_tab; i++ {
		fmt.Print(string(tab_rune[i]))
	}
	fmt.Print("\n")
}

func StringToSliceRune(word_to_found string) []rune { //transforme un string en un tableau de rune
	var result []rune
	for i := 0; i < len(word_to_found); i++ {
		result = append(result, (rune(word_to_found[i])))
	}
	return result
}

func Equal(a, b []rune) bool { //fonction compare le tableau de runes(condition de victoire)
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func PrintHangMan(life int) {
	//DEBUT CREATION FLUX BOURREAU
	file_bourreau_1, _ := os.Open("bourreau_1.txt")         //ouverture bourreau_1.txt
	file_bourreau_2, _ := os.Open("bourreau_2.txt")         //ouverture bourreau_2.txt
	file_bourreau_3, _ := os.Open("bourreau_3.txt")         //ouverture bourreau_3.txt
	file_bourreau_4, _ := os.Open("bourreau_4.txt")         //ouverture bourreau_4.txt
	file_bourreau_5, _ := os.Open("bourreau_5.txt")         //ouverture bourreau_5.txt
	file_bourreau_6, _ := os.Open("bourreau_6.txt")         //ouverture bourreau_6.txt
	file_bourreau_7, _ := os.Open("bourreau_7.txt")         //ouverture bourreau_7.txt
	file_bourreau_8, _ := os.Open("bourreau_8.txt")         //ouverture bourreau_8.txt
	file_bourreau_9, _ := os.Open("bourreau_9.txt")         //ouverture bourreau_9.txt
	scanner_bourreau_1 := bufio.NewScanner(file_bourreau_1) //flux bourreau_1.txt
	scanner_bourreau_2 := bufio.NewScanner(file_bourreau_2) //flux bourreau_2.txt
	scanner_bourreau_3 := bufio.NewScanner(file_bourreau_3) //flux bourreau_3.txt
	scanner_bourreau_4 := bufio.NewScanner(file_bourreau_4) //flux bourreau_4.txt
	scanner_bourreau_5 := bufio.NewScanner(file_bourreau_5) //flux bourreau_5.txt
	scanner_bourreau_6 := bufio.NewScanner(file_bourreau_6) //flux bourreau_6.txt
	scanner_bourreau_7 := bufio.NewScanner(file_bourreau_7) //flux bourreau_7.txt
	scanner_bourreau_8 := bufio.NewScanner(file_bourreau_8) //flux bourreau_8.txt
	scanner_bourreau_9 := bufio.NewScanner(file_bourreau_9) //flux bourreau_9.txt
	//FIN CREATION FLUX BOURREAU

	//DEBUT PRINT BOURREAU SELON NOMBRE DE VIE
	if life == 9 {
		for scanner_bourreau_1.Scan() { //print 9 lifes
			fmt.Println(scanner_bourreau_1.Text())
		}
	}
	if life == 8 {
		for scanner_bourreau_2.Scan() { //print 8 lifes
			fmt.Println(scanner_bourreau_2.Text())
		}
	}
	if life == 7 {
		for scanner_bourreau_3.Scan() { //print 7 lifes
			fmt.Println(scanner_bourreau_3.Text())
		}
	}
	if life == 6 {
		for scanner_bourreau_4.Scan() { //print 6 lifes
			fmt.Println(scanner_bourreau_4.Text())
		}
	}
	if life == 5 {
		for scanner_bourreau_5.Scan() { //print 5 lifes
			fmt.Println(scanner_bourreau_5.Text())
		}
	}
	if life == 4 {
		for scanner_bourreau_6.Scan() { //print 4 lifes
			fmt.Println(scanner_bourreau_6.Text())
		}
	}
	if life == 3 {
		for scanner_bourreau_7.Scan() { //print 3 lifes
			fmt.Println(scanner_bourreau_7.Text())
		}
	}
	if life == 2 {
		for scanner_bourreau_8.Scan() { //print 2 lifes
			fmt.Println(scanner_bourreau_8.Text())
		}
	}
	if life == 1 {
		for scanner_bourreau_9.Scan() { //print 9 lifes
			fmt.Println(scanner_bourreau_9.Text())
		}
	}
	//FIN PRINT BOURREAU SELON VIE
}

func Rejouer() bool { //fonction pour rejouer ou non
	var reponse string
	fmt.Println("Voulez vous rejouer ? O pour rejouer , N pour arreter")
	fmt.Scan(&reponse)
	if reponse == "O" {
		return true
	} else if reponse == "N" {
		return false
	} else {
		fmt.Println("Veuillez rentrer soit O pour rejouer, soit N pour arreter")
		fmt.Scan(&reponse)
	}
	return false
}

func main() {
	rand.Seed(time.Now().UnixNano()) //seed
	nbr_words := rand.Intn(85) + 1   //génération d'un nombre aléatoire

	file_words, _ := os.Open("words.txt")         //ouverture words.txt
	scanner_words := bufio.NewScanner(file_words) //flux words.txt

	input := bufio.NewScanner(os.Stdin) //flux input

	var tab_words []string //tableau de mot

	for scanner_words.Scan() { //remplire le tableau de mot d'après words.txt
		tab_words = append(tab_words, scanner_words.Text())
	}
	// fmt.Println(tab_words[nbr_words]) ============================================================== imprime le mot du jeu

	life := 10 //nombre de vie

	to_found := tab_words[nbr_words] //mot à trouver

	//var to_found_RuneVersion []rune                    //tableau de rune futur miroir de to_found
	to_found_RuneVersion := StringToSliceRune(to_found) //to_found_RuneVersion devient miroir de to_found

	//fmt.Println(to_found_RuneVersion) ================================TEST

	var correct int //bollean si la lettre est trouver ou non

	var founded []rune                   //founded est le mot version "____" et ce print à caque lettre trouver, devient le mot à trouver jusqu'a victoire
	for i := 0; i < len(to_found); i++ { //initialisation de founded pour qu'il devient un miroir de "_"
		founded = append(founded, '_')
	}

	PrintMyTabRune(founded, len(to_found))                //appel de PrintMyTabRune pour print founded + taille du mot à trouver(besoins technique)
	fmt.Println("Bonne chance, vous avez 10 tentatives.") //début du jeu
	for {
		correct = 0
		fmt.Println()
		fmt.Println("Lettre choisie: ")
		input.Scan()
		input_text := input.Text()

		for j := 0; j < len(to_found); j++ { //lettre detector, change egalement le mot révélé selon les lettres trouver
			if input_text == (string(to_found[j])) {
				correct = 1
				founded[j] = (rune(input_text[0]))
			}
		}

		if correct == 0 {
			life--
			PrintHangMan(life)
			if life >= 1 { //pour évité de print après le mot
				fmt.Print("Cette lettre n'est pas dans le mot, il vous reste exactement ", life, " .")
			}

		} else if correct == 1 {
			PrintMyTabRune(founded, len(to_found)) //appel de PrintMyTabRune pour print le mot révélé à chaque tour gagné
		}

		if life <= 0 {
			fmt.Println("Hah cheh. I am not José Morigno.")
			break
		}

		if Equal(to_found_RuneVersion, founded) {
			fmt.Println("Bingo! I am José Morigno.")
			break
		}
	}

}
