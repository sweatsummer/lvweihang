// lv0 mysql 语句略//
// golang-mysql 操作(增，删，改，查)//
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type User struct {
	id      int
	name    string
	math    int
	english int
}

func main() {
	//连接mysql数据库//
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/students")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	var u []User
	var tmp User
	//检索students.score数据//
	rows, err := db.Query("select * from score ")
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		err := rows.Scan(&tmp.id, &tmp.name, &tmp.math, &tmp.english)
		if err != nil {
			log.Println(err)
			return
		}
		u = append(u, tmp)
	}
	fmt.Println(u)
	//插入小红数据//
	result2, err := db.Exec("insert into score (name,math,english) values (?,?,?)", "小红", 130, 145)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(result2.LastInsertId())
	fmt.Println(result2.RowsAffected())
	//更新数据（将id为2的数学改为140）
	result1, err := db.Exec("update score set math=? where id=?", 140, 2)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(result1.RowsAffected())
	fmt.Println(result1.RowsAffected())

	////删除数据//
	result, err := db.Exec("delete from score where id=?", 7)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}
