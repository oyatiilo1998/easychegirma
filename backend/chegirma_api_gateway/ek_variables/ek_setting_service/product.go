package ek_setting_service

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID             string    `json:"id" bson:"_id"`
	Name           string    `json:"name" bson:"name"`
	Description    string    `json:"description" bson:"description"`
	URL            string    `json:"url" bson:"url"`
	ImageURL       string    `json:"image_url" bson:"image_url"`
	Category       Category  `json:"category" bson:"category"`
	DiscountAmount int       `json:"discount_amount" bson:"discount_amount"`
	PriceBefore    int       `json:"price_before" bson:"price_before"`
	PriceAfter     int       `json:"price_after" bson:"price_after"`
	Owner          User      `json:"owner" bson:"owner"`
	FromTime       time.Time `json:"from_time" bson:"from_time"`
	ToTime         time.Time `json:"to_time" bson:"to_time"`
}
type EntityProduct struct {
	Name   string `json:"name" bson:"name"`
	RuName string `json:"ru_name" bson:"ru_name"`
	Code   uint32 `json:"code" bson:"code"`
}

type CreateProduct struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	Name           string             `json:"name" bson:"name"`
	Description    string             `json:"description" bson:"description"`
	URL            string             `json:"url" bson:"url"`
	ImageURL       string             `json:"image_url" bson:"image_url"`
	CategoryID     primitive.ObjectID `json:"category_id" bson:"category_id"`
	DiscountAmount int                `json:"discount_amount" bson:"discount_amount"`
	PriceBefore    int                `json:"price_before" bson:"price_before"`
	PriceAfter     int                `json:"price_after" bson:"price_after"`
	OwnerId        primitive.ObjectID `json:"owner_id" bson:"owner_id"`
	FromTime       time.Time          `json:"from_time" bson:"from_time"`
	ToTime         time.Time          `json:"to_time" bson:"to_time"`
}
type CreateUpdateProductSwag struct {
	ID             string    `json:"id" bson:"_id"`
	Name           string    `json:"name" bson:"name"`
	Description    string    `json:"description" bson:"description"`
	URL            string    `json:"url" bson:"url"`
	ImageURL       string    `json:"image_url" bson:"image_url"`
	CategoryID     string    `json:"category_id" bson:"category_id"`
	DiscountAmount int       `json:"discount_amount" bson:"discount_amount"`
	PriceBefore    int       `json:"price_before" bson:"price_before"`
	PriceAfter     int       `json:"price_after" bson:"price_after"`
	OwnerId        string    `json:"owner_id" bson:"owner_id"`
	FromTime       time.Time `json:"from_time" bson:"from_time"`
	ToTime         time.Time `json:"to_time" bson:"to_time"`
}
