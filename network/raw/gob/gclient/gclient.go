package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

type Person struct {
	Name  Name
	Email []Email
}

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	Kind    string
	Address string
}

func (p Person) String() string {
	s := p.Name.Personal + " " + p.Name.Family
	for _, v := range p.Email {
		s += "\n" + v.Kind + ": " + v.Address
	}
	return s
}

func main() {
	person := Person{
		Name: Name{Family: "Newmarch", Personal: "Jan"},
		Email: []Email{Email{Kind: "home", Address: "[email protected]"},
			Email{Kind: "work", Address: "[email protected]"}}}
	service := "0.0.0.0:1200"
	conn, err := net.Dial("tcp", service)
	checkError(err)
	
	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)
	
	for i := 0; i < 4; i++ {
		encoder.Encode(person)
		var newPerson Person
		decoder.Decode(&newPerson)
		fmt.Println("==========CLIENT==========")
		fmt.Println(newPerson.String())
	}
	conn.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
