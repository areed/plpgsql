package plpgsql

import (
	"os"
	"testing"
)

var DB = MustOpen(os.Getenv("PLPGSQL_CONNECTION"))

func TestOpen(t *testing.T) {
	db, err := Open(os.Getenv("PLPGSQL_CONNECTION"))
	if err != nil {
		t.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		t.Error(err)
	}
	db.Close()
}

func TestMustOpen(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatal(r)
		}
	}()
	db := MustOpen(os.Getenv("PLPGSQL_CONNECTION"))
	if err := db.Ping(); err != nil {
		t.Error(err)
	}
	db.Close()
}

var return_void = `
CREATE OR REPLACE FUNCTION return_void() RETURNS void AS $$
BEGIN
	RETURN;
END;
$$ LANGUAGE plpgsql;`

func TestVoid(t *testing.T) {
	if _, err := DB.Exec(return_void); err != nil {
		t.Fatal(err)
	}
	err := Void(DB, "return_void")
	if err != nil {
		t.Fatal(err)
	}
}

var return_twelve = `
CREATE OR REPLACE FUNCTION return_twelve() RETURNS int AS $$
BEGIN
	RETURN 12;
END;
$$ LANGUAGE plpgsql;`

func TestInt64(t *testing.T) {
	_, err := DB.Exec(return_twelve)
	if err != nil {
		t.Fatal(err)
	}
	i, pgerr := Int64(DB, "return_twelve")
	if pgerr != nil {
		t.Errorf("%t %+v", err)
	}
	if i != 12 {
		t.Errorf("expect %d to equal 12", i)
	}
}

var return_hello = `
CREATE OR REPLACE FUNCTION return_hello() RETURNS text AS $$
BEGIN
	RETURN 'hello';
END;
$$ LANGUAGE plpgsql;`

func TestString(t *testing.T) {
	if _, err := DB.Exec(return_hello); err != nil {
		t.Fatal(err)
	}
	s, err := String(DB, "return_hello")
	if err != nil {
		t.Fatal(err)
	}
	if s != "hello" {
		t.Errorf("expected %q, got %q", "hello", s)
	}
}
