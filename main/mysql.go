package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "app",
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	regions, err := regionByName("湖南省")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("regions found: %v\n", regions)
}

type Region struct {
	ID    int64
	Pid   int64
	Name  string
	Level int64
}

func regionByName(name string) ([]Region, error) {
	var regions []Region

	rows, err := db.Query("SELECT * FROM region WHERE name = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()

	for rows.Next() {
		var region Region
		if err := rows.Scan(&region.ID, &region.Pid, &region.Name, &region.Level); err != nil {
			return nil, fmt.Errorf("regionByName %q: %v", name, err)
		}
		regions = append(regions, region)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("regionByName %q: %v", name, err)
	}
	return regions, nil
}
