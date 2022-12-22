package pgsql

import (
	"fmt"
)

var (
	ErrInvalidInputID = fmt.Errorf("error:%s", "invalid input id received")
)

func isNotValidID(id int) bool {
	return id > 0
}

func buildStatement() string {
	return fmt.Sprintf("SELECT name from %s where id=$1;", "public.test")
}

func (s *store) GetByID(id int) ([]byte, error) {
	if isNotValidID(id) {
		return nil, ErrInvalidInputID
	}
	statement := buildStatement()
	rows, err := s.connection.Query(statement)
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
