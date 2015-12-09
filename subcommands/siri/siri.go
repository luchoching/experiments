// https://gist.github.com/svett/a6f02026270b443d5e46
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	askCommand := flag.NewFlagSet("ask", flag.ExitOnError)
	questionFlag := askCommand.String("question", "", "Question that you are asking for.")

	if len(os.Args) == 1 {
		fmt.Println("usage: siri <command> [<args.]")
		fmt.Println("The most commonly used git commands are:")
		fmt.Println("ask Ask questions")
	}

	switch os.Args[1] {
	case "ask":
		askCommand.Parse(os.Args[2:])
	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		os.Exit(2)
	}

	if askCommand.Parsed() {
		if *questionFlag == "" {
			fmt.Println("Please supply the question using -question option.")
			return
		}
		fmt.Printf("You asked: %q\n", *questionFlag)
	}

}
