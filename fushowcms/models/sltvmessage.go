package models

import (
	_ "github.com/go-sql-driver/mysql"
)

func (sltvm *SLTVMessage) AddSLTVMessage() bool {
	rows, err := Engine.Insert(sltvm)
	if rows > 0 && err == nil {
		return true
	}
	return false
}
