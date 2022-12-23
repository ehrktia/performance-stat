package pgsql

import (
	"encoding/base64"
	"fmt"
)

const (
	tablename = "public.test"
)

var (
	ErrInvalidInputID = fmt.Errorf("error:%s", "invalid input id received")
)

func isNotValidID(id int) bool {
	return id <= 0
}

func buildStatement() string {
	return fmt.Sprintf("SELECT name from %s where id=$1;", tablename)
}

func (s *store) GetByID(id int) ([]byte, error) {
	if isNotValidID(id) {
		return nil, ErrInvalidInputID
	}
	statement := buildStatement()
	rows, err := s.connection.Query(statement, id)
	if err != nil {
		return nil, err
	}
	var name string
	for rows.Next() {
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}

	}

	return []byte(name), nil
}

func buildInsertQuery() string {
	return fmt.Sprintf("insert into %s values($1,$2);", tablename)

}

func (s *store) PutData() error {
	stmt := buildInsertQuery()
	sourceSeed := "abcdefghijklmnopqrstuvwxyz1234567890"
	name := base64.StdEncoding.EncodeToString([]byte(sourceSeed))
	for i := 1; i <= 10000; i++ {
		if _, err := s.connection.Exec(stmt, i, name); err != nil {
			return err
		}
	}
	return nil
}

func buildGetAllStatement() string {
	return fmt.Sprintf("select * from %s;", tablename)

}

func (s *store) GetAll() error {
	stmt := buildGetAllStatement()
	_, err := s.connection.Exec(stmt)
	if err != nil {
		return err
	}
	return nil

}
