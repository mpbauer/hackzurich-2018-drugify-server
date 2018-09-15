package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type DrugHistory struct {
	ID         string    `bson:"_id,omitempty" json:"id"`
	Username   string    `bson:"username" json:"username"`
	Title      string    `bson:"title" json:"title"`
	AuthHolder string    `bson:"authHolder" json:"authHolder"`
	AtcCode    string    `bson:"atcCode" json:"atcCode"`
	Substances string    `bson:"substances" json:"substances"`
	AuthNrs    string    `bson:"authNrs" json:"authNrs"`
	CreatedAt  time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time `bson:"updatedAt" json:"updatedAt"`
	FromDate   string    `bson:"fromDate" json:"fromDate"`
	ToDate     string    `bson:"toDate" json:"toDate"`
}

// Get all drug history items
func (db *DB) GetFullDrugHistory(userId string) ([]DrugHistory, error) {
	query := bson.M{"username": userId}
	var history []DrugHistory

	err := db.C("drugHistory").Find(query).All(&history)

	// Find({}).All() returns nil if no entities were found so the value is replaced by an empty project slice
	if history == nil {
		history = []DrugHistory{}
	}

	return history, err
}

// Create a new drug history item
func (db *DB) InsertDrugHistoryItem(historyItem DrugHistory) error {
	historyItem.CreatedAt = time.Now()
	err := db.C("drugHistory").Insert(historyItem)
	return err
}
