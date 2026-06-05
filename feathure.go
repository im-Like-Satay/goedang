package main

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func TambahBarang(db *sql.DB, id string, nama string, stok int) error {
	query := "INSERT INTO barang (id, nama, stok) VALUES (?, ?, ?)"
	_, err := db.Exec(query, id, nama, stok)
	return err
}

func LihatBarang(db *sql.DB) error {
	rows, err := db.Query("SELECT id, nama, stok FROM barang")
	if err != nil {
		return err
	}
	defer rows.Close()

	table := tablewriter.NewWriter(os.Stdout)

	table.Header([]string{"ID", "NAME", "STOCK"})

	adaData := false
	for rows.Next() {
		adaData = true
		var id, name string
		var stock int

		err = rows.Scan(&id, &name, &stock)
		if err != nil {
			return err
		}

		table.Append(id, name, strconv.Itoa(stock))
	}

	if !adaData {
		fmt.Println("(Wirehouse is emty)")
	}
	fmt.Println()

	table.Render()

	return nil
}

func HapusBarang(db *sql.DB, id string) error {
	sql := "DELETE FROM barang WHERE id = ?"
	_, err := db.Exec(sql, id)
	return err
}

func ClearTerminal() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
