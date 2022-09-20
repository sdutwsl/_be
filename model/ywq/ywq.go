package ywq

import (
	"_be/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type YWQ struct {
	id        uint
	date_time string
}

func GetAllRecords() []YWQ {
	var ret []YWQ
	var sql_source = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True", config.MYSQL_USER, config.MYSQL_PASS, config.MYSQL_HOST, config.MYSQL_DB)
	var sql_query = "SELECT * FROM ywq"
	db, err := sql.Open("mysql", sql_source)
	if err != nil {
		return ret
	}

	rows, err := db.Query(sql_query)

	if err != nil {
		return ret
	}

	for rows.Next() {
		var y YWQ
		rows.Scan(&y.id, &y.date_time)
		ret = append(ret, y)
		fmt.Println(y)
	}

	return ret
}
