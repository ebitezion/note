package main

import "testing"

func TestCreate(t *testing.T) {
	if id := Create("Test title", "Test Content"); id == 10 {
		t.Logf("Expected %v Got %v", 9, id)
	}

}

func TestUpdate(t *testing.T) {
	if id := Update(2, "title testing", "content testing"); id == 2 {
		t.Logf("Expected %v Got %v", 2, id)
	}
}
