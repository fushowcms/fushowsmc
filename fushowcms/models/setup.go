package models

import (
	_ "github.com/go-sql-driver/mysql"
)

func (su *SetUp) AddSetUp() bool {
	rows, err := Engine.Insert(su)
	if rows > 0 && err == nil {
		return true
	}
	return false
}
