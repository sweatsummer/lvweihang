package dao

import (
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

func InitDB() {
	//访问本地数据库//
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/message_borad") //因无法修改数据库名称，暂借用留言版项目的数据库//
	if err != nil {
		log.Fatalf("connect mysql error :%v", err)
	}
	DB = db
	//检验连接//
	fmt.Println(db.Ping())
}
