package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
	"io/ioutil"
    "log"
)

type Game struct {
	Word             func     // Word composed of '_', ex: H_ll_
	ToFind           func     // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int        // Number of attempts left
	Found            bool       //If the letter was found, return true
}

func main() {
	var morigno Game
	morigno.Word = RandomWord()
	morigno.ToFind = LetterInWord()
	morigno.Attempts = 10
	morigno.Found = false
	morigno.HangmanPositions = "hangman.txt"
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

func HiddenWord(word string) {
	for _, letters := range word {
		if morigno.Found == false {
			letters += rune(95)
		}
	}
}

func RevealFewLetters(j string) {
	for n := 0; n < len(j)-1; n++ {
		rand.Seed(time.Now().Unix())
		lettre := j[rand.Intn(len(j))]
		fmt.Println(lettre)
	}
}

func LetterInWord(find string, letters []string) bool {
	for _, letter := range letters {
		if find == letter {
			return true
		}
	}
	morigno.Attempts -= 1
	return false
}

func Step() {
	if morigno.ToFind == 9 {
		content, err := ioutil.ReadFile("pendu1.txt")
     	if err != nil {
          	log.Fatal(err)
     	}
    	fmt.Println(string(content))
		}

	if morigno.ToFind == 8 {
		content, err := ioutil.ReadFile("pendu2.txt")
     	if err != nil {
			log.Fatal(err)
     	}
    	fmt.Println(string(content))
		}

	if morigno.ToFind == 7 {
		content, err := ioutil.ReadFile("pendu3.txt")
     	if err != nil {
			log.Fatal(err)
     	}
    	fmt.Println(string(content))
		}

	if morigno.ToFind == 6 {
		content, err := ioutil.ReadFile("pendu4.txt")
     	if err != nil {
			log.Fatal(err)
     	}
	 	fmt.Println(string(content))
		}

	if morigno.ToFind == 5 {
		content, err := ioutil.ReadFile("pendu5.txt")
     	if err != nil {
			log.Fatal(err)
    	}
    	fmt.Println(string(content))
		}

	if morigno.ToFind == 4 {
		content, err := ioutil.ReadFile("pendu6.txt")
     	if err != nil {
			log.Fatal(err)
     	}
    	fmt.Println(string(content))
		}

	if morigno.ToFind == 3 {
		content, err := ioutil.ReadFile("pendu7.txt")
     	if err != nil {
			log.Fatal(err)
     	}
    	fmt.Println(string(content))
	}

	if morigno.ToFind == 2 {
		content, err := ioutil.ReadFile("pendu8.txt")
     	if err != nil {
			log.Fatal(err)
    	}
    	fmt.Println(string(content))
	}
	if morigno.ToFind == 1 {
		content, err := ioutil.ReadFile("pendu9.txt")
     	if err != nil {
			log.Fatal(err)
     	}
    	fmt.Println(string(content))
		}
	if morigno.ToFind == 0 {
		content, err := ioutil.ReadFile("pendu10.txt")
     	if err != nil {
			log.Fatal(err)
    	 }
    	fmt.Println(string(content))
		}
}

func EndOfGame() {
	if morigno.ToFind == RandomWord {
		return "Gagné"
	} if morigno.Attempts == 0 {
		return "Perdu"
	}
}
