package views

import (
	"bufio"
	"fmt"
	"github.com/DandDevy/hangman_in_Golang/controllers"
	"os"
)
const(
	QuitChar = '0'
)
/**
HangManCLI is struct that has methods on it to sent and get data from the user.
 */
type HangManCLI struct {
	Reader      *bufio.Reader
	wordToGuess string
	currentGuess string
	controller  controllers.Controller
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
	HMC.Reader=bufio.NewReader(os.Stdin)

	HMC.controller.NewController().SetWord(HMC.GetWordToGuess())

	//HMC.controller.SetWord(HMC.GetWordToGuess())

	HMC.play()

	//HMC.setReader(bufio.NewReader(os.Stdin))
	return HMC
}

/**
This is a welcome screen that prints to the console.
 */
func (HMC HangManCLI) WelcomeScreen()  {
	fmt.Println("##################################################################")
	fmt.Println("##################################################################")
	fmt.Println("##################################################################\n\n")
	fmt.Println("#################### WELCOME TO HANGMAN ##########################\n\n")
	fmt.Println("##################################################################\n\n")
}

/**
This is a Good bye screen that prints to the console.
*/
func (HMC HangManCLI) GoodByeScreen()  {
	fmt.Println("##################################################################a")
	fmt.Println("\n\n############################ GOOD BYE ############################\n\n")
	fmt.Println("##################################################################")
	fmt.Println("##################################################################")
	fmt.Println("##################################################################")
}

/**
This uses ReadString from the Reader of the struct.
 */
func (HMC HangManCLI) GetWordToGuess() string {

	fmt.Print("Enter the word you like to be guessed:")
	var toGuess string

	toGuess, _ = HMC.Reader.ReadString('\n') // reading from the command line

	HMC.wordToGuess = toGuess



	fmt.Println("you have chosen:", toGuess)

	return toGuess
}

func (HMC HangManCLI) GetLetterGuessed() byte {

	fmt.Print("Guess a character(0 for quit):")
	letterGuessedAsByte, _ := HMC.Reader.ReadByte()

	fmt.Println("\n you have chosen: ", string(letterGuessedAsByte))
	//fmt.Printf("\nyou have chosen: %v %T \n", string(letterGuessedAsByte), string(letterGuessedAsByte))

	return letterGuessedAsByte

}

func (HMC HangManCLI) play() {

	HMC.WelcomeScreen()

	var keepPlaying = true

	for keepPlaying {
		guess := HMC.GetLetterGuessed()
		isRightGuess, updatedGuess:= HMC.controller.Check(guess)
		HMC.displayWord(updatedGuess)

		if isRightGuess {
			HMC.congratulationsScreen()
		}

		keepPlaying = HMC.controller.KeepPlaying()
	}

	HMC.GoodByeScreen()

}

func (HMC HangManCLI) congratulationsScreen() {
	fmt.Println("##################################################################\n\n")
	fmt.Println("###################### congratulations ###########################\n\n")
	fmt.Println("##################################################################\n\n")

}

func (HMC HangManCLI) displayWord(found string) {
	fmt.Println("Your current guess \n |-->", found)
}
