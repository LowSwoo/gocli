package gocli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Command struct {
	Name        string
	Description string
	Action      func(Context)
}

type Context struct {
	CommandParams []string
}

type App struct {
	commands []Command
}

func NewApp() *App {
	app := App{}
	basic_commands := []Command{
		{
			Name:        "exit",
			Description: "Exit the programm",
			Action: func(Context) {
				os.Exit(0)
			},
		},
		{
			Name:        "help",
			Description: "help",
			Action: func(Context) {
				for _, c := range app.commands {
					fmt.Printf("Command: %v\tDescription:%v\n", c.Name, c.Description)
				}
			},
		},
		{
			Name:        "clear",
			Description: "Clear the terminal",
			Action: func(Context) {
				fmt.Println("\033[2J")
			},
		},
	}
	app.AddComands(basic_commands)
	return &app
}

func (app *App) AddComands(commands []Command) {
	app.commands = append(app.commands, commands...)
}
func (app *App) HasCommand(name string) *Command {
	for _, c := range app.commands {
		if name == c.Name {
			return &c
		}
	}
	return nil
}

func (app *App) Run() {
	var inputString []string
	scanner := bufio.NewScanner(os.Stdin)
	var context Context
	for {
		fmt.Printf(">> ")
		if scanner.Scan() {
			inputString = strings.Split(scanner.Text(), " ")
		}
		c := app.HasCommand(inputString[0])
		if len(inputString) == 1 {
			context.CommandParams = nil
		} else {
			context.CommandParams = inputString[1:]
		}
		if c != nil {
			c.Action(context)
		} else {
			for _, c := range app.commands {
				fmt.Printf("Command: %v\tDescription:%v\n", c.Name, c.Description)
			}
		}
	}
}
