package cmd

type item struct {
	id      int
	title   string
	content string
	date    string
}

type items struct {
	item []item
}

/////////////////////////////////////////////////////////////////////
// Helper func
/////////////////////////////////////////////////////////////////////

func GetItem(id int) item {
	//find id from db

	//handle err

	//print output
	return item{}
}

/////////////////////////////////////////////////////////////////////////////
//DB helper func
////////////////////////////////////////////////////////////////////////////
//dummy data
var dummy = []map[string]item{
	{"1": {1, "my grocery task", "purchase at dollar store ok", "12-03-21"}},
	{"2": {2, "my reading task", "Read the bible back to back", "12-03-21"}},
	{"3": {3, "my work task", "Programme the app mehn	", "12-03-21"}},
}

/*var dummy2 = items{
	[
		{1, "my grocery task", "purchase at dollar store ok", "12-03-21"},
	]
}*/

func DBGet(id int) (item item, err error) {
	//get item from db based on id

	//manage err
	if err != nil {
		return item, err
	}
	return
}
