package ymgo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	database *mongo.Database
}

// Collection TODO Support CollectionOption
func (d *Database) Collection(name string) (collection *Collection) {
	c := d.database.Collection(name)
	collection = &Collection{
		collection: c,
	}
	return
}
