package models

type Todo struct {
	ID          string `bson:"_id" json:"_id"`
	Title       string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`
}
