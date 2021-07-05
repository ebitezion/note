package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Conn      *sql.DB
	err       error
	UserInput string
)

/*
 The database constants
*/
//var db = page.DBHolder{}

const (
	Port   = ":8080"
	DBhost = "localhost"
	DBuser = "root"

	DBport = "3306"
	DBname = "notedb"
)

func init() {
	connString := fmt.Sprintf("%s:@tcp(%s:%s)/%s?parseTime=true", DBuser, DBhost, DBport, DBname)

	Conn, err = sql.Open("mysql", connString)
	if err != nil {
		panic(err)

	}
	if err := Conn.Ping(); err != nil {
		fmt.Printf("Db failed to ping: %s", err)
	}

}

func main() {

	defer Conn.Close()
	App()
}

func App() {
	header()
	//display the list of items home page
	ReadAll()

	//Prompt for input
	prompt()
	//Responding to user input
	//now text is our command captured
	command()
}
func header() {
	fmt.Println("======================================================================================================")
	fmt.Println("\t\t\t                           *NOTE TAKE APP*")
	fmt.Println("======================================================================================================")

}
func prompt() {

	fmt.Printf("[Press -d to delete followed by ID]\n [Press -c to create] \n [Press -v followed by Id to view] \n[Press -x to exit]\n")
	fmt.Println("======================================================================================================")
	fmt.Print(">>>")
	//capture input
	input := os.Stdin
	scanner := bufio.NewScanner(input)
	//var text string
	if scanner.Scan() {
		UserInput = scanner.Text()
		fmt.Println(UserInput)
	}
}
func command() {
	//use the switch case
	args := strings.Fields(UserInput)
	switch args[0] {
	case "-d":
		id, err := strconv.ParseInt(args[1], 10, 32)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		Delete(int(id))
		App()
	case "-c":
		CreatePage()
	case "-v":
		id, err := strconv.ParseInt(string(args[1]), 10, 32)
		if err != nil {
			fmt.Println("Parsing error: ", err.Error())
			return
		}
		DetailsPage(int(id))
	case "-x":
		os.Exit(1)
	default:
		App()
	}

}
func CreatePage() {
	var title, content string
	//prompt user for title and content
	fmt.Println("ENTER TITLE >>>")
	input := os.Stdin
	scanner := bufio.NewScanner(input)

	if scanner.Scan() {
		title = scanner.Text()

	}
	fmt.Println("ENTER CONTENT >>>")
	scanner = bufio.NewScanner(input)
	if scanner.Scan() {

		content = scanner.Text()

	}
	Create(title, content)
	App()

}
func EditPage(id int) {
	var title, content string
	//prompt user for title and content
	fmt.Println("ENTER TITLE >>>")
	input := os.Stdin
	scanner := bufio.NewScanner(input)

	if scanner.Scan() {
		title = scanner.Text()

	}
	fmt.Println("ENTER CONTENT >>>")
	scanner = bufio.NewScanner(input)
	if scanner.Scan() {

		content = scanner.Text()

	}
	Update(id, title, content)
	DetailsPage(id)

}

func DetailsPage(Id int) {
	header()
	Read(Id)
	Detailsprompt()
	Detailscommand()
}

func Detailsprompt() {

	fmt.Printf("[Press -e followed by ID to edit]\n [Press -d followed by ID to delete] \n [Press -h to go to home page] \n[Press -x to exit]\n")
	fmt.Println("======================================================================================================")
	fmt.Print(">>>")
	//capture input
	input := os.Stdin
	scanner := bufio.NewScanner(input)
	//var text string
	if scanner.Scan() {
		UserInput = scanner.Text()
		fmt.Println(UserInput)
	}
}

func Detailscommand() {
	//use the switch case
	args := strings.Fields(UserInput)
	switch args[0] {
	case "-e":
		id, err := strconv.ParseInt(args[1], 10, 32)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		EditPage(int(id))
	case "-d":
		id, err := strconv.ParseInt(args[1], 10, 32)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		Delete(int(id))
		DetailsPage(int(id))
	case "-h":
		App()
	case "-x":
		os.Exit(1)
	default:
		App()
	}

}
