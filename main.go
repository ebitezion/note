package main

import (
	"bufio"
	"fmt"
	"note/cmd"
	"os"
)

func main() {
	//fmt.Println(cmd.DBGet(1))
	//App()
	//cmd.FileDB("log.txt")
	cmd.Create("log.txt", "Donald Trump for President 2024", "The rightwing will shine again")
}

func App() {
	fmt.Println("======================================================================================================")
	fmt.Println("\t\t\t                           *NOTE TAKE APP*")
	fmt.Println("======================================================================================================")
	fmt.Println("\n\n")
	//display the list of items home page
	//itemsPage()
	fmt.Println("\n\n")

	//Responding to user input

	//Prompt for input

	fmt.Printf("[Press -d to delete followed by ID]\n [Press -c to create] \n [Press -v followed by Id to view] \n")
	fmt.Println("======================================================================================================")
	fmt.Print(">>>")
	//capture input
	input := os.Stdin
	scanner := bufio.NewScanner(input)
	var text string
	if scanner.Scan() {
		text = scanner.Text()
		fmt.Println(text)
	}
	//now text is our command captured
	//use the switch case
	switch text {
	case "-d":
		print("delete ...")
	case "-c":
		print("create ...")
	case "-v":
		print("view ...")
	default:
		print("unknown command ...")

	}
}

/*func itemsPage(page cmd.Item) {
	//populate the items struct with dbdata
	var page cmd.Item
	//fmt.Println(cmd.DB)
	for i, v := range cmd.DB {
		page = append(page, v)
		fmt.Println(page[i].Id, "\t", page[i].Title, "\t", "\t", page[i].Date)

	}

}*/
