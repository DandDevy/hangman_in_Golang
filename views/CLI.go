package views

import (
	"bufio"
	"fmt"
	"os"
)

func welcomeScreen()  {
	fmt.Println("#################### WELCOME TO HANGMAN ##########################")
}

func getWordToGuess() string {
	var wordToGuess string
	reader := bufio.NewReader(os.Stdin) //creating a reader
	wordToGuess, _ = reader.ReadString('\n') // reading from the command line

	fmt.Println("you have chosen:", wordToGuess)

	return wordToGuess
}

func getLetterGuessed() chan {

}
