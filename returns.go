package plpgsql

import (
	"database/sql"

	"github.com/lib/pq"
)

//Void executes a PL/pgSQL function returning nothing.
func Void(db *sql.DB, fn string, args ...interface{}) *pq.Error {
	_, err := QueryRow(db, fn, args...)
	return err
}

//Int64 executes a PL/pgSQL function returning an int64.
func Int64(db *sql.DB, fn string, args ...interface{}) (int64, *pq.Error) {
	var i int64
	err := ExecFn(db, fn, args, &i)
	return i, err
}

//String executes a PL/pgSQL function returning a string.
func String(db *sql.DB, fn string, args ...interface{}) (string, *pq.Error) {
	var s string
	err := ExecFn(db, fn, args, &s)
	return s, err
}
