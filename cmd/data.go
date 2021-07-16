package cmd

import (
	"fmt"
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

func Create(title, content string) {
	//get last id from db and increment by 1
	lastID := len(DB)

	//println("The last id is", lastID)

	//increment the last id by 1
	lastID += 1
	//println("The incremented is", lastID)

	//get current date time
	t := time.Now()
	today := t.Format("Mon Jan _2 15:04:05 2006")
	println(today)

	//append all param to db e.g 	{3, "my work task", "Programme the app mehn	", "12-03-21"},
	newdata := append(DB, Item{int(lastID), title, content, today})
	fmt.Println(newdata[lastID-1])
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
