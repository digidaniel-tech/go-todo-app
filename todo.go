package main

import (
	"fmt"
	"os"
	"time"
	"todo/internal/commands"
)

func main() {
    args := os.Args;

    if (len(args) <= 1) {
        fmt.Println("No command entered")
        os.Exit(1)
    }

    command := args[1]
    commands.AddTodo(command, time.Now())
    fmt.Println("Done!")
}
