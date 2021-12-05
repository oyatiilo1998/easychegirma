package chegirma_setting_service

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       string `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Login    string `json:"login" bson:"login"`
	Password string `json:"password" bson:"password"`
	ImageUrl string `json:"image_url" bson:"image_url"`
	Link     string `json:"link" bson:"link"`
}

type CreateUser struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Login    string             `json:"login" bson:"login"`
	Password string             `json:"password" bson:"password"`
	ImageUrl string             `json:"image_url" bson:"image_url"`
	Link     string             `json:"link" bson:"link"`
}
type CreateUpdateUserSwag struct {
	ID       string `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Login    string `json:"login" bson:"login"`
	Password string `json:"password" bson:"password"`
	ImageUrl string `json:"image_url" bson:"image_url"`
	Link     string `json:"link" bson:"link"`
}
