package game

import (
	"math/rand"

	hangman ".."
)

var InitString []string = hangman.MotRandom()
var RandomMot string = string(InitString[rand.Intn(len(InitString))])
var Lettre string
var ArrayAnswer []rune = hangman.InitTab(RandomMot)
var ArrayInit []rune = []rune(RandomMot)
var tentative int = 10
var Lettreok bool
var WordFind = false

func Game() bool {
	if InitString == nil { //si le fichier n'est pas trouvé et donc il n'y a pas de mot disponible pour le jeu, arret du jeu
		return false
	}
	Lettreok = hangman.EntrerLettre(Lettre)
	if Lettreok == true {
		hangman.Compare(Lettre, ArrayInit, ArrayAnswer) //regarde si la lettre est contenu dans le mot
		compare := hangman.Compare(Lettre, ArrayInit, ArrayAnswer)
		if compare == false { //si la lettre n'est pas contenue
			tentative--
		} else if compare == true { //si la lettre est contenue dans le mot, remplacer dans le tableau de réponse le caractère _ par la lettre mise par le joueur
			arrayRune := []rune(Lettre)
			pos := []int(nil)
			for i := 0; i < len(ArrayInit); i++ {
				if arrayRune[0] == ArrayInit[i] {
					pos = append(pos, i)
				}
				for i := 0; i < len(pos); i++ {
					ArrayAnswer[pos[i]] = ArrayInit[pos[i]]
				}
			}
		}
	} else {
		Lettreok = false
	}
	return true
}
