package service

import "go-ms-profile/repository"

type Service struct {
	mg *repository.MongoRepository
}

func NewProfileService(mg *repository.MongoRepository) *Service {
	return &Service{
		mg: mg,
	}
}


