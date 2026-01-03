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
