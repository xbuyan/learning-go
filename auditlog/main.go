package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  auditlog add <actor> <action> <document>")
		fmt.Println("  auditlog verify")
		fmt.Println("  auditlog show")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 5 {
			fmt.Println("Usage: auditlog add <actor> <action> <document>")
			return
		}
		actor := os.Args[2]
		action := os.Args[3]
		document := os.Args[4]

		documentHash, err := hashFile(document)
		if err != nil {
			fmt.Println("Error hashing document:", err)
			return
		}

		err = appendEntry(actor, action, document, documentHash)
		if err != nil {
			fmt.Println("Error saving entry:", err)
			return
		}
		fmt.Println("Entry added successfully.")

	case "verify":
		err := verifyChain()
		if err != nil {
			fmt.Println(err)
			return
		}

	case "show":
		entries, err := loadLog()
		if err != nil {
			fmt.Println("Error loading log:", err)
			return
		}
		if len(entries) == 0 {
			fmt.Println("Log is empty.")
			return
		}
		for _, entry := range entries {
			fmt.Println("---")
			fmt.Println("ID:           ", entry.ID)
			fmt.Println("Timestamp:    ", entry.Timestamp)
			fmt.Println("Actor:        ", entry.Actor)
			fmt.Println("Action:       ", entry.Action)
			fmt.Println("Document:     ", entry.Document)
			fmt.Println("Document Hash:", entry.DocumentHash)
			fmt.Println("Previous Hash:", entry.PreviousHash)
			fmt.Println("Entry Hash:   ", entry.EntryHash)
		}

	default:
		fmt.Println("Unknown command:", command)
	}
}
