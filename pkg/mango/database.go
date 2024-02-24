package mango

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	database *mongo.Database
}

// Collection TODO Support CollectionOption
func (d *Database) Collection(o ICollection) (collection *Collection) {
	c := d.database.Collection(o.CollectionName())
	collection = &Collection{
		collection: c,
	}
	return
}

func (d *Database) CollectionByName(name string) (collection *Collection) {
	c := d.database.Collection(name)
	collection = &Collection{
		collection: c,
	}
	return
}
