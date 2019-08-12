package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func get_word_to_guess()  {

}

func main()  {
	fmt.Println("################Welcome to Hangman###################")

	var play_game bool = true

	reader := bufio.NewReader(os.Stdin)

	for play_game {

		fmt.Print("Would you like to leave? >>>")
		text, _ := reader.ReadString('\n')

		//text = text[:len(text) -1] // remove the \n

		text = strings.TrimSuffix(text, "\n") //another way of removing \n

		fmt.Println(text)

		if text == "yes" {
			play_game = false
		}

	}
}
