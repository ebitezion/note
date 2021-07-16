package cmd

import "testing"

func TestData(t *testing.T) {
	_, err := DBGet(1)
	if err != nil {
		t.Error("Expected Error as Nil Got ", err)
	}

}

func TestCreate(t *testing.T) {

	Create("Donald Trump- The leader with passion for the USA", "Today i listened to him speak and was reminded of the touching speech made by Martin Luther king")
	println("")
	t.Fail()

}

func TestReadFileDB(t *testing.T) {
	ReadFileDB("log.txt")
	t.Fail()
}
