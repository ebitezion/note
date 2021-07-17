package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

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
	//fmt.Println(cmd.DBGet(1))
	App()
	//cmd.FileDB("log.txt")
	//Create("MAGA-more than just cliche", "You just have to love the Trumpism ideology")
	//cmd.LoadPage("log.txt")
	//Read(2)
	//RaedAll()
	//Update(2, "Winnifred- the dazzling Queen", "Ever so often we are blessed with")
}

func App() {
	header()
	//display the list of items home page
	ReadAll()
	fmt.Println("\n\n")

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
	fmt.Println("\n\n")

}
func prompt() {

	fmt.Printf("[Press -d to delete followed by ID]\n [Press -c to create] \n [Press -v followed by Id to view] \n")
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
		print("delete ...")
	case "-c":

		//prompt user for title and content
		fmt.Println("ENTER TITLE >>>")
		input := os.Stdin
		scanner := bufio.NewScanner(input)
		if scanner.Scan() {
			title := scanner.Text()
			fmt.Println("ENTER CONTENT >>>")
			content := scanner.Text()
			Create(title, content)
			fmt.Println("Created")

		}

	case "-v":
		id, err := strconv.ParseInt(string(args[1]), 10, 32)
		if err != nil {
			fmt.Println("Parsing error: ", err.Error())
			return
		}
		DetailsPage(int(id))
	default:
		os.Exit(1)
	}

}
func Create(title, content string) {

	//get current date time
	t := time.Now()
	today := t.Format("Mon Jan _2 15:04:05 2006")

	stmt, err := Conn.Prepare("Insert into page(title,content,date) Values(?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	res, err := stmt.Exec(title, content, today)
	if err != nil {
		fmt.Println(err)
		return
	}
	ID, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("lastID: ", ID)
}

type Page struct {
	Id      int
	Title   string
	Content string
	Date    string
}

func Read(id int) {
	//query table with the id
	stmt, err := Conn.Prepare("Select * from page where id=?")
	if err != nil {
		panic(err.Error())

	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		panic(err.Error())
	}
	page := Page{}
	for rows.Next() {
		err := rows.Scan(&page.Id, &page.Title, &page.Content, &page.Date)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(page.Id, "\t", page.Title, "\t", page.Content, "\t", page.Date)
	}
}
func DetailsPage(Id int) {
	header()
	Read(Id)
	Detailsprompt()
	Detailscommand()
}
func Detailsprompt() {

	fmt.Printf("[Press -E followed by ID to edit]\n [Press -d followed by ID to delete] \n [Press -H to go to home page] \n")
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
	switch UserInput {
	case "-E":
		print("Create ...")
	case "-D":
		print("Delete ...")
	case "-H":
		print("view ...")
	default:
		os.Exit(1)
	}

}

func ReadAll() {
	//query table with the id
	rows, err := Conn.Query("Select * from page ")
	if err != nil {
		panic(err.Error())

	}
	defer rows.Close()
	page := Page{}
	for rows.Next() {
		err := rows.Scan(&page.Id, &page.Title, &page.Content, &page.Date)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(page.Id, "\t", page.Title, "\t", page.Date)

	}
}

func Update(Id int, title, content string) {
	stmt, err := Conn.Prepare("Update page set title=?, content=?, date=? where id=?")
	if err != nil {
		panic(err.Error())

	}
	defer stmt.Close()
	t := time.Now()
	date := t.Format("Mon Jan _2 15:04:05 2006")
	res, err := stmt.Exec(title, content, date, Id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	ra, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(ra)
}
