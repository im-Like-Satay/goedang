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
				fmt.Println("Format salah. Gunakan: tambah [id] [nama] [stok]")
				continue
			}
			id := args[1]
			nama := args[2]
			var stok int

			_, err := fmt.Sscanf(args[3], "%d", &stok)
			if err != nil {
				fmt.Println("Eror: Stok harus angka.")
				continue
			}

			err = TambahBarang(db, id, nama, stok)
			if err != nil {
				fmt.Println("Eror: Gagal menyimpan data.")
			} else {
				fmt.Println("Sukses menambahkan barang.")
			}

		case "delete":
			id := args[1]
			err := HapusBarang(db, id)
			if err != nil {
				fmt.Printf("error : %v \n", err)
			}

		default:
			fmt.Printf("Perintah '%s' tidak dikenal.\n", perintah)
		}

		if err := scanner.Err(); err != nil {
			log.Fatalf("Terjadi kesalahan pada input terminal: %v", err)
		}
	}
}
