package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

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

func main() {
	var solution string
	// Getting random character
	mot_a_trouver := solution[rand.Intn(len(solution))]

	// Display the character
	fmt.Println(string(mot_a_trouver))
}
