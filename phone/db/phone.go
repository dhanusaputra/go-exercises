package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//Phone represents the phone_numbers table in the DB
type Phone struct {
	ID     int
	Number string
}

//Open ...
func Open(driverName, dataSource string) (*DB, error) {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

//DB ...
type DB struct {
	db *sql.DB
}

//Close ...
func (db *DB) Close() error {
	return db.db.Close()
}

//Seed ...
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
	for _, number := range data {
		if _, err := insertPhone(db.db, number); err != nil {
			return err
		}
	}
	return nil
}

func insertPhone(db *sql.DB, phone string) (int, error) {
	var id int
	statement := `
	INSERT INTO phone_numbers(value) VALUES($1) RETURNING id
	`
	err := db.QueryRow(statement, phone).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

//AllPhones ...
func (db *DB) AllPhones() ([]Phone, error) {
	rows, err := db.db.Query("SELECT id, value from phone_numbers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []Phone
	for rows.Next() {
		var p Phone
		if err := rows.Scan(&p.ID, &p.Number); err != nil {
			return nil, err
		}
		ret = append(ret, p)
	}
	if err := rows.Err(); err != nil {
		return ret, nil
	}
	return ret, nil
}

//FindPhone ...
func (db *DB) FindPhone(number string) (*Phone, error) {
	var p Phone
	err := db.db.QueryRow("SELECT * FROM phone_numbers WHERE value =$1", number).Scan(&p.ID, &p.Number)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

//UpdatePhone ...
func (db *DB) UpdatePhone(p *Phone) error {
	statement := `UPDATE phone_numbers SET value =$2 WHERE id=$1`
	_, err := db.db.Exec(statement, p.ID, p.Number)
	return err
}

//DeletePhone ...
func (db *DB) DeletePhone(id int) error {
	statement := `DELETE FROM phone_numbers WHERE id=$1`
	_, err := db.db.Exec(statement, id)
	return err
}

//Migrate ...
func Migrate(driverName, dataSource string) error {
	db, err := sql.Open(driverName, dataSource)
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
	statement := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS phone_numbers (
		id SERIAL,
		value VARCHAR(255)
	)`)
	_, err := db.Exec(statement)
	return err
}

//Reset ...
func Reset(driverName, dataSource, dbName string) error {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return err
	}
	err = resetDB(db, dbName)
	if err != nil {
		return err
	}
	return db.Close()
}

func resetDB(db *sql.DB, name string) error {
	_, err := db.Exec("DROP DATABASE IF EXISTS " + name)
	if err != nil {
		return err
	}
	return createDB(db, name)
}

func createDB(db *sql.DB, name string) error {
	_, err := db.Exec("CREATE DATABASE " + name)
	return err
}
