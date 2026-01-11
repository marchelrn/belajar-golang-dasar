package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	query := "INSERT INTO customer(id, name) VALUES('admin', null)"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	query := "SELECT id, name FROM customer where name IS NULL"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success select joko")
	defer rows.Close()

	for rows.Next() {
		var id string
		var name sql.NullString
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("ID:", id)
		if name.Valid {
			fmt.Println("Name:", name.String)
		} else {
			fmt.Println("Name: NULL (Data Kosong)")
		}
	}
}

//func TestScanNull(t *testing.T) {
//	db := GetConnection()
//	defer db.Close()
//	ctx := context.Background()
//
//	_, err := db.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS customer (id VARCHAR(50) PRIMARY KEY, name VARCHAR(50) NULL)")
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	_, err = db.ExecContext(ctx, "INSERT INTO customer(id, name) VALUES('dondi', NULL)")
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	query := "SELECT id, name FROM customer WHERE name IS NULL"
//	rows, err := db.QueryContext(ctx, query)
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer rows.Close()
//
//	for rows.Next() {
//		var id string
//		var name sql.NullString
//
//		err := rows.Scan(&id, &name)
//		if err != nil {
//			t.Fatal(err)
//		}
//
//		fmt.Println("ID:", id)
//		//if name.Valid {
//		//	fmt.Println("Name:", name.String)
//		//} else {
//		//	fmt.Println("Name: NULL (Data Kosong)")
//		//}
//		fmt.Println("Name:", name)
//	}
//}

func TestQueryComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	query := "SELECT id, name,email,balance,birth_date,rating,married,created_at FROM customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance sql.NullInt32
		var rating sql.NullFloat64
		var birthDate, createdAt sql.NullTime
		var married sql.NullBool
		err := rows.Scan(&id, &name, &email, &balance, &birthDate, &rating, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("---------------------")
		fmt.Println("ID:", id)
		fmt.Println("Name:", name)
		if email.Valid {
			fmt.Println("Email:", email.String)
		}
		if balance.Valid {
			fmt.Println("Balance:", balance.Int32)
		}
		if rating.Valid {
			fmt.Println("Rating:", rating.Float64)
		}
		if birthDate.Valid {
			fmt.Println("Birth Date:", birthDate.Time)
		}
		if married.Valid {
			fmt.Println("Married:", married.Bool)
		}
		fmt.Println("---------------------")
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	username := "admin'; --"
	password := "admin123"

	query := "SELECT username FROM users WHERE username='" + username + "' AND password='" + password + "' LIMIT 1"
	fmt.Println(query)
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Login Success:", username)
	} else {
		fmt.Println("Login Failed")
	}
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	username := "admin' --"
	password := "admin123"

	query := "SELECT username FROM users WHERE username= $1 AND password= $2 LIMIT 1"
	fmt.Println(query)
	rows, err := db.QueryContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Login Success:", username)
	} else {
		t.Fatal("Login Failed, username or password incorrect")
	}
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	username := "doni'; DROP TABLE users; --"
	password := "doni123"

	query := "INSERT INTO users(username, password) VALUES($1, $2)"
	_, err := db.ExecContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new customer")
}
