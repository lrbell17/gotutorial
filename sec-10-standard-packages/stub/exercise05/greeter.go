package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/lrbell17/gotutorial/sec-10-standard-packages/stub/exercise05/db"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("What is your name? ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Printf("Hello %v\n", strings.Title(name))

	isPresent := db.CheckDbForUser(name)
	if isPresent {
		fmt.Println("Welcome back!")
	} else {
		fmt.Println("Welcome first time user!")
	}

	t := time.Now()
	fmt.Printf("The date is %v and the time is %v\n", t.Format(time.RFC1123)[:16], t.Format(time.Kitchen))

}
