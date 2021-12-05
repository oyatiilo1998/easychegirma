package ek_setting_service

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	ID         string `json:"id" bson:"_id"`
	RuName     string `json:"ru_name" bson:"ru_name"`
	Name       string `json:"name" bson:"name"`
	ExternalID uint32 `json:"external_id" bson:"external_id"`
	Soato      uint32 `json:"soato" bson:"soato"`
	Code       uint32 `json:"code" bson:"code"`
}
type EntityCategory struct {
	Name   string `json:"name" bson:"name"`
	RuName string `json:"ru_name" bson:"ru_name"`
	Code   uint32 `json:"code" bson:"code"`
}

type CreateCategory struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	RuName     string             `json:"ru_name" bson:"ru_name"`
	Name       string             `json:"name" bson:"name"`
	ExternalID uint32             `json:"external_id" bson:"external_id"`
	Soato      uint32             `json:"soato" bson:"soato"`
	Code       uint32             `json:"code" bson:"code"`
}
type CreateUpdateCategorySwag struct {
	Name       string `json:"name" binding:"required"`
	RuName     string `json:"ru_name" binding:"required"`
	ExternalID uint32 `json:"external_id" bson:"external_id"`
	Soato      uint32 `json:"soato" bson:"soato"`
	Code       uint32 `json:"code" binding:"required"`
}
