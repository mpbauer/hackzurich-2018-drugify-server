package models

import (
	"gopkg.in/mgo.v2"
	"log"
	"github.com/mpbauer/zhaw-issue-tracker-server/errorhandling"
)

const (
	COLLECTION = "drug"
)

type DB struct {
	*mgo.Database
}

func NewDB(url string, dbName string) (*DB, error){

	session, err := mgo.Dial(url)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	db := session.DB(dbName)
	return &DB{db}, nil
}

func checkIfObjectNotFound(err *error) {
	if *err == mgo.ErrNotFound{
		*err = errorhandling.NewErrNotFound("object not found")
	}
}