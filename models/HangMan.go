package models


/**
HangMan is a game of hang-man that can be initialised by NewHangMan.
 */
type HangMan struct {
	wordToGuess         string
	numAttemptsLeft     int
	maxNumberOfAttempts int

}

/**
Initialises the HangMan.
Takes the word to guess as a parameter.
 */
func (HM HangMan) NewHangMan(WTG string) HangMan{
	HM.wordToGuess = WTG
	HM.maxNumberOfAttempts = 10
	HM.numAttemptsLeft = HM.maxNumberOfAttempts
	return HM
}

