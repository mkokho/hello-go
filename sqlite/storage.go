package main

import (
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type object struct {
	Id int `json:"id"`
	Body string `json:"body"`
}

type storage interface {
	addObject(o object) error
	getObject(id int) (object, error)
}

type memStorage struct {
	objects map[int]object
}

func newMemStorage() *memStorage {
	return &memStorage{
		objects: map[int]object{},
	}
}

func (s *memStorage) addObject(o object) error {
	s.objects[o.Id] = o
	return nil
}

func (s *memStorage) getObject(id int) (object, error) {
	o, ok := s.objects[id]
	if !ok {
		return object{}, fmt.Errorf("object %s not found", id)
	}
	return o, nil
}

type sqlStorage struct {
	db *sql.DB
}

func newSqlStorage(c *config) *sqlStorage {
	db, err := sql.Open("sqlite3", c.dbFile)
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
	create table if not exists objects  (id integer not null primary key, body text);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}

	return &sqlStorage{
		db: db,
	}
}


func (s *sqlStorage) addObject(o object) error {
	_, err := s.db.Exec("insert into objects (id, body) values (?, ?)", o.Id, o.Body)
	return err
}

func (s *sqlStorage) getObject(id int) (object, error) {
	var obj object
	rows, err := s.db.Query("select id, body from objects where id = ?", id)
	if err != nil {
		return obj, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Body)
		return obj, nil
	}

	return obj, fmt.Errorf("Not found: %v", id)
}
