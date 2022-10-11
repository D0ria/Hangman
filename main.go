package hangman

func main() {

	var morigno Game
	morigno.Word = RandomWord
	morigno.ToFind = HiddenWord
	morigno.Attempts = 3
	morigno.HangmanPositions = Step()
	morigno.Found = false
}
