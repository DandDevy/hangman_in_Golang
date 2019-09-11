package models

import (
	"strings"
)

/**
hangMan is a game of hang-man that can be initialised by NewHangMan.
 */

const maxNumberOfAttempts int  = 10

type HangMan struct {
	wordToGuess		string
	numAttempts		int
	currentGuess	string
	userGuess		string
}

//func NewHangMan(wordToGuess string) *hangMan {
//	return &hangMan{wordToGuess: wordToGuess}
//}

func (HM *HangMan) WordToGuess() string {
	return HM.wordToGuess
}

func (HM *HangMan) SetWordToGuess(wordToGuess string) {
	HM.wordToGuess = wordToGuess
	for i := 0; i < len(wordToGuess); i++ {
		HM.currentGuess += "_"
	}
}


/**
Constructor for the hang man.
Takes the word to guess as a parameter.
 */
func (HM HangMan) NewHangMan(WTG string) HangMan{
	HM.wordToGuess = WTG
	HM.numAttempts = 0
	return HM
}

/**
Constructor for game without any parameters
*/
func (HM HangMan) NewHangManWithoutParam() HangMan{
	HM.numAttempts = 0
	return HM
}

/**
Returns boolean to allowing game to go on.
 */
func (HM HangMan) KeepPlaying() bool {
	res := false
	if HM.numAttempts < maxNumberOfAttempts || len(HM.wordToGuess) == 0{
		res = true
	}

	return res
}

/**
updates the word to guess. By removing the letter(s) that was correctly guessed.
 */
func (HM HangMan) updateWordToGuess(correctGuess byte) string {
	var newWord string
	strings.Replace(HM.wordToGuess, string(correctGuess), "", -1)

	return newWord
}

/**
Gets the size of the word to guess/number of letters left to guess.
*/
func (HM HangMan) GetSizeOfWordToGuess() int {
	return len(HM.wordToGuess)
}

/**
Checks if the guess is correct. Returns bool.
 */
func (HM HangMan) CheckWord(guess byte) (bool, string) {

	isGoodGuess := false
	println("guess in Checkword:", string(guess))


	if HM.KeepPlaying() {

		if guess != 0 {	//In golang when something is allocated it is given the value of 0 at the beginning

			if strings.Contains(HM.wordToGuess, string(guess)) {

				println("word to guess contains the guess")
				isGoodGuess = true
				HM.updateWordToGuess(guess)

			} else {
				HM.numAttempts++
			}
		}
	}

	return isGoodGuess,HM.GetUserGuess()
}

/**
Get the the string of current guesses of user ( _ * length of word at start)
 */
func (HM *HangMan) GetUserGuess() string {
	return HM.userGuess
}







