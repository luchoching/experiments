// http://blog.ralch.com/tutorial/golang-subcommands/
package main

import (
	"flag"
	"fmt"
)

func main() {
	username := flag.String("user", "root", "Username for this server")

	flag.Parse()
	fmt.Printf("Your username is %q", *username)
}
