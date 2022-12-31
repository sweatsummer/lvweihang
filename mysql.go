package main

import "fmt"

// 数据库（a）：创建表描述//
// （b）不知道，没听过//
func main() {
	fmt.Println("mysql -u root -p\n\ncreate database company;" +
		"\n\nuse company;\n\ncreate table 'company_pay' (\n'id' int(11) default 0,\n" +
		"'name' varchar(20) default '',\n'work' varchar(20),\n" +
		"'achievement' int(11) default 0,\n); ")
}
