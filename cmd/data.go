package cmd

import (
	"bufio"
	"fmt"
	"os"
	"time"
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

/*var dummydb = items{

	item: []Item{{1, "my grocery task", "purchase at dollar store ok", "12-03-21"},
		{2, "my reading task", "Read the bible back to back", "12-03-21"},
		{3, "my work task", "Programme the app mehn	", "12-03-21"},
	},
}*/
var DB = []Item{{1, "my grocery task", "purchase at dollar store ok", "12-03-21"},
	{2, "my reading task", "Read the bible back to back", "12-03-21"},
	{3, "my work task", "Programme the app mehn	", "12-03-21"},
}

func Create(filename, title, content string) {
	//ReadFileDB(filename)
	//get last id from db and increment by 1
	lastID := 1 //len(DB)

	//increment the last id by 1
	lastID += 1

	//get current date time
	t := time.Now()
	today := t.Format("Mon Jan _2 15:04:05 2006")

	//append all param to db e.g 	{3, "my work task", "Programme the app mehn	", "12-03-21"},
	WriteFileDB(filename, Item{int(lastID), title, content, today})

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

func Load(page *Item) {

}

func ReadFileDB(filename string) (int, []string, error) {
	//lets write to a file
	f, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("error with oprning file %v", err)
		return 0, nil, err
	}
	defer f.Close()
	scanner := bufio.NewReader(f)
	var (
		i   int = 0
		str []string
	)
	for {
		inputstr, err := scanner.ReadString('\n')
		if err != nil {
			return 0, nil, err
		}
		str = append(str, inputstr)
		i += 1

	}
	return i, str, nil
}
