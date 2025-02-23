package db

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const dbFile = "db/users.txt"

func CheckDbForUser(name string) bool {

	name = strings.TrimSpace(strings.Title(name))

	db, err := os.Open(dbFile)
	if err != nil {
		fmt.Printf("Error opening DB file: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	reader := bufio.NewReader(db)

	for {
		line, eofErr := reader.ReadString('\n')
		if name == strings.TrimSpace(line) {
			return true
		}
		if eofErr != nil {
			addUserToDb(name)
			return false
		}
	}

}

func addUserToDb(name string) {
	db, err := os.OpenFile(dbFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening DB file: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	_, writeErr := db.WriteString(name + "\n")

	if writeErr != nil {
		fmt.Printf("Error adding user: %v\n", writeErr)
		os.Exit(1)
	}

}
