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

type categoryRepo struct {
	collection *mongo.Collection
}

func NewCategoryRepo(db *mongo.Database) repo.CategoryStorageI {
	return &categoryRepo{
		collection: db.Collection("CategoryCollection")}
}

func (cr *categoryRepo) Create(category *ss.Category) (string, error) {
	objectID, err := primitive.ObjectIDFromHex(category.Id)
	if err != nil {
		return "", err
	}
	createCategory := vrss.CreateCategory{
		ID:     objectID,
		Name:   category.Name,
		RuName: category.RuName,
		Code:   category.Code,
		Soato:  category.Soato,
	}
	resp, err := cr.collection.InsertOne(
		context.Background(),
		createCategory,
	)
	if err != nil {
		return "", err
	}
	return resp.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (cr *categoryRepo) Get(id string) (*ss.Category, error) {
	var categoryDecode vrss.Category
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	if err := cr.collection.FindOne(
		context.Background(),
		bson.M{
			"_id": objectID,
		}).Decode(&categoryDecode); err != nil {
		return nil, err
	}

	return &ss.Category{
		Id:     categoryDecode.ID,
		Name:   categoryDecode.Name,
		RuName: categoryDecode.RuName,
		Code:   categoryDecode.Code,
		Soato:  categoryDecode.Soato,
	}, nil
}

func (cr *categoryRepo) GetAll(page, limit, code uint32, name string) ([]*ss.Category, uint32, error) {
	var (
		response []*ss.Category
		cities   []*vrss.Category
		filter   = bson.D{}
	)
	if name != "" {
		filter = append(filter, primitive.E{Key: "name", Value: bson.D{
			primitive.E{Key: "$regex", Value: name},
			primitive.E{Key: "$options", Value: "im"},
		}})
	}
	if code != 0 {
		filter = append(filter, primitive.E{Key: "code", Value: code})
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
func (cr *categoryRepo) Delete(id string) error {
	result, err := cr.collection.DeleteOne(context.Background(), bson.M{"id": id})
	if err != nil || result.DeletedCount == 0 {
		return err
	}
	return nil
}
func (cr *categoryRepo) Update(category *ss.Category) error {
	objectID, err := primitive.ObjectIDFromHex(category.Id)
	if err != nil {
		return err
	}
	update := bson.M{
		"$set": bson.M{
			"name":    category.Name,
			"code":    category.Code,
			"ru_name": category.RuName,
			"soato":   category.Soato,
		}}
	filter := bson.M{"_id": bson.M{"$eq": objectID}}
	_, err = cr.collection.UpdateOne(
		context.Background(),
		filter,
		update)
	return err
}

func (cr *categoryRepo) CategoryExists(id string) (bool, error) {
	categoryObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}
	filter := bson.D{primitive.E{Key: "_id", Value: categoryObjectID}}
	count, err := cr.collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return false, err
	}
	return count == 1, nil
}

func (cr *categoryRepo) GetByCode(soato int) (*ss.Category, error) {
	var categoryDecode vrss.Category
	if err := cr.collection.FindOne(
		context.Background(),
		bson.M{
			"soato": soato,
		}).Decode(&categoryDecode); err != nil {
		return nil, err
	}

	return &ss.Category{
		Id:     categoryDecode.ID,
		Name:   categoryDecode.Name,
		RuName: categoryDecode.RuName,
		Code:   categoryDecode.Code,
		Soato:  categoryDecode.Soato,
	}, nil
}
