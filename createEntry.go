package main

import (
	"database/sql"
	"net/http"
	"time"

	// "encoding/json"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	// "errors"
)

func createEntry(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
	dt := time.Now()
	var newEntry entry
	if err := c.BindJSON(&newEntry); err != nil {
		return
	}
	// tmp := []byte(newEntry)
	// m := map[string]interface{}{}
	// if err := json.Unmarshal(tmp, &m); err != nil {
	// 	panic(err)
	// }
	db, err := sql.Open("sqlite3", "./entries.db")
	checkErr(err)
	// defer close
	stmt, _ := db.Prepare("INSERT INTO entries (id, entry, date, user_name) VALUES (?, ?, ?, ?)")
	stmt.Exec(nil, newEntry.Entry, dt.Format("01-02-2006 15:04:05"), newEntry.User_name)
	// db.Exec("INSERT INTO entries SELECT json_extract(value, '$.id'), json_extract(value, '$.entry'), json_extract(value, '$.date'), json_extract(value, '$.user_name') FROM json_each(readfile('body.json'));")
	stmt.Close()
	// entries = append(entries, newEntry)
	c.IndentedJSON(http.StatusCreated, newEntry)
}
