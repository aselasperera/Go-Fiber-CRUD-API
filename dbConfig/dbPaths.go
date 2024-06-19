package dbConfig

import (
	mongo "go.mongodb.org/mongo-driver/mongo"
)

var DATABASE *mongo.Database

const DATABASE_URL = "mongodb+srv://aselasperera:Asela9907@users.2htrxsy.mongodb.net/?retryWrites=true&w=majority&appName=users"

const DATABASE_NAME = "TestCGaaSApp-APP211"
