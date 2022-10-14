package hangman

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func Lecture() {
	f, err := os.Open("words.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func MotRandom() []string { //fonction qui choisi un mot aléatoire dans une des 3 listes
	var tableauMot []string
	min := 1
	max := 3
	rand.Seed(time.Now().UnixNano())
	whichWord := (rand.Intn(max-min+1) + min) //choisi un nombre aléatoire entre 1 et 3( pour les listes de mots )
	words := "hangman-classic/words" + strconv.Itoa(whichWord) + ".txt"
	f, err := os.Open(words)

	if err != nil { //gestion de l'erreur
		fmt.Println("Fichier introuvable")
		return tableauMot
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		rand.Seed(time.Now().UnixNano())
		tableauMot = append(tableauMot, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Fichier introuvable")
	}
	return tableauMot
}

func InitTab(motRandom string) []rune { // créer le tableau de réponse avec des lettres déjà mise
	tabinit := []rune(motRandom)
	tabreponse := []rune(nil)
	for i := 0; i < len(tabinit); i++ {
		tabreponse = append(tabreponse, '_')
	}
	indice := (len(tabinit) / 2) - 1
	for i := 0; i < indice; i++ {
		rand.Seed(time.Now().UnixNano())
		position := rand.Intn(len(tabinit))
		letter := tabinit[position]
		tabreponse[position] = letter
		for j := 0; j < len(tabinit); j++ {
			if letter == tabinit[j] {
				tabreponse[j] = letter
			}
		}
	}
	return tabreponse
}

func Compare(Lettre string, InitTab []rune, Reponsetab []rune) bool {
	TabRune := []rune(Lettre)
	for i := 0; i < len(InitTab); i++ {
		if TabRune[0] == InitTab[i] {
			return true
		}
	}
	return false
}

func EntrerLettre(Lettre string) bool { // fonction qui demande une lettre au joueur et la vérifie
	SpecialCar := []string{"é", "è", "ê", "à", "â", "-", "ù", "û", "ç"}
	for i := 0; i < len(SpecialCar); i++ {
		if Lettre == SpecialCar[i] {
			return true
		}
	}
	TabRuneLettre := []rune(Lettre)
	if TabRuneLettre[0] >= 'a' && TabRuneLettre[0] <= 'z' {
		TabRuneLettre[0] -= 32
	}
	if TabRuneLettre[0] >= 'A' && TabRuneLettre[0] <= 'Z' {
		return true
	} else {
		return false
	}
}

func Etape(tentative int) { //fonction qui affiche la position du pendu par rapport au nombre de vie restante
	var tabMorigno []byte
	vie := 10
	position := vie - tentative
	content, err := ioutil.ReadFile("hangman.txt")

	if err != nil {
		fmt.Println("[fichier non trouvé]")
	} else {
		for i := 0; i < 71; i++ {
			tabMorigno = append(tabMorigno, content[i+(71*position)])

		}
		morigno := (string(tabMorigno))
		fmt.Println(morigno)

	}
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
