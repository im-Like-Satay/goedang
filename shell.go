package main

import (
	"bufio"
	"fmt"
	"goedang/db"
	"log"
	"os"
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
			err := LihatBarang(db)
			if err != nil {
				fmt.Printf("Eror: %v\n", err)
			}

		case "add":
			if len(args) < 4 {
				fmt.Println("Wrong fromat. Use: add [id] [nama] [stok]")
				continue
			}
			id := args[1]
			nama := args[2]
			var stok int

			_, err := fmt.Sscanf(args[3], "%d", &stok)
			if err != nil {
				fmt.Println("Eror: Stock must be number.")
				continue
			}

			err = TambahBarang(db, id, nama, stok)
			if err != nil {
				fmt.Println("Eror: Failed save data.")
			} else {
				fmt.Println("Success added item.")
			}

		case "delete":
			id := args[1]
			err := HapusBarang(db, id)
			if err != nil {
				fmt.Printf("error : %v \n", err)
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
