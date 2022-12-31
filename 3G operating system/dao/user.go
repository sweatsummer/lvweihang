package dao

import "3G_operating_system/model"

// SearchUserByUserName 查找
func SearchUserByUserName(name string) (u model.User, err error) {
	row := DB.QueryRow("select id,name,password from user where name=?", name)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.ID, &u.UserName, &u.Password)
	return
}

// InsertUser 插入//
func InsertUser(u model.User) (err error) {
	_, err = DB.Exec("insert into user(name,password) values (?,?)", u.UserName, u.Password)
	if err != nil {
		return
	}
	return err
}

// InputRecharge 输入
func InputRecharge(number string) (err error) {
	_, err = DB.Exec("insert into user(recharge) values (?)", number)
	if err != nil {
		return
	}
	return err
}
