package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var dbconn *sql.DB

type TaskInfo struct {
	Task_id   int    `json:"task_id"`
	User_name string `json:"user_name"`
	Task_user string `json:"task_user"`
	Bonus     int    `json:"bonus"`
	Status    string `json:"status"`
	Remark    string `json:"remark"`
	Comment   string `json:"comment"`
}

func init() {
	mysql_server := os.Getenv("MYSQL_SERVER")
	if mysql_server == "" {
		mysql_server = "0.0.0.0:3306"
	}
	db, err := sql.Open("mysql", "root:root@tcp("+mysql_server+")/tokentask?charset=utf8")
	if err != nil {
		log.Panic("failed to open mysql ", err)
	}
	if err = db.Ping(); err != nil {
		log.Panic("failed to ping mysql ", err)
	}
	dbconn = db

}

func userlogin(user, pass string) bool {
	row, err := dbconn.Query("select 1 from t_user where name=? and pass=?", user, pass)
	if err != nil {
		fmt.Println("user not exists or password err ", err, user)
		return false
	}
	return row.Next()
}

func task_query() []TaskInfo {
	var tasks []TaskInfo

	rows, err := dbconn.Query("select task_id, user_name, task_user, bonus, status, remark, ifnull(comment,'none') from t_task order by task_id")
	if err != nil {
		fmt.Println("failed for query sql:", err)
		return tasks
	}

	for rows.Next() {
		var task TaskInfo
		err := rows.Scan(&task.Task_id, &task.User_name, &task.Task_user, &task.Bonus, &task.Status, &task.Remark, &task.Comment)
		if err != nil {
			fmt.Println("failed to Scan result set", err)
			break
		}
		tasks = append(tasks, task)
	}

	return tasks
}
