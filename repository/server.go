package repository

import (
	"backend/json-server/model"
	"context"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

type ServerRepository struct {
	db *qmgo.Database
}

func NewServerRepository(db *qmgo.Database) *ServerRepository {
	return &ServerRepository{db}
}

func (s *ServerRepository) Find(ctx context.Context, filter bson.M) (model.Server, error) {
	var result model.Server
	err := s.db.Collection("json").Find(ctx, filter).One(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *ServerRepository) Insert(ctx context.Context, models model.Server) error {
	_, err := s.db.Collection("json").InsertOne(ctx, models)
	if err != nil {
		return err
	}

	return nil
}
