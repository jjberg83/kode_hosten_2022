package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	//"os/exec"
	"bufio"
	"os"
	//"runtime"
)

// Task 7: Simple Shell
//
// This task focuses on building a simple shell that accepts
// commands that run certain OS functions or programs. For OS
// functions refer to golang's built-in OS and ioutil packages.
//
// The shell should be implemented through a command line
// application; allowing the user to execute all the functions
// specified in the task. Info such as [path] are command arguments
//
// Important: The prompt of the shell should print the current directory.
// For example, something like this:
//   /Users/meling/Dropbox/work/opsys/2020/meling-stud-labs/lab3>
//
// We suggest using a space after the > symbol.
//
// Your program should be able to at least the following functions:
// 	- exit
// 		- exit the program
// 	- cd [path]
// 		- change directory to a specified path
// 	- ls
// 		- list items and files in the current path
// 	- mkdir [path]
// 		- create a directory with the specified path
// 	- rm [path]
// 		- remove a specified file or folder
// 	- create [path]
// 		- create a file with a specified name
// 	- cat [file]
// 		- show the contents of a specified file
// 			- any file, you can use the 'hello.txt' file to check if your
// 			  implementation works
// 	- help
// 		- show a list of available commands
//
// You may also implement any number of optional functions, here are some ideas:
// 	- help [command]
// 		- give additional info on a certain command
// 	- ls [path]
// 		- make ls allow for a specified path parameter
// 	- rm -r
// 		WARNING: Be aware of where you are when you try to execute this command
// 		- recursively remove a directory
// 			- meaning that if the directory contains files, remove
// 			  all the files within the directory first, then the
// 			  directory itself
// 	- calc [expression]
// 		- Simple calculator program that can calculate a given expression
// 			- example expressions could be + - * \ pow
// 	- ipconfig
// 		- show ip interfaces
// 	- history
// 		- show command history
// 		- Alternatively implement this together with pressing up on your
// 		  keypad to load the previous command
// 		- clrhistory to clear history
// 	- tail [n]
// 		- show last n lines of a file
// 	- head [n]
// 		- show first n lines of a file
// 	- writefile [text]
// 		- write specified text to a specified file
//
// 	Or, alternatively, implement your own functionality not specified as you please
//
// Additional notes:
// 	- If you want to use colors in your terminal program you can see the package
// 		"github.com/fatih/color"
//
// 	- Helper functions may lead to cleaner code
//

// Terminal contains
type Terminal struct {
	// DONE (student): Add field(s) if necessary
	Command string
}

// I alle metodene jeg har laget under googlet jeg bare
// f.eks 'mkdir system call golang'

func (t *Terminal) Cat(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	text, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	for tall := range text {
		fmt.Print(string(text[tall]))

	}
}

//text, err != ioutil ReaAdll()

func (t *Terminal) Create(path, fileName string) {
	file, err := os.Create(path + "/" + fileName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("File created successfully")
		defer file.Close()
	}
}

func (t *Terminal) Rm(path string) {
	err := os.Remove(path)
	if err != nil {
		fmt.Println(err)
	}
}

func (t *Terminal) Cd(directory string) {
	chDirErr := os.Chdir(directory)
	if chDirErr != nil {
		fmt.Println(chDirErr)
	}
	_, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

}

func (t *Terminal) Exit() {
	os.Exit(0)
}

func (t *Terminal) Mkdir(directory string) {
	err := os.Mkdir(directory, 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func (t *Terminal) Ls(pwd string) {
	f, err := os.Open(pwd)
	if err != nil {
		fmt.Println(err)
		return
	}

	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range files {
		fmt.Println(v.Name(), v.IsDir())
	}
}

// Execute executes a given command
func (t *Terminal) Execute(command, argument, pwd string) string {
	// DONE (student): Make this run a given command
	switch command {
	// disse peker til metodene jeg har laget til structen Terminal ovenfor
	case "mkdir":
		t.Mkdir(argument)
		return ""

	case "ls":
		t.Ls(pwd)
		return ""

	case "exit":
		t.Exit()
		return ""
	case "cd":
		t.Cd(argument)
		return ""
	case "rm":
		t.Rm(argument)
		return ""
	case "create":
		t.Create(pwd, argument)
		return ""
	case "cat":
		t.Cat(argument)
		return ""
	}
	return ""
}

// This is the main function of the application.
// User input should be continuously read and checked for commands
// for all the defined operations.
// See hjtps://golang.org/pkg/bufio/#Reader and especially the ReadLine
// function.
func main() {
	reader := bufio.NewReader(os.Stdin)
	var t = Terminal{}
	// den evige loopen under oppfører seg som en Terminal
	// den viser pwd, venter på kommando, og når kommandoen
	// kommer, utfører den kommandoen. Deretter venter den
	// på ny kommando, og loopen gjentas...
	for {
		pwd, _ := os.Getwd()
		fmt.Print(pwd + " > ")
		les, _ := reader.ReadString('\n')
		// kode under gjør at første ord i det som ble lest over
		// (altså kommandoen), kommer på index 0, og alt etter
		// (altså argumenter) kommer på følgende indekser i listen
		// NB: Jeg har bare skrevet kode for 1 argument
		lesSlice := strings.Split(les, " ")
		kommando := strings.TrimSpace(lesSlice[0])
		var argument string
		if len(lesSlice) > 1 { // m.a.o - det er et argument etter kommandoen
			argument = strings.TrimSpace(lesSlice[1])
		}
		//fmt.Println("Kommando er ", kommando, "a")
		//fmt.Println("Argument er ", argument, "a")
		t.Execute(kommando, argument, pwd)
	}

}
