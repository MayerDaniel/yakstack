package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func initStack() []string {
	arr := []string{}
	usr, err := user.Current()
	check(err)
	file, err := os.Open(usr.HomeDir + "/.yakstack")
	if err != nil {
		return arr
	}
	check(err)
	defer file.Close()
	decoder := json.NewDecoder(file)
	stackConf := []string{}
	err = decoder.Decode(&stackConf)
	if err != nil {
		return arr
	}
	check(err)
	return stackConf
}

func saveStack(arr []string) {
	usr, err := user.Current()
	check(err)
	file, _ := json.MarshalIndent(arr, "", " ")
	err = ioutil.WriteFile(usr.HomeDir+"/.yakstack", file, 0644)
	check(err)
}

func usage() {
	commands := `The commands are:

     peek     show current task
     push     put a new task on top of the stack
     pop      pop the current task off of the top of the stack
     list     show all tasks`
	fmt.Printf("Usage: %s <command> \n\n", os.Args[0])
	fmt.Printf("%s\n\n", commands)
	os.Exit(0)
}

func peek(stack []string) {
	if len(stack) < 1 {
		fmt.Printf("You have no tasks!\n")
	} else {
		fmt.Printf("Your current task is %s\n", stack[0])
	}
}

func push(stack []string, task string) []string {
	fmt.Printf("Task '%s' pushed to stack!\n", task)
	return append(stack, task)
}

func pop(stack []string) []string {
	if len(stack) < 1 {
		fmt.Printf("You have no tasks!\n")
		return stack
	}
	last := len(stack) - 1
	fmt.Printf("Task '%s' completed!\n", stack[last])
	return stack[:last]

}

func list(stack []string) {
	if len(stack) < 1 {
		fmt.Printf("You have no tasks!\n")
	} else {
		fmt.Printf("Tasks in order of priority:\n\n")
		n := 1
		for i := len(stack) - 1; i >= 0; i-- {
			fmt.Printf("\t%d.\t%s\n", n, stack[i])
			n++
		}
		fmt.Printf("\n")
	}
}

func main() {
	stack := initStack()
	args := os.Args
	if len(args) < 2 {
		usage()
	}
	if len(args) > 2 && args[1] != "push" {
		fmt.Printf("ERROR: too many positional arguments.\n\n")
		usage()
	}
	switch arg := args[1]; arg {
	case "peek":
		peek(stack)
	case "push":
		stack = push(stack, strings.Join(args[2:], " "))
	case "pop":
		stack = pop(stack)
	case "list":
		list(stack)
	default:
		fmt.Printf("ERROR: invalid command.\n\n")
		usage()
	}
	saveStack(stack)

}
