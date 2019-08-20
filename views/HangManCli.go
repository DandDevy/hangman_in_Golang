package views

import (
	"bufio"
	"fmt"
	"os"
)
const QUIT_CHAR = '0'
/**
HangManCLI is struct that has methods on it to sent and get data from the user.
 */
type HangManCLI struct {
	Reader *bufio.Reader
}

/**
Getter for reader of type *bufio.Reader.
 */
func (HMC HangManCLI) getReader()  *bufio.Reader{
	return HMC.Reader
}

/**
Setter for reader of type *bufio.Reader.
*/
func (HMC HangManCLI) setReader(pointerBufioReader *bufio.Reader)  {
	HMC.Reader = pointerBufioReader
}

/**
This is a constructor for the HangManCli. It creates a reader to read data from the console.
 */
func (HMC HangManCLI) NewHangManCli() HangManCLI {
	//HMC:= HangManCLI{
	//	Reader: bufio.NewReader(os.Stdin),
	//}
	//HMC.Reader=bufio.NewReader(os.Stdin)

	HMC.setReader(bufio.NewReader(os.Stdin))
	return HMC
}

/**
This is a welcome screen that prints to the console.
 */
func (HMC HangManCLI) WelcomeScreen()  {
	fmt.Println("##################################################################\n\n")
	fmt.Println("#################### WELCOME TO HANGMAN ##########################\n\n")
}

/**
This uses ReadString from the Reader of the struct.
 */
func (HMC HangManCLI) GetWordToGuess() string {

	fmt.Print("Enter the word you like to be guessed:")
	var wordToGuess string

	wordToGuess, _ = HMC.Reader.ReadString('\n') // reading from the command line

	fmt.Println("you have chosen:", wordToGuess)

	return wordToGuess
}

func (HMC HangManCLI) GetLetterGuessed() byte {

	var letterGuessed byte
	fmt.Print("Guess a character:(0 for quit):")
	letterGuessed, _= HMC.Reader.ReadByte()

	fmt.Printf("\nyou have chose: %s %d\n", letterGuessed)

	return letterGuessed

}
