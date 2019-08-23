package models

import (
	"strings"
)

/**
hangMan is a game of hang-man that can be initialised by NewHangMan.
 */

const maxNumberOfAttempts int  = 10

type HangMan struct {
	wordToGuess         string
	numAttempts         int
}

//func NewHangMan(wordToGuess string) *hangMan {
//	return &hangMan{wordToGuess: wordToGuess}
//}

func (HM *HangMan) WordToGuess() string {
	return HM.wordToGuess
}

func (HM *HangMan) SetWordToGuess(wordToGuess string) {
	HM.wordToGuess = wordToGuess
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
func (HM HangMan) UpdateWordToGuess(correctGuess string)  {
	strings.Replace(HM.wordToGuess, correctGuess, "", -1)
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
func (HM HangMan) CheckWord(guess string) bool {

	res := false

	if guess != "" {

		if strings.Contains(HM.wordToGuess, guess) {
			res = true
		}
	}

	return res
}



