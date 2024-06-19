package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserName string             `bson:"user_name" json:"user_name"`
	UserID   string             `bson:"user_id" json:"user_id"`
	Email    string             `bson:"email" json:"email"`
	DOB      string             `bson:"dob" json:"dob"` // Date of Birth
	Gender   string             `bson:"gender" json:"gender"`
}
