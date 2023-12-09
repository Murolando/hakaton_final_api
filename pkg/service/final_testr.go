package service

import (
	"github.com/Murolando/hakaton_final_api/ent"
	"github.com/Murolando/hakaton_final_api/pkg/repository"
)

type FinalTestService struct {
	repo *repository.Repository
}

func NewFinalTestService(repo *repository.Repository) *FinalTestService {
	return &FinalTestService{
		repo: repo,
	}
}
func (s *FinalTestService) StartFinalTest(userId int) (*ent.FinalTest, error) {

	return s.repo.FinalTestPst.StartFinalTest(userId)
}
