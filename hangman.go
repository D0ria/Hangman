package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Game struct {
	Word             string     // Word composed of '_', ex: H_ll_
	ToFind           string     // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int        // Number of attempts left
	HangmanPositions [10]string // It can be the array where the positions parsed in "hangman.txt" are stored
	Found            bool       //If the letter was found, return true
}

func Reader() {
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

func RandomWord(i int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(i)
}

func RevealFewLetters(j string) {
	for n := 0; n < len(j)-1; n++ {
		rand.Seed(time.Now().Unix())
		lettre := j[rand.Intn(len(j))]
		fmt.Println(lettre)
	}
}

func HiddenWord(word string) {
	for _, letters := range word {
		if hangman.Found == false {
			letters += rune(95)
		}
	}
}

func Step(tab) 

tab = [





=========
,
         
      |  
      |  
      |  
      |  
      |  
=========
,
  +---+  
      |  
      |  
      |  
      |  
      |  
=========
,
  +---+  
  |   |  
      |  
      |  
      |  
      |  
=========
,
  +---+  
  |   |  
  O   |  
      |  
      |  
      |  
=========
,
  +---+  
  |   |  
  O   |  
  |   |  
      |  
      |  
=========
,
  +---+  
  |   |  
  O   |  
 /|   |  
      |  
      |  
=========
,
  +---+  
  |   |  
  O   |  
 /|\  |  
      |  
      |  
=========
,
  +---+  
  |   |  
  O   |  
 /|\  |  
 /    |  
      |  
=========
,
  +---+  
  |   |  
  O   |  
 /|\  |  
 / \  |  
      |  
=========
]