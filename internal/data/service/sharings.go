package service

import (
	"github.com/romarq/visualtez-storage/internal/data/dto"
	Repository "github.com/romarq/visualtez-storage/internal/data/repository"
	"github.com/romarq/visualtez-storage/pkg/utils"
)

type SharingsService struct {
	SharingsRepository Repository.SharingsRepository
}

func InitSharingsService(r Repository.SharingsRepository) SharingsService {
	return SharingsService{SharingsRepository: r}
}

// GetBigMapByID - Get sharing by hash
func (s *SharingsService) GetSharing(hash string) (dto.Sharing, error) {
	encrypted, err := s.SharingsRepository.GetSharing(hash)
	return encrypted, err
}

// InsertSharing - Inserts a sharing record
func (s *SharingsService) InsertSharing(req dto.CreateSharing_Params) (dto.Sharing, error) {
	// Generate hash
	hash := utils.SHA256(req.Content)
	sharing := dto.Sharing{
		Hash:    hash,
		Content: req.Content,
	}

	return sharing, s.SharingsRepository.InsertSharing(sharing)
}
