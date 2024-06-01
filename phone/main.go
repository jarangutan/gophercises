package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	phonedb "github.com/jarangutan/gophercises/phone/db"
	_ "github.com/mattn/go-sqlite3"
)

const file string = "phone.db"

const (
	driver     = "sqlite3"
	dataSource = "phone.db"
)

func main() {
	fmt.Println("Resetting phone number table")
	must(phonedb.Reset(driver, dataSource))

	fmt.Println("Creating phone number table")
	must(phonedb.Migrate(driver, dataSource))

	db, err := phonedb.Open(driver, dataSource)
	must(err)
	defer db.Close()

	// fmt.Println("Pinging database")
	// err = db.Ping()
	// must(err)

	fmt.Println("Seeding phone number table")
	must(db.Seed())

	phones, err := db.AllPhones()
	must(err)

	for _, p := range phones {
		fmt.Printf("Working on %+v\n", p)
		number := Normalize(p.Number)
		if number != p.Number {
			fmt.Println("Updating or removing...")
			existing, err := db.FindPhone(number)
			must(err)
			if existing != nil {
				must(db.DeletePhone(p.ID))
			} else {
				p.Number = number
				must(db.UpdatePhone(&p))
			}
		} else {
			fmt.Println("No changes required")
		}
	}

	phones, err = db.AllPhones()
	must(err)
	fmt.Println("Phones in db after normalizing")
	for _, p := range phones {
		fmt.Printf("Entry %+v\n", p)
	}
}

func Normalize(input string) string {
	var b strings.Builder
	for _, ch := range input {
		if unicode.IsDigit(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
