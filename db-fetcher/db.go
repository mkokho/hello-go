package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type db struct {
	conn *sqlx.DB
}

func mustSetupDB() *db {
	db, err := setupMariaDB()
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Fatal("Could not setup MariaDB.")
	}
	return db
}

func setupMariaDB() (*db, error) {
	dsn := fmt.Sprintf("remote:e5e37e4811ae@tcp(%v:%v)/db", "ec2-54-218-82-6.us-west-2.compute.amazonaws.com", "23306")

	conn, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(); err != nil {
		return nil, err
	}

	return &db{conn}, nil
}

type Transaction struct {
	Id        int `db:"id"`
	Member_id int `db:"member_id"`
}

func (db *db) fetch(minId int, limit int) ([]Transaction, error) {
	ts := []Transaction{}
	err := db.conn.Select(&ts, "select id, member_id from pos_transactions where id > ? limit ?", minId, limit)
	if err != nil {
		return nil, err
	}

	return ts, nil
}
