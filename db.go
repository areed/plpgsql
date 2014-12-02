//Package plpgsql provides some utilities for running PL/pgSQL functions.
package plpgsql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

//Open returns a connection to a postgres database.
func Open(conn string) (*sql.DB, error) {
	return sql.Open("postgres", conn)
}

//MustOpen returns a connection to a postgres database or panics.
func MustOpen(conn string) *sql.DB {
	db, err := Open(conn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

//QueryRow is similar to sql.QueryRow but ignores sql.ErrNoRows and returns errors raised within a PL/pgSQL function.
func QueryRow(db *sql.DB, fn string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s%s;", fn, paramSql[len(args)]), args...)
	if err != nil {
		return nil, err
	}
	rows.Next()
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return rows, nil
}

//ExecFn runs a PL/pgSQL function.
func ExecFn(db *sql.DB, fn string, args []interface{}, dests ...interface{}) error {
	rows, err := QueryRow(db, fn, args...)
	if err != nil {
		return err
	}
	return ScanRow(rows, dests...)
}

//ScanRow scans the return values of a PL/pgSQL function into the provided destination arguments.
func ScanRow(rows *sql.Rows, dests ...interface{}) error {
	err := rows.Scan(dests...)
	if err != nil {
		return err
	}
	return nil
}

var paramSql = map[int]string{
	0:  "()",
	1:  "($1)",
	2:  "($1, $2)",
	3:  "($1, $2, $3)",
	4:  "($1, $2, $3, $4)",
	5:  "($1, $2, $3, $4, $5)",
	6:  "($1, $2, $3, $4, $5, $6)",
	7:  "($1, $2, $3, $4, $5, $6, $7)",
	8:  "($1, $2, $3, $4, $5, $6, $7, $8)",
	9:  "($1, $2, $3, $4, $5, $6, $7, $8, $9)",
	10: "($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
	11: "($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
	12: "($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)",
}
