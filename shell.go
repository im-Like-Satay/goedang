package main

import (
	"bufio"
	"fmt"
	"goedang/db"
	"log"
	"os"
	"strconv"
	"strings"
)

func Shell() {
	db := db.InitDB()
	defer db.Close()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("goedang> ")

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			continue
		}

		args := strings.Fields(input)
		perintah := args[0]

		switch perintah {
		case "exit":
			return

		case "show":
			err := ShowItems(db)
			if err != nil {
				fmt.Printf("Eror:\n %v \n", err)
			}

		case "add":
			if len(args) < 4 {
				fmt.Println("Wrong fromat. Use: add [name] [stok] [price]")
				continue
			}
			name := args[1]
			stock, ErrStock := strconv.Atoi(args[2])
			price, ErrPrice := strconv.Atoi(args[3])

			if ErrStock != nil || ErrPrice != nil {
				fmt.Println("Eror: Stock must be number.")
				continue
			}

			err := AddItems(db, name, stock, price)
			if err != nil {
				fmt.Println("Eror: Failed save data.")
			} else {
				fmt.Println("Success added item.")
			}

		case "update":
			name := args[1]
			stock, _ := strconv.Atoi(args[2])
			price, _ := strconv.Atoi(args[3])
			newName := args[4]

			UpdateDB(db, name, stock, price, newName)
			fmt.Println("Success update DB")

		case "delete":
			name := args[1]
			err := DeleteItems(db, name)
			if err != nil {
				fmt.Printf("Error :\n %v \n", err)
			}

		case "clear":
			ClearTerminal()

		default:
			fmt.Printf("Unknown command `%s'.\n", perintah)
		}

		if err := scanner.Err(); err != nil {
			log.Fatalf("Terjadi kesalahan pada input terminal: %v", err)
		}
	}
}
