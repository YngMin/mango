package ymgo

import (
	"github.com/iancoleman/strcase"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"strings"
)

type Collection struct {
	collection *mongo.Collection
}

type InsertOneResult *mongo.InsertOneResult

type InsertManyResult *mongo.InsertManyResult

type UpdateResult *mongo.UpdateResult

func (c *Collection) InsertOne(ctx Context, document any) (result InsertOneResult, err error) {
	result, err = c.collection.InsertOne(ctx, document)
	return
}

func (c *Collection) InsertMany(ctx Context, documents []any) (result InsertManyResult, err error) {
	result, err = c.collection.InsertMany(ctx, documents)
	return
}

func (c *Collection) FindOne(ctx Context, filter bson.M, dest any) (err error) {
	result := c.collection.FindOne(ctx, filter)
	if resultErr := result.Err(); resultErr != nil {
		err = resultErr
		return
	}

	err = result.Decode(dest)
	if err != nil {
		return
	}
	return
}

func (c *Collection) Find(ctx Context, filter bson.M, dest any) (err error) {
	var cur *mongo.Cursor
	cur, err = c.collection.Find(ctx, filter)
	if err != nil {
		return
	}

	defer func() {
		_ = cur.Close(ctx)
	}()

	err = cur.All(ctx, dest)
	if err != nil {
		return
	}

	return
}

func (c *Collection) UpdateOne(ctx Context, filter bson.M, update bson.M) (result UpdateResult, err error) {
	result, err = c.collection.UpdateOne(ctx, filter, update)
	return
}

func (c *Collection) UpdateMany(ctx Context, filter bson.M, update bson.M) (result UpdateResult, err error) {
	result, err = c.collection.UpdateMany(ctx, filter, update)
	return
}

func (c *Collection) UpdateByID(ctx Context, id primitive.ObjectID, update bson.M) (result UpdateResult, err error) {
	result, err = c.collection.UpdateByID(ctx, id, update)
	return
}

func getCollectionName(o any) string {
	typeName := reflect.TypeOf(o).String()
	split := strings.Split(typeName, ".")
	return strcase.ToLowerCamel(split[len(split)-1])
}
