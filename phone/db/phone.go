package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Phone struct {
	ID     int
	Number string
}

type DB struct {
	db *sql.DB
}

func Open(driverName, dataSourceName string) (*DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func (db *DB) Ping() error {
	err := db.db.Ping()
	return err
}

func (db *DB) Close() error {
	return db.db.Close()
}

func (db *DB) Seed() error {
	data := []string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}

	for _, p := range data {
		if _, err := insertPhone(db.db, p); err != nil {
			return err
		}
	}
	return nil
}

func insertPhone(db *sql.DB, phone string) (int, error) {
	insert := "INSERT INTO phone_numbers(value) VALUES($1) RETURNING id;"
	var id int
	err := db.QueryRow(insert, phone).Scan(&id)
	if err != nil {
		return -1, nil
	}
	return id, nil
}

func (db *DB) AllPhones() ([]Phone, error) {
	rows, err := db.db.Query("SELECT * FROM phone_numbers")
	if err != nil {
		return nil, err
	}

	var ret []Phone
	for rows.Next() {
		var p Phone
		if err := rows.Scan(&p.ID, &p.Number); err != nil {
			return nil, err
		}
		ret = append(ret, p)
	}
	return ret, nil
}

func (db *DB) FindPhone(number string) (*Phone, error) {
	row := db.db.QueryRow("SELECT id, value FROM phone_numbers WHERE value=?", number)

	var p Phone
	err := row.Scan(&p.ID, &p.Number)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &p, nil
}

func (db *DB) UpdatePhone(p *Phone) error {
	fmt.Println(p)
	_, err := db.db.Exec("UPDATE phone_numbers SET value=$2 WHERE id=$1 ", p.Number, p.ID)
	return err
}

func (db *DB) DeletePhone(id int) error {
	_, err := db.db.Exec("DELETE FROM phone_numbers WHERE id=$1", id)
	return err
}

func Migrate(driverName, dataSourceName string) error {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return err
	}
	err = createPhoneNumbersTable(db)
	if err != nil {
		return err
	}
	return db.Close()
}

func createPhoneNumbersTable(db *sql.DB) error {
	const create string = `
	 CREATE TABLE IF NOT EXISTS phone_numbers (
	  id INTEGER PRIMARY KEY,
	  value TEXT
	 );`

	_, err := db.Exec(create)
	return err
}

func Reset(driverName, dataSourceName string) error {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return err
	}
	err = dropPhoneNumbersTable(db)
	if err != nil {
		return err
	}
	return db.Close()
}

func dropPhoneNumbersTable(db *sql.DB) error {
	const drop string = `DROP TABLE IF EXISTS phone_numbers;`
	_, err := db.Exec(drop)
	return err
}
