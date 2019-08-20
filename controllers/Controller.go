package controllers

import (
	"hangman_in_Golang/models"
)

/**
Struct carrying reference to HangMan game for Command Line Interface.
 */
type Controller struct {
	HangMan models.HangMan
}

func (C Controller) NewController(HM models.HangMan)  Controller{
	C.setHangMan(HM)
	return C
}

/**
Getter for reference variable for HangMan game.
 */
func (C Controller) getHangMan()  models.HangMan{
	return C.HangMan
}

/**
Setter for reference for HangMan game.
 */
func (C Controller) setHangMan(HM models.HangMan) {
	C.HangMan = HM
}

func start()  {
	
}

func setWordToGuess(WTG string)  {

}