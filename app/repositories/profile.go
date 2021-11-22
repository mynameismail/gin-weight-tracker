package repositories

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"gin-weight-tracker/services"
)

type Profile struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Birth    string `json:"birth"`
}

func (h Profile) collection() *mongo.Collection {
	collname := "profiles"
	coll := services.MongoClient.Database(getDBName()).Collection(collname)
	return coll
}

func (h Profile) All() []Profile {
	var profiles []Profile

	cursor, err := h.collection().Find(ctx, bson.D{})
	if err != nil {
		return profiles
	}

	for cursor.Next(ctx) {
		var profile bson.M
		err = cursor.Decode(&profile)
		if err != nil {
			return profiles
		}

		profiles = append(profiles, Profile{
			ID: profile["_id"].(primitive.ObjectID).Hex(),
			Username: profile["username"].(string),
			Birth: profile["birth"].(string),
		})
	}

	return profiles
}
