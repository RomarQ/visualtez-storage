package repository

import (
	"context"

	"github.com/romarq/visualtez-storage/internal/data/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SharingsRepository struct {
	Collection *mongo.Collection
}

func InitSharingsRepository(DB *mongo.Database) SharingsRepository {
	return SharingsRepository{Collection: DB.Collection("sharings")}
}

// GetSharing - Queries the database to get a given sharing by hash
func (r *SharingsRepository) GetSharing(hash string) (dto.Sharing, error) {
	var result dto.Sharing
	r.Collection.InsertOne(context.TODO(), bson.M{"hash": hash})
	err := r.Collection.FindOne(context.TODO(), bson.M{"hash": hash}).Decode(&result)
	return result, err
}

// InsertSharing - Inserts a sharing record
func (r *SharingsRepository) InsertSharing(sharing dto.Sharing) error {
	_, err := r.Collection.InsertOne(context.TODO(), sharing)
	return err
}
