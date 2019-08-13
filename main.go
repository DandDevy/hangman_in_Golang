package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getWordToGuess() string {
	var wordToGuess string

	fmt.Print("Enter the word to guess>>>")

	reader := bufio.NewReader(os.Stdin)
	wordToGuess, _ = reader.ReadString('\n')
	wordToGuess = strings.TrimSuffix(wordToGuess, "\n")

	fmt.Println("You have chosen:", wordToGuess)

	return wordToGuess

}

func main()  {
	fmt.Println("################Welcome to Hangman###################")


		wordToGuess := getWordToGuess()


	var playGame bool = true

	reader := bufio.NewReader(os.Stdin)

	for playGame {

		fmt.Print("Would you like to leave? >>>")
		text, _ := reader.ReadString('\n')

		//text = text[:len(text) -1] // remove the \n

		text = strings.TrimSuffix(text, "\n") //another way of removing \n

		fmt.Println(text)

		if text == "yes" {
			playGame = false
		}

	}
}
