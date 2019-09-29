package mysqldb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// MySQL DB Driver
	_ "github.com/go-sql-driver/mysql"
)

// Db Connection to be imported
var Db *sql.DB

// ImagesModel ...
type ImagesModel struct {
	TransactionID int32  `json:"transaction_id"`
	ImageID       string `json:"image_id"`
	Scale         int32  `json:"scale"`
	ImageURL      string `json:"image_url"`
}

func init() {
	dbUsername, _ := os.LookupEnv("MYSQL_SERVICE_USERNAME")
	dbPassword, _ := os.LookupEnv("MYSQL_SERVICE_PASSWORD")

	var err error

	dataSourceName := fmt.Sprintf("%s:%s@tcp(mysql-service:3306)/nodemsdb", dbUsername, dbPassword)
	fmt.Println("dataSourceName: ", dataSourceName)
	Db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("\nDb Connection Error: %v\n", err)
	}

	// Ping the DB
	err = Db.Ping()
	if err != nil {
		log.Fatalf("\nDb Ping Error: %v\n", err)
	}

	DBStats := Db.Stats()
	fmt.Println("DBStats: ", DBStats)
}
