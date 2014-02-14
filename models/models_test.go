package models

import (
	"fmt"
	"testing"

	"github.com/lunny/xorm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	var err error
	orm, err = xorm.NewEngine("sqlite3", "./test.db")
	if err != nil {
		fmt.Println(err)
	}

	orm.ShowSQL = true
	orm.ShowDebug = true

	err = orm.Sync(&User{}, &Repo{})
	if err != nil {
		fmt.Println(err)
	}

	root = "test"
}

func TestCreateRepository(t *testing.T) {
	user := User{Id: 1, Type: Individual}
	_, err := CreateRepository(&user, "test")
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteRepository(t *testing.T) {
	user := User{Id: 1, Type: Individual}
	err := DeleteRepository(&user, "test")
	if err != nil {
		t.Error(err)
	}
}
