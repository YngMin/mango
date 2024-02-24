package mango

import (
	"go.mongodb.org/mongo-driver/bson"
	"ymgo/pkg/errors"
)

type FindQuery struct {
	ctx    Context
	filter bson.M
	sort   []string
}

func Find(ctx Context) *FindQuery {
	return &FindQuery{
		ctx:    ctx,
		filter: make(bson.M),
		sort:   make([]string, 0),
	}
}

func (f *FindQuery) Equals(key string, value any) *FindQuery {
	f.filter[key] = value
	return f
}

func (f *FindQuery) In(key string, values any) *FindQuery {
	f.filter[key] = bson.M{
		"$in": values,
	}
	return f
}

func (f *FindQuery) Like(key, value string, ignoreCase bool) *FindQuery {
	regexValue := bson.M{
		"$regex": value,
	}
	if ignoreCase {
		regexValue["$options"] = "i"
	}
	f.filter[key] = regexValue
	return f
}

func (f *FindQuery) GT(key string, value any) *FindQuery {
	f.filter[key] = bson.M{
		"$gt": value,
	}
	return f
}

func (f *FindQuery) GTE(key string, value any) *FindQuery {
	f.filter[key] = bson.M{
		"$gte": value,
	}
	return f
}

func (f *FindQuery) LT(key string, value any) *FindQuery {
	f.filter[key] = bson.M{
		"$lt": value,
	}
	return f
}

func (f *FindQuery) LTE(key string, value any) *FindQuery {
	f.filter[key] = bson.M{
		"$lte": value,
	}
	return f
}

func (f *FindQuery) Sort(key string, desc bool) *FindQuery {
	if desc {
		key = "-" + key
	}
	f.sort = append(f.sort, key)
	return f
}

func (f *FindQuery) One(dest ICollection) (err error) {
	err = f.validate()
	if err != nil {
		return
	}

	collection := f.ctx.db.Collection(dest)
	err = collection.FindOne(f.ctx, f.filter, dest)
	if err != nil {
		return
	}
	return
}

func (f *FindQuery) All(dest ICollection) (err error) {
	err = f.validate()
	if err != nil {
		return
	}

	collection := f.ctx.db.Collection(dest)
	err = collection.Find(f.ctx, f.filter, dest)
	if err != nil {
		return
	}
	return
}

func (f *FindQuery) validate() (err error) {
	if f.ctx.db == nil {
		err = errors.ErrNeedDatabase
		return
	}
	return
}
