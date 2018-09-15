package models

import "gopkg.in/mgo.v2/bson"

type Drug struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	Title      string `bson:"title" json:"title"`
	AuthHolder string `bson:"authHolder" json:"authHolder"`
	AtcCode    string `bson:"atcCode" json:"atcCode"`
	Substances string `bson:"substances" json:"substances"`
	AuthNrs    string `bson:"authNrs" json:"authNrs"`
}

func (db *DB) FindDrug(swissMedicId string) (Drug, error) {
	var drug Drug
	query := bson.M{"authNrs": swissMedicId}

	err := db.C("drugs").Find(query).One(&drug)
	checkIfObjectNotFound(&err)
	return drug, err
}
