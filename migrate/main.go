package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	my "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	// addr := fmt.Sprintf("mysql://%v:%v@tcp(%v:%v)/%v", "root", "root123", "127.0.0.1", "3306", "test")
	addr := fmt.Sprintf("mysql://root:root@tcp(%v:%v)/public", "127.0.0.1", 1)

	p := &my.Mysql{}
	_, err := p.Open(addr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("ok")
}
