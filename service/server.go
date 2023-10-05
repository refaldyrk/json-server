package service

import (
	"backend/json-server/model"
	"backend/json-server/repository"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ServerService struct {
	serverRepository *repository.ServerRepository
}

func NewServerService(serverRepository *repository.ServerRepository) *ServerService {
	return &ServerService{serverRepository}
}

func (s *ServerService) InsertNewJSON(ctx context.Context, jsons any) (string, error) {
	id := uuid.NewString()

	byteJson, err := json.Marshal(jsons)
	if err != nil {
		return "", err
	}

	models := model.Server{
		ID:        primitive.NewObjectID(),
		ServerID:  id,
		Data:      byteJson,
		Raw:       map[string]any{},
		CreatedAt: 0,
	}

	err = s.serverRepository.Insert(ctx, models)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *ServerService) GetJSON(ctx context.Context, server_id string) (model.Server, error) {
	result, err := s.serverRepository.Find(ctx, bson.M{"server_id": server_id})
	if err != nil {
		return result, err
	}

	//Unmarshal JSON
	var data map[string]any
	err = json.Unmarshal(result.Data, &data)
	if err != nil {
		return model.Server{}, err
	}

	result.Raw = data

	return result, nil
}
