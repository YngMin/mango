package ymgo

import (
	"go.mongodb.org/mongo-driver/bson"
	"ymgo/pkg/errors"
)

type Find struct {
	ctx    Context
	filter bson.M
	sort   []string
}

func FindQuery(ctx Context) *Find {
	return &Find{
		ctx:    ctx,
		filter: make(bson.M),
		sort:   make([]string, 0),
	}
}

func (f *Find) Equals(key string, value any) *Find {
	f.filter[key] = value
	return f
}

func (f *Find) In(key string, values any) *Find {
	f.filter[key] = bson.M{
		"$in": values,
	}
	return f
}

func (f *Find) Like(key, value string, ignoreCase bool) *Find {
	regexValue := bson.M{
		"$regex": value,
	}
	if ignoreCase {
		regexValue["$options"] = "i"
	}
	f.filter[key] = regexValue
	return f
}

func (f *Find) GT(key string, value any) *Find {
	f.filter[key] = bson.M{
		"$gt": value,
	}
	return f
}

func (f *Find) GTE(key string, value any) *Find {
	f.filter[key] = bson.M{
		"$gte": value,
	}
	return f
}

func (f *Find) LT(key string, value any) *Find {
	f.filter[key] = bson.M{
		"$lt": value,
	}
	return f
}

func (f *Find) LTE(key string, value any) *Find {
	f.filter[key] = bson.M{
		"$lte": value,
	}
	return f
}

func (f *Find) Sort(key string, desc bool) *Find {
	if desc {
		key = "-" + key
	}
	f.sort = append(f.sort, key)
	return f
}

func (f *Find) FindOne(dest any) (err error) {
	err = f.validate()
	if err != nil {
		return
	}

	collection := f.ctx.db.Collection(getCollectionName(dest))
	err = collection.FindOne(f.ctx, f.filter, dest)
	if err != nil {
		return
	}
	return
}

func (f *Find) Find(dest any) (err error) {
	err = f.validate()
	if err != nil {
		return
	}

	collection := f.ctx.db.Collection(getCollectionName(dest))
	err = collection.Find(f.ctx, f.filter, dest)
	if err != nil {
		return
	}
	return
}

func (f *Find) validate() (err error) {
	if f.ctx.db == nil {
		err = errors.ErrNeedDatabase
		return
	}
	return
}
