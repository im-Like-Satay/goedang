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

func AddItems(db *sql.DB, name string, stock int, price int) error {
	query := "INSERT INTO items (name, stock, price) VALUES (?, ?, ?)"
	_, err := db.Exec(query, name, stock, price)
	return err
}

func ShowItems(db *sql.DB) error {
	rows, err := db.Query("SELECT id, name, stock, price, time_stamp FROM items")
	if err != nil {
		return err
	}
	defer rows.Close()

	table := tablewriter.NewWriter(os.Stdout)

	table.Header([]string{"ID", "NAME", "STOCK", "PRICE", "TIME ADD"})

	adaData := false
	for rows.Next() {
		adaData = true
		var id, name string
		var stock int
		var price int
		var time_stamp string

		err = rows.Scan(&id, &name, &stock, &price, &time_stamp)
		if err != nil {
			return err
		}

		table.Append(id, name, strconv.Itoa(stock), strconv.Itoa(price), time_stamp)
	}

	if !adaData {
		fmt.Println("(db is emty)")
	}
	fmt.Println()

	table.Render()

	return nil
}

func DeleteItems(db *sql.DB, name string) error {
	sql := "DELETE FROM items WHERE name = ?"
	_, err := db.Exec(sql, name)
	return err
}

func UpdateDB(db *sql.DB, name string, stock int, price int, NewName string) error {
	sql := "UPDATE items SET name = ?, stock = ?, price = ? WHERE name = ?"
	_, err := db.Exec(sql, name, stock, price, NewName)
	if err != nil {
		return err
	}

	return nil
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
