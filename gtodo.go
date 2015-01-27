package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var itemCount = 0

//PLACE HOLDER USE A FILE INSTEAD

var todoTags = [100]string{}
var todoName = [100]string{}
var todoDescription = [100]string{}

/////////////////////////////////

func main() {
	fmt.Println("> Viewing TODO List (", os.Args, ")\n")

	if len(os.Args) > 1 {
		if os.Args[1] == "view" {
			readList()
			viewTodoList(0, 0)
		} else if os.Args[1] == "add" {
			readList()
			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Tag: ")
			tag, _ := reader.ReadString('\n')

			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')

			fmt.Print("Description: ")
			description, _ := reader.ReadString('\n')

			tag = strings.TrimSpace(tag)
			name = strings.TrimSpace(name)
			description = strings.TrimSpace(description)

			addItem(tag, name, description)
			writeList()
		} else if os.Args[1] == "rem" {
			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			removeItem(name)
		}
	} else {
		fmt.Println("No Args Given!")
	}
}

func viewTodoList(sortType int, sortRange int) {
	paddingTag := 5
	paddingName := 30
	paddingDecription := 30

	fmt.Println(padLine("TAG", paddingTag, " "), "| ", padLine("NAME", paddingName, " "), "| ", padLine("DESCRIPTION", paddingDecription, " "))
	fmt.Println(padLine("", paddingTag, "-"), "| ", padLine("", paddingName, "-"), "| ", padLine("", paddingDecription, "-"))

	for i := range todoName {
		if todoName[i] != "" {
			fmt.Println(padLine("", paddingTag, " "), "| ", padLine("", paddingName, " "), "| ", padLine("", paddingDecription, " "))
			fmt.Println(padLine(todoTags[i], paddingTag, " "), "| ", padLine(todoName[i], paddingName, " "), "| ", padLine(todoDescription[i], paddingDecription, " "))
		}
	}
}

func padLine(s string, targLength int, pad string) string {
	paddedString := s
	padLen := targLength - len(s)

	if len(s) > targLength {
		return s[0:targLength-3] + "..."
	} else {
		for i := 0; i < padLen; i++ {
			paddedString += pad
		}
	}
	return paddedString
}

//N/A
func getName() {

}

func getTag(itemName string) string {
	for i := range todoName {
		if todoName[i] == itemName {
			return todoTags[i]
		}
	}
	return ""
}

func getDescription(itemName string) string {
	for i := range todoName {
		if todoName[i] == itemName {
			return todoDescription[i]
		}
	}
	return ""
}

func addItem(tag string, name string, description string) bool {
	todoTags[itemCount] = tag
	todoName[itemCount] = name
	todoDescription[itemCount] = description
	//If operation completes succesful return true, else false
	return true
}

func removeItem(name string) bool {
	for i := range todoName {
		if todoName[i] == name {
			todoName[i] = ""
			todoTags[i] = ""
			todoDescription[i] = ""

		}
	}
	writeList()
	//If operation completes succesful return true, else false
	return false
}

func writeList() {
	s := ""
	for i := range todoName {
		if todoName[i] != "" {
			s += todoTags[i] + " -BRK- " + todoName[i] + " -BRK- " + todoDescription[i] + " -N-"
		}
	}
	ioutil.WriteFile("gtodo.dat", []byte(s), 0644)
}

func readList() {
	data, _ := ioutil.ReadFile("gtodo.dat")
	dataLined := strings.Split(string(data), "-N-")

	for i := range dataLined {
		dataSplit := strings.Split(dataLined[i], " -BRK- ")

		for i2 := range dataSplit {
			if i2 == 0 {
				todoTags[i] = dataSplit[i2]
			} else if i2 == 1 {
				todoName[i] = dataSplit[i2]
			} else if i2 == 2 {
				todoDescription[i] = dataSplit[i2]
			}
		}

		if dataSplit[0] != "" {
			itemCount++
		}

	}
}
