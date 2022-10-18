package hangman

import (
	"fmt"
	"io/ioutil"
	"log"
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
