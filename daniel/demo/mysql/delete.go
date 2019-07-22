package main

import (
	"./logic"
	"fmt"
)

func main()  {
	db, err := logic.NewMysql()
	logic.CheckErr(err)
	defer db.Close()

	sql := "delete from t_renewal_v2_deal where Fdeal_id=?"
	stmt, err := db.Prepare(sql)
	logic.CheckErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(73628321)
	logic.CheckErr(err)

	affect, err := res.RowsAffected()
	logic.CheckErr(err)
	fmt.Println("受影响的行数:", affect)
}
