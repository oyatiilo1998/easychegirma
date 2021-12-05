package mongo

import (
	"context"

	vrss "chegirma_setting_service/ek_variables/chegirma_setting_service"
	ss "chegirma_setting_service/genproto/setting_service"
	"chegirma_setting_service/pkg/helper"
	"chegirma_setting_service/storage/repo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepo struct {
	collection *mongo.Collection
}

func NewUserRepo(db *mongo.Database) repo.UserStorageI {
	return &userRepo{
		collection: db.Collection("UserCollection")}
}

func (cr *userRepo) Create(user *ss.User) (string, error) {
	objectID, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		return "", err
	}
	createUser := vrss.CreateUser{
		ID:       objectID,
		Name:     user.Name,
		Link:     user.Link,
		Password: user.Password,
		Login:    user.Login,
		ImageUrl: user.ImageUrl,
	}
	resp, err := cr.collection.InsertOne(
		context.Background(),
		createUser,
	)
	if err != nil {
		return "", err
	}
	return resp.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (cr *userRepo) Get(id string) (*ss.User, error) {
	var userDecode vrss.User
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	if err := cr.collection.FindOne(
		context.Background(),
		bson.M{
			"_id": objectID,
		}).Decode(&userDecode); err != nil {
		return nil, err
	}

	return &ss.User{
		Id:       userDecode.ID,
		Name:     userDecode.Name,
		Link:     userDecode.Link,
		Password: userDecode.Password,
		Login:    userDecode.Login,
		ImageUrl: userDecode.ImageUrl,
	}, nil
}

func (cr *userRepo) GetAll(page, limit uint32, name string) ([]*ss.User, uint32, error) {
	var (
		response []*ss.User
		cities   []*vrss.User
		filter   = bson.D{}
	)
	if name != "" {
		filter = append(filter, primitive.E{Key: "name", Value: bson.D{
			primitive.E{Key: "$regex", Value: name},
			primitive.E{Key: "$options", Value: "im"},
		}})
	}

	opts := options.Find()
	skip := (page - 1) * limit
	opts.SetLimit(int64(limit))
	opts.SetSkip(int64(skip))
	opts.SetSort(bson.M{
		"name": 1,
	})
	count, err := cr.collection.CountDocuments(context.Background(), filter)

	if err != nil {
		return nil, 0, err
	}

	rows, err := cr.collection.Find(
		context.Background(),
		filter,
		opts,
	)
	if err != nil {
		return nil, 0, err
	}

	if err := rows.All(context.Background(), &cities); err != nil {
		return nil, 0, err
	}
	if err := helper.StructToProto(cities, &response); err != nil {
		return nil, 0, err
	}
	return response, uint32(count), nil
}
func (cr *userRepo) Delete(id string) error {
	result, err := cr.collection.DeleteOne(context.Background(), bson.M{"id": id})
	if err != nil || result.DeletedCount == 0 {
		return err
	}
	return nil
}
func (cr *userRepo) Update(user *ss.User) error {
	objectID, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		return err
	}
	update := bson.M{
		"$set": bson.M{
			"name":      user.Name,
			"password":  user.Password,
			"login":     user.Login,
			"image_url": user.ImageUrl,
			"link":      user.Link,
		}}
	filter := bson.M{"_id": bson.M{"$eq": objectID}}
	_, err = cr.collection.UpdateOne(
		context.Background(),
		filter,
		update)
	return err
}

var UserExists struct {
	Id     string `json:"id" bson:"id"`
	Exists bool   `json:"exists" bson:"exists"`
}

func (cr *userRepo) UserExists(login, password string) (*ss.UserExistsResponse, error) {
	var userDecode vrss.User
	filter := bson.D{primitive.E{Key: "login", Value: login}, primitive.E{Key: "password", Value: password}}
	if err := cr.collection.FindOne(
		context.Background(), filter).Decode(&userDecode); err != nil {
		return &ss.UserExistsResponse{
			Exist: false,
			Id:    "",
		}, err
	}

	return &ss.UserExistsResponse{Exist: true, Id: userDecode.ID}, nil
}
