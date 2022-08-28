package usecase

import "go.mongodb.org/mongo-driver/bson/primitive"

func StringToObjectID(value string) (object primitive.ObjectID, err error) {
	return primitive.ObjectIDFromHex(value)
}
