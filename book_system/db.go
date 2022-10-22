package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() (err error) {
	addr := "root:123456@tcp(127.0.0.1:3306)/practice"
	db, err = sql.Open("mysql", addr)
	if err != nil {
		return err
	}
	//最大连接
	db.SetMaxOpenConns(100)
	//最大空闲
	db.SetMaxIdleConns(16)
	return
}

func QueryAllbook() (*[]Book, error) {
	sqlStr := "select id,title,price from book"
	r, _ := db.Query(sqlStr)
	defer r.Close()
	bookSlice := make([]Book, 0, 10)
	for r.Next() {
		var b Book
		err := r.Scan(&b.ID, &b.Title, &b.Price)
		if err != nil {
			return nil, err
		}
		bookSlice = append(bookSlice, b)
	}
	return &bookSlice, nil
}

func InsertBook(title string, price int64) (err error) {
	sqlStr := "insert into book(title,price) values(?,?)"
	db.Exec(sqlStr, title, price)
	if err != nil {
		fmt.Println("插入失败")
		return
	}
	return
}

func DeleteBook(id int64) (err error) {
	sqlStr := "delete from book where id=?"
	_, err = db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("删除失败")
		return err
	}
	return
}
