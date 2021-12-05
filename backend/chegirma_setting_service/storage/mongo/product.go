package mongo

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"chegirma_setting_service/ek_variables"
	vrss "chegirma_setting_service/ek_variables/chegirma_setting_service"
	ss "chegirma_setting_service/genproto/setting_service"
	"chegirma_setting_service/pkg/helper"
	"chegirma_setting_service/storage/repo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type productRepo struct {
	collection *mongo.Collection
}

func NewProductRepo(db *mongo.Database) repo.ProductStorageI {
	return &productRepo{
		collection: db.Collection("ProductCollection")}
}

func (cr *productRepo) Create(product *ss.CreateUpdateProduct) (string, error) {
	objectID, err := primitive.ObjectIDFromHex(product.Id)
	if err != nil {
		return "", err
	}
	categoryID, err := primitive.ObjectIDFromHex(product.CategoryId)
	if err != nil {
		return "", err
	}
	ownerID, err := primitive.ObjectIDFromHex(product.OwnerId)
	if err != nil {
		return "", err
	}
	toTime, err := time.Parse(ek_variables.TimeLayout, product.ToTime)
	if err != nil {
		return "", err
	}
	fromTime, err := time.Parse(ek_variables.TimeLayout, product.FromTime)
	if err != nil {
		return "", err
	}
	createProduct := vrss.CreateProduct{
		ID:             objectID,
		Name:           product.Name,
		URL:            product.Url,
		ImageURL:       product.ImageUrl,
		DiscountAmount: int(product.DiscountAmount),
		PriceBefore:    int(product.PriceBefore),
		PriceAfter:     int(product.PriceAfter),
		CategoryID:     categoryID,
		FromTime:       fromTime,
		ToTime:         toTime,
		OwnerId:        ownerID,
		Description:    product.Description,
	}
	resp, err := cr.collection.InsertOne(
		context.Background(),
		createProduct,
	)
	if err != nil {
		return "", err
	}
	return resp.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (cr *productRepo) Get(id string) (*ss.Product, error) {
	var (
		product  []*vrss.Product
		response *ss.Product
		pipeline = mongo.Pipeline{}
	)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	pipeline = append(pipeline,
		bson.D{
			primitive.E{Key: "$match", Value: bson.D{
				primitive.E{Key: "_id", Value: objectID}}}},
		bson.D{
			primitive.E{Key: "$lookup", Value: bson.D{
				primitive.E{Key: "from", Value: "CategoryCollection"},
				primitive.E{Key: "localField", Value: "category_id"},
				primitive.E{Key: "foreignField", Value: "_id"},
				primitive.E{Key: "as", Value: "category"}}}},
		bson.D{
			primitive.E{Key: "$lookup", Value: bson.D{
				primitive.E{Key: "from", Value: "UserCollection"},
				primitive.E{Key: "localField", Value: "owner_id"},
				primitive.E{Key: "foreignField", Value: "_id"},
				primitive.E{Key: "as", Value: "owner"}}}},
		bson.D{primitive.E{Key: "$unwind", Value: bson.D{
			primitive.E{Key: "path", Value: "$category"}, {Key: "preserveNullAndEmptyArrays", Value: true}}}},
		bson.D{primitive.E{Key: "$unwind", Value: bson.D{
			primitive.E{Key: "path", Value: "$owner"}, {Key: "preserveNullAndEmptyArrays", Value: true}}}})

	row, err := cr.collection.Aggregate(
		context.Background(),
		pipeline)
	defer func() {
		_ = row.Close(context.Background())
	}()
	if err != nil {
		return nil, err
	}

	if err := row.All(context.Background(), &product); err != nil {
		return nil, err
	}

	if len(product) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	byte, err := json.Marshal(product[0])
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(byte, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func (cr *productRepo) GetAll(page, limit uint32, category_id, owner_id, name string) ([]*ss.Product, uint32, error) {
	var (
		response []*ss.Product
		products []*vrss.Product
		filter   = bson.D{}
		pipeline = mongo.Pipeline{}
		skip     = (page - 1) * limit
	)
	if name != "" {
		filter = append(filter, bson.E{Key: "name", Value: bson.D{
			primitive.E{Key: "$regex", Value: name},
			primitive.E{Key: "$options", Value: "im"},
		}})
		pipeline = append(pipeline, bson.D{
			primitive.E{Key: "$match", Value: bson.D{
				primitive.E{Key: "name", Value: bson.D{
					primitive.E{Key: "$regex", Value: name},
					primitive.E{Key: "$options", Value: "im"},
				}}}}})
	}
	if category_id != "" {
		categoryID, err := primitive.ObjectIDFromHex(category_id)
		if err != nil {
			return nil, 0, err
		}
		filter = append(filter, bson.E{Key: "category_id", Value: categoryID})
		pipeline = append(pipeline, bson.D{
			primitive.E{Key: "$match", Value: bson.D{
				primitive.E{Key: "category_id", Value: categoryID}}}})
	}
	if owner_id != "" {
		ownerID, err := primitive.ObjectIDFromHex(owner_id)
		if err != nil {
			return nil, 0, err
		}
		filter = append(filter, bson.E{Key: "owner_id", Value: ownerID})
		pipeline = append(pipeline, bson.D{
			primitive.E{Key: "$match", Value: bson.D{
				primitive.E{Key: "owner_id", Value: ownerID}}}})
	}
	pipeline = append(pipeline,
		bson.D{primitive.E{Key: "$skip", Value: skip}},
		bson.D{primitive.E{Key: "$limit", Value: limit}},
		bson.D{primitive.E{Key: "$sort", Value: bson.M{"name": -1}}},
		bson.D{
			primitive.E{Key: "$lookup", Value: bson.D{
				primitive.E{Key: "from", Value: "CategoryCollection"},
				primitive.E{Key: "localField", Value: "category_id"},
				primitive.E{Key: "foreignField", Value: "_id"},
				primitive.E{Key: "as", Value: "category"}}}},
		bson.D{
			primitive.E{Key: "$lookup", Value: bson.D{
				primitive.E{Key: "from", Value: "UserCollection"},
				primitive.E{Key: "localField", Value: "owner_id"},
				primitive.E{Key: "foreignField", Value: "_id"},
				primitive.E{Key: "as", Value: "owner"}}}},
		bson.D{primitive.E{Key: "$unwind", Value: bson.D{
			primitive.E{Key: "path", Value: "$category"}, {Key: "preserveNullAndEmptyArrays", Value: false}}}},
		bson.D{primitive.E{Key: "$unwind", Value: bson.D{
			primitive.E{Key: "path", Value: "$owner"}, {Key: "preserveNullAndEmptyArrays", Value: true}}}})
	count, err := cr.collection.CountDocuments(context.Background(), filter)

	if err != nil {
		return nil, 0, err
	}

	rows, err := cr.collection.Aggregate(
		context.Background(), pipeline)

	if err != nil {
		return nil, 0, err
	}

	if err := rows.All(context.Background(), &products); err != nil {
		return nil, 0, err
	}
	if err := helper.StructToProto(products, &response); err != nil {
		return nil, 0, err
	}
	fmt.Println(response)
	return response, uint32(count), nil
}
func (cr *productRepo) Delete(id string) error {
	result, err := cr.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil || result.DeletedCount == 0 {
		return err
	}
	return nil
}
func (cr *productRepo) Update(product *ss.CreateUpdateProduct) error {
	objectID, err := primitive.ObjectIDFromHex(product.Id)
	if err != nil {
		return err
	}
	categoryID, err := primitive.ObjectIDFromHex(product.CategoryId)
	if err != nil {
		return err
	}
	fromTime, err := time.Parse(ek_variables.TimeLayout, product.FromTime)
	if err != nil {
		return err
	}
	toTime, err := time.Parse(ek_variables.TimeLayout, product.ToTime)
	if err != nil {
		return err
	}
	update := bson.M{
		"$set": bson.M{
			"url":             product.Url,
			"name":            product.Name,
			"description":     product.Description,
			"image_url":       product.ImageUrl,
			"discount_amount": product.DiscountAmount,
			"price_before":    product.PriceBefore,
			"price_after":     product.PriceAfter,
			"category_id":     categoryID,
			"from_time":       fromTime,
			"to_time":         toTime,
		}}
	filter := bson.M{"_id": bson.M{"$eq": objectID}}
	_, err = cr.collection.UpdateOne(
		context.Background(),
		filter,
		update)
	return err
}

func (cr *productRepo) ProductExists(id string) (bool, error) {
	productObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}
	filter := bson.D{primitive.E{Key: "_id", Value: productObjectID}}
	count, err := cr.collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return false, err
	}
	return count == 1, nil
}
