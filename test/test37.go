package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./order.sqlite.db")
	if err != nil {
		println("11")
		println(err)
	}
	sql_table := `
    CREATE TABLE IF NOT EXISTS express_order(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        uid INTEGER NOT NULL,
        weight DOUBLE NOT NULL ,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );
    `
	db.Exec(sql_table)
	insertExpressOrder, err := db.Prepare("insert into express_order(uid,weight) values (?,?);")
	res, err := insertExpressOrder.Exec("101011", "2.34")
	if err != nil {
		println(12)
		println(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		println(13)
		println(err)
	}
	println(id)
}
