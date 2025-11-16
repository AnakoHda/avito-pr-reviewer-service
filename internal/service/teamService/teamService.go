package teamService

import (
	"avito-pr-reviewer-service/internal/domain"
	"context"
)

type UsersRepository interface {
	CreateAndUpdateUsers(ctx context.Context, users []domain.User) error
	ListUsersByTeamName(ctx context.Context, teamName string) ([]domain.User, error)
}
type TeamRepository interface {
	CreateTeam(ctx context.Context, teamName string) error
}
type Service struct {
	userRepo UsersRepository
	teamRepo TeamRepository
}

func New(usersRepository UsersRepository, teamRepository TeamRepository) *Service {
	return &Service{
		userRepo: usersRepository,
		teamRepo: teamRepository,
	}
}

func (s *Service) AddTeamWithUsers(ctx context.Context, team domain.Team) error {
	// создание команды
	if err := s.teamRepo.CreateTeam(ctx, team.TeamName); err != nil {
		return err
	}

	if err := s.userRepo.CreateAndUpdateUsers(ctx, team.Members); err != nil {
		return err
	}
	return nil
}

func (s *Service) getTeam(ctx context.Context, teamName string) (*domain.Team, error) {
	// существует ли команда и получить скисок
	existedUsersInTeam, err := s.userRepo.ListUsersByTeamName(ctx, teamName)
	if err != nil {
		return nil, err
	}
	return domain.NewTeam(teamName, existedUsersInTeam)
}
