package main

import (
	"hangman_in_Golang/views"
)


func main()  {

		//hmc:= views.HangManCli{Reader: bufio.NewReader(os.Stdin)}

		var hmc views.HangManCli

		hmc = hmc.NewHangManCli()

		hmc.WelcomeScreen()

		wordToGuess := hmc.GetWordToGuess()



		//fmt.Println(hmc.Reader)

	//	var wordToGuess string
	//	wordToGuess, _ =hmc.Reader.ReadString('\n')
	//fmt.Printf("wordToGuess: %s\n", wordToGuess)



	var playGame bool = true

	//tester := Tester{"test1"}
	//
	//tester.name = "asd"
	//
	//
	//test2 := Tester{}
	//test2.newTester()
	//
	//fmt.Printf("test: %d\n", test2)
	//
	//
	//t := Tester{name:"asd"}
	//
	//test3 := &t
	//
	//fmt.Printf("test3: %p\n",*test3)
	//*test3.name="newName"

	for playGame {
		char := hmc.GetLetterGuessed()

		if char == '0' {
			playGame = false
		} else if char == wordToGuess[0] {
			break
		}

	}
}
