package main

import (
	"fmt"

	"time"
)

type Page struct {
	Id      int
	Title   string
	Content string
	Date    string
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

func Update(Id int, title, content string) int64 {
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
		return 0
	}
	ra, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	fmt.Println(ra)
	return ra
}

func Delete(id int) {
	sql := "DELETE FROM page WHERE id=? "
	stmt, err := Conn.Prepare(sql)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer stmt.Close()
	res, err := stmt.Exec(id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	lastid, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Last op id", lastid)

}

func Create(title, content string) int {

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
		return 0
	}
	ID, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return int(ID)
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
