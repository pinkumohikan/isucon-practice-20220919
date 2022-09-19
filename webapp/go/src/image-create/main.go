package main

import(
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

func main() {
	seedBuf := make([]byte, 8)
	crand.Read(seedBuf)
	rand.Seed(int64(binary.LittleEndian.Uint64(seedBuf)))

	dsn := fmt.Sprintf("isucon:isucon@tcp(127.0.0.1:3306)/isubata?parseTime=true&loc=Local&charset=utf8mb4&interpolateParams=true")

	log.Printf("Connecting to db: %q", dsn)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Println(err)
	}

	rows, err := db.Query("SELECT name, data FROM image")
	if err != nil {
		return
	}

	for rows.Next() {
		var name string
		var data []byte
		rows.Scan(&name, &data)

		/ 書き込み
		err = ioutil.WriteFile("/home/isucon/isucon-practice-20220919/webapp/public/icons/"+name, data, 0666)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}