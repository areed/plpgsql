package plpgsql

import "database/sql"

//Void executes a PL/pgSQL function returning nothing.
func Void(db *sql.DB, fn string, args ...interface{}) error {
	_, err := QueryRow(db, fn, args...)
	return err
}

//Int64 executes a PL/pgSQL function returning an int64.
func Int64(db *sql.DB, fn string, args ...interface{}) (int64, error) {
	var i int64
	err := ExecFn(db, fn, args, &i)
	return i, err
}

//String executes a PL/pgSQL function returning a string.
func String(db *sql.DB, fn string, args ...interface{}) (string, error) {
	var s string
	err := ExecFn(db, fn, args, &s)
	return s, err
}
