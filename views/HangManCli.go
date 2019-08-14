package views

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type HangManCli struct {
	reader Reader
}

func newHangManCli()  HangManCli{
	r := bufio.NewReader(os.Stdin)
	hmc := HangManCli{
		reader: r,
	}

	return hmc
}

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

func getLetterGuessed() byte {

	var letterGuessed byte
	reader := bytes.NewReader(os.Stdin)
	letterGuessed, _=reader.ReadByte()

	fmt.Println("you have chose:", letterGuessed)

	return letterGuessed

}
