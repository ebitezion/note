package cmd

import (
	"bufio"
	"fmt"
	"os"
)

type Item struct {
	Id      int
	Title   string
	Content string
	Date    string
}

type Items struct {
	Item []Item
}

/////////////////////////////////////////////////////////////////////
// Helper func
/////////////////////////////////////////////////////////////////////

func GetItem(id int) Item {
	//find id from db

	//handle err

	//print output
	return Item{}
}

/////////////////////////////////////////////////////////////////////////////
//DB helper func
////////////////////////////////////////////////////////////////////////////
//dummy data
var dummy = []map[string]Item{
	{"1": {1, "my grocery task", "purchase at dollar store ok", "12-03-21"}},
	{"2": {2, "my reading task", "Read the bible back to back", "12-03-21"}},
	{"3": {3, "my work task", "Programme the app mehn	", "12-03-21"}},
}

var DB = []Item{{1, "my grocery task", "purchase at dollar store ok", "12-03-21"},
	{2, "my reading task", "Read the bible back to back", "12-03-21"},
	{3, "my work task", "Programme the app mehn	", "12-03-21"},
}

func DBGet(id int) (i Item, err error) {
	//get item from db based on id

	//manage err
	if err != nil {
		return
	}

	i = DB[id]
	return
}

func WriteFileDB(filename string, content Item) {

	//open the file
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {

		return
	}
	defer f.Close()

	fmt.Fprintln(f, content)
	//scanner := bufio.NewWriter(f)
	/*	for _, val := range content {
		info, err := scanner.Write([]byte(val))
		if err != nil {
			println("err", err)
			return
		}
		scanner.Flush()

	}*/

}

func LoadPage(filename string) {
	_, items, err := ReadFileDB(filename)

	if err != nil {
		fmt.Errorf("Error %v \n", err)
		return
	}
	if items == nil {
		fmt.Errorf("Item(s) :%v \n", items)
		return
	}

	for _, item := range items {
		fmt.Print("TODO")
		fmt.Println("\t", string(item[0:]), "\t", string(item[1]), "\t\t\t", string(item[5:]))
	}
}

func ReadFileDB(filename string) (int, []string, error) {
	//lets write to a file
	f, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("error with oprning file %v", err)
		return 0, nil, err
	}
	defer f.Close()
	//ioutil.ReadFile(filename)
	scanner := bufio.NewScanner(f)
	var (
		i        int = 0
		str      []string
		inputstr string
	)
	for scanner.Scan() {
		i++
		inputstr = scanner.Text()

		str = append(str, inputstr)

	}
	println("i?", i)
	return i, str, nil
}
