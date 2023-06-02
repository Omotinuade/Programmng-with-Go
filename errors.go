package main

import "fmt"

func ain() {
	// PrintName()
	Listen()
}

func PrintName() {
	var name string

	fmt.Print("Name: ")
	panic("INPUT A NAME")
	fmt.Scan(&name)

	fmt.Printf(" name is %s\n")
}