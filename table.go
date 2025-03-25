package main

import "fmt"

func printTableHeader(idWidth, nameWidth int, firstHeaderText, secondHeaderText string) {
	fmt.Print("╔")
	for i := 0; i < idWidth+2; i++ {
		fmt.Print("═")
	}
	fmt.Print("╦")
	for i := 0; i < nameWidth+2; i++ {
		fmt.Print("═")
	}
	fmt.Println("╗")

	fmt.Printf("║ %-*s ║ %-*s ║\n", idWidth, firstHeaderText, nameWidth, secondHeaderText)

	fmt.Print("╠")
	for i := 0; i < idWidth+2; i++ {
		fmt.Print("═")
	}
	fmt.Print("╬")
	for i := 0; i < nameWidth+2; i++ {
		fmt.Print("═")
	}
	fmt.Println("╣")

}

func printTableRow(id int, name string, idWidth, nameWidth int) {
	fmt.Printf("║ %-*d ║ %-*s ║\n", idWidth, id, nameWidth, name)
}

func printTableFooter(idWidth, nameWidth int) {
	fmt.Print("╚")
	for i := 0; i < idWidth+2; i++ {
		fmt.Print("═")
	}
	fmt.Print("╩")
	for i := 0; i < nameWidth+2; i++ {
		fmt.Print("═")
	}
	fmt.Print("╝\n")
}

func printBorder(idWidth, nameWidth int, borderChar string) {
	fmt.Print("╠")
	for i := 0; i < idWidth+2; i++ {
		fmt.Print(borderChar)
	}
	fmt.Print("╦")
	for i := 0; i < nameWidth+2; i++ {
		fmt.Print(borderChar)
	}
	fmt.Println("╣")
}
