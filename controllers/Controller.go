package controllers

import (
	"github.com/DandDevy/hangman_in_Golang/models"
)

/**
Struct carrying reference to hangMan game for Command Line Interface.
 */
type Controller struct {
	hangMan models.HangMan
}

/**
Constructor for Controller
 */
func (c Controller) NewController()  Controller{
	c.hangMan.NewHangManWithoutParam()

	return c
}

func Start()  {

}
/**
Sets word to guess for the hang man game.
 */
func (c Controller) SetWord(word string) {
	c.hangMan.SetWordToGuess(word)
}

/**
Getter for reference variable for hangMan game.
 */
func (c Controller) GetHangMan()  models.HangMan{
	return c.hangMan
}

/**
Setter for reference for hangMan game.
 */
func (c Controller) SetHangMan(HM models.HangMan) {
	c.hangMan = HM
}

/**
Gets the size of the word to guess/number of letters left to guess.
 */
func (c Controller) GetSizeOfWordToGuess() int {

	return c.hangMan.GetSizeOfWordToGuess()
}

/**
Returns boolean to allowing game to go on.
*/
func (c Controller) KeepPlaying() bool {

	return c.hangMan.KeepPlaying()
}

/**
Checks a byte and returns a bool and string of current desk.
 */
func (c Controller) Check(guess byte) (bool, string) {
	return c.hangMan.CheckWord(guess)
}

func (c Controller) GetFoundWord() string {

	return c.hangMan.GetUserGuess()
}


