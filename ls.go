package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/color"
)

type ls struct {
	command        []string
	command_output []fs.FileInfo
	hidden bool
	full   bool
}

func main() {
	a := ls{command: os.Args, hidden: false, full: false}
	a.get()
	a.put()
}
func (data *ls) get() {
	var command string
	if len(data.command) > 1 {
		for _, val := range data.command {
			if val == "-a"   {
				data.hidden = true
			} else if val == "-l"  {
				data.full = true
			} else if val == "-la" {
				data.hidden = true
				data.full = true
			} else {
				command = data.command[len(data.command)-1]
			}
		}
		if data.hidden || data.full {
			if len(data.command) > 2 {
				command = data.command[len(data.command)-1]
			} else {
				command = "."
			}
		}
	} else {
		data.hidden = false
		command = "."
	}
	out, err := ioutil.ReadDir(command)
	if err != nil {
		log.Fatal(err)
	}
	data.command_output = out
}

func (data *ls) put() {
	for _, file := range data.command_output {
		if !data.hidden {
			if string(file.Name()[0]) == "." {
				continue
			}
		}
		if data.full {
			output := color.New(color.FgGreen).SprintFunc()
			output2 := color.New(color.FgBlue).SprintFunc()
			if file.IsDir() {
				output3 := color.New(color.FgYellow).SprintFunc()
				fmt.Printf("%s %s %s  \n", output(file.Mode()), output2(file.ModTime().String()), output3(file.Name()))
			} else {
				output3 := color.New(color.FgCyan).SprintFunc()
				fmt.Printf("%s %s %s \n", output(file.Mode()), output2(file.ModTime().String()), output3(file.Name()))
			}
		} else {
			if file.IsDir() {
				color.Yellow(file.Name())
			} else {
				color.Cyan(file.Name())
			}
		}
	}
}
