package main

import (
	"github.com/DandDevy/hangman_in_Golang/views"
)

var (
	HMC views.HangManCLI
)


func main()  {

	//HMC:= views.HangManCLI{Reader: bufio.NewReader(os.Stdin)}


	HMC.NewHangManCli()

	//fmt.Println(HMC.Reader)

	//	var wordToGuess string
	//	wordToGuess, _ =HMC.Reader.ReadString('\n')
	//fmt.Printf("wordToGuess: %s\n", wordToGuess)

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
}
