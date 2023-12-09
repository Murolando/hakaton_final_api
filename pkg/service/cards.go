package service

import (
	"github.com/Murolando/hakaton_final_api/ent"
	"github.com/Murolando/hakaton_final_api/pkg/repository"
)

type CardService struct {
	repo *repository.Repository
}

func NewCardService(repo *repository.Repository) *CardService {
	return &CardService{
		repo: repo,
	}
}

func (s *CardService) AddCard(card *ent.Card, userId int) (bool, error) {
	return s.repo.AddCard(card,userId)
}
func (s *CardService) GetCard(userId int) ([]*ent.Card, error){
	return s.repo.GetCard(userId)
}
func (s *CardService) DeleteCard(cardId int) (bool, error){
	
	return s.repo.DeleteCard(cardId)
}
