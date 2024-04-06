package mango

import (
	"github.com/YngMin/mango/pkg/sliceutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Collection[T any] struct {
	database *mongo.Database
	*mongo.Collection
}

func GetCollection[T any](client *Client, databaseName, collectionName string) *Collection[T] {
	database := client.Database(databaseName)
	collection := database.Collection(collectionName)
	return &Collection[T]{
		database:   database,
		Collection: collection,
	}
}

func (c *Collection[T]) InsertOne(ctx Context, document T) (result *mongo.InsertOneResult, err error) {
	result, err = c.Collection.InsertOne(ctx, document)
	return
}

func (c *Collection[T]) InsertMany(ctx Context, documents []T) (result *mongo.InsertManyResult, err error) {
	result, err = c.Collection.InsertMany(ctx, sliceutil.Map(documents, func(idx int) any {
		return documents[idx]
	}))
	return
}

func (c *Collection[T]) FindOne(ctx Context, filter bson.M) (document T, err error) {
	result := c.Collection.FindOne(ctx, filter)
	if err = result.Err(); err != nil {
		return
	}

	err = result.Decode(&document)
	if err != nil {
		return
	}
	return
}

func (c *Collection[T]) Find(ctx Context, filter bson.M) (documents []T, err error) {
	var cur *mongo.Cursor
	cur, err = c.Collection.Find(ctx, filter)
	if err != nil {
		return
	}

	defer func() {
		_ = cur.Close(ctx)
	}()

	documents = make([]T, 0)
	for cur.Next(ctx) {
		var doc T
		if err = cur.Decode(&doc); err != nil {
			return
		}
		documents = append(documents, doc)
	}
	return
}

func (c *Collection[T]) UpdateOne(ctx Context, filter, update bson.M) (result *mongo.UpdateResult, err error) {
	result, err = c.Collection.UpdateOne(ctx, filter, update)
	return
}

func (c *Collection[T]) UpdateMany(ctx Context, filter, update bson.M) (result *mongo.UpdateResult, err error) {
	result, err = c.Collection.UpdateMany(ctx, filter, update)
	return
}

func (c *Collection[T]) UpdateByID(ctx Context, id primitive.ObjectID, update bson.M) (result *mongo.UpdateResult, err error) {
	result, err = c.Collection.UpdateByID(ctx, id, update)
	return
}
