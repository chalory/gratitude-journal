package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	// "errors"
)

func GetEntry(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	db, err := sql.Open("sqlite3", "./entries.db")
	checkErr(err)
	stmt, err := db.Query("SELECT * FROM entries ORDER BY date DESC;")
	// stmt, err := db.Query("SELECT json_group_array(json_object('id', id, 'entry', entry, 'date', date, 'user_name', user_name)) AS json_result FROM (SELECT * FROM entries ORDER BY id);")
	err = stmt.Err()
	checkErr(err)
	msg := make([]entry, 0)

	for stmt.Next() {
		theEntry := entry{}
		err = stmt.Scan(&theEntry.ID, &theEntry.Entry, &theEntry.Date, &theEntry.User_name)
		checkErr(err)
		msg = append(msg, theEntry)
	}

	err = stmt.Err()
	checkErr(err)
	fmt.Println(msg)
	// db.Exec("SELECT json_group_array(json_object('id', id, 'entry', entry, 'date', date, 'user_name', user_name)) AS json_result FROM (SELECT * FROM entries ORDER BY id);")
	c.IndentedJSON(http.StatusOK, msg)
	stmt.Close()
}
