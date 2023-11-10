package main

import "fmt"

type Event struct {
	Name         string
	Date         string
	Participants int
}

func main() {
	birthday := Event{Name: "Birthday", Date: "24/11/94", Participants: 29}
	halloween := Event{Name: "Halloween", Date: "31/10/23", Participants: 2}

	fmt.Printf("Event name is \"%v\". We will celebrate it on %v and %v peoples are invited\n", birthday.Name, birthday.Date, birthday.Participants)
	fmt.Printf("Event name is \"%v\". We will celebrate it on %v and %v peoples are invited\n", halloween.Name, halloween.Date, halloween.Participants)

	newyear := CreateEvent("New Year", "31/12/23", 2)
	fmt.Printf("Event name is \"%v\". We will celebrate it on %v and %v peoples are invited\n", newyear.Name, newyear.Date, newyear.Participants)
}

func CreateEvent(name, date string, participants int) Event {
	return Event{name, date, participants}
}
