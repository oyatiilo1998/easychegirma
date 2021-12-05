package helper

import (
	"encoding/json"

	"chegirma_setting_service/pkg/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleError(log logger.Logger, err error, message string, req interface{}) error {

	if err == mongo.ErrNoDocuments {
		log.Error(message+", Not Found", logger.Error(err), logger.Any("req", req))
		return status.Error(codes.NotFound, "Not Found")
	} else if err != nil {
		log.Error(message, logger.Error(err), logger.Any("req", req))
		return status.Error(codes.Internal, message+err.Error())
	}
	return nil
}

func StructToProto(structure, protoMessage interface{}) error {
	byte, err := json.Marshal(structure)
	if err != nil {
		return err
	}
	err = json.Unmarshal(byte, protoMessage)
	return err
}

func MarshalUnmarshal(reqStructure, resStructure interface{}) error {
	byte, err := json.Marshal(reqStructure)
	if err != nil {
		return err
	}
	err = json.Unmarshal(byte, resStructure)
	return err
}
