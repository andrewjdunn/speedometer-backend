package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"hciware.com/speedometer/record"
)

func SpeedRecords() ([]record.Record, error) {
	var db *sql.DB = openDatabase()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM speed_record where timestamp >= DATE_SUB(NOW(), INTERVAL 5 DAY)")

	var records []record.Record
	if err != nil {
		return nil, fmt.Errorf("SpeedRecords %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var rec record.Record
		var pingOkNumber string
		if err := rows.Scan(&rec.ID, &rec.TimeStamp, &rec.Latency, &rec.UploadSpeed, &rec.DownloadSpeed, &rec.Distance, &pingOkNumber); err != nil {
			return nil, fmt.Errorf("SpeedRecords %v", err)
		}
		rec.PingOk = pingOkNumber == "1"
		records = append(records, rec)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("SpeedRecords %v", err)
	}
	return records, nil
}

func StoreRecord(record record.Record) (int64, error) {
	var db *sql.DB = openDatabase()
	defer db.Close()

	result, err := db.Exec("INSERT INTO speed_record (timestamp, latency, uploadspeed, downloadspeed, distance, pingok) VALUES (?, ?, ?, ?, ?, ?)",
		record.TimeStamp, record.Latency.Milliseconds(), record.UploadSpeed, record.DownloadSpeed, record.Distance, record.PingOk)
	if err != nil {
		return 0, fmt.Errorf("addRecord: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addRecord: %v", err)
	}
	return id, nil
}

func openDatabase() *sql.DB {
	var db *sql.DB

	cfg := mysql.Config{
		User:      os.Getenv("DBUSER"),
		Passwd:    os.Getenv("DBPASS"),
		Net:       "tcp",
		Addr:      "127.0.0.1:3306",
		DBName:    "speedometer",
		ParseTime: true,
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

	return db
}
