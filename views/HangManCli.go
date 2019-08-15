package views

import (
	"bufio"
	"fmt"
	"os"
)
const QUIT_CHAR = '0'
/**
HangManCli is struct that has methods on it to sent and get data from the user.
 */
type HangManCli struct {
	Reader *bufio.Reader
}

/**
This is a constructor for the HangMAnCli. It creates a reader to read data from the console.
 */
func (hmc HangManCli) NewHangManCli() HangManCli {
	//hmc:= HangManCli{
	//	Reader: bufio.NewReader(os.Stdin),
	//}
	hmc.Reader=bufio.NewReader(os.Stdin)

	return hmc
}

/**
This is a welcome screen that prints to the console.
 */
func (hmc HangManCli) WelcomeScreen()  {
	fmt.Println("##################################################################\n\n")
	fmt.Println("#################### WELCOME TO HANGMAN ##########################\n\n")
}

/**
This uses ReadString from the Reader of the struct.
 */
func (hmc HangManCli) GetWordToGuess() string {

	fmt.Print("Enter the word you like to be guessed:")
	var wordToGuess string

	wordToGuess, _ = hmc.Reader.ReadString('\n') // reading from the command line

	fmt.Println("you have chosen:", wordToGuess)

	return wordToGuess
}

func (hmc HangManCli) GetLetterGuessed() byte {

	var letterGuessed byte
	fmt.Print("Guess a character:(0 for quit):")
	letterGuessed, _= hmc.Reader.ReadByte()

	fmt.Printf("\nyou have chose: %s %d\n", letterGuessed)

	return letterGuessed

}
