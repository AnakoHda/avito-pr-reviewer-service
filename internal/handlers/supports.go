package handlers

import (
	"avito-pr-reviewer-service/internal/domain"
	"avito-pr-reviewer-service/internal/generated/api/dto"
	"encoding/json"
	"net/http"
)

const (
	ErrBadRequest dto.ErrorResponseErrorCode = "BAD_REQUEST"
	ErrInternal   dto.ErrorResponseErrorCode = "Internal error"
)

type TeamResponse struct {
	Team dto.Team `json:"team"`
}
type UserResponse struct {
	User dto.User `json:"user"`
}
type PullRequestsShort struct {
	PullRequestsShort []dto.PullRequestShort `json:"pull_requests"`
}
type PullRequestResponse struct {
	PullRequest dto.PullRequest `json:"pr"`
}

func FromTeamDTOToTeam(dto dto.Team) (domain.Team, error) {
	var tmpTeam domain.Team
	tmpTeam.TeamName = dto.TeamName
	for _, member := range dto.Members {
		tmpUser, err := domain.NewUser(domain.UserId(member.UserId), member.Username, dto.TeamName, member.IsActive)
		if err != nil {
			return domain.Team{}, err
		}
		tmpTeam.Members = append(tmpTeam.Members, *tmpUser)
	}
	return tmpTeam, nil
}
func FromTeamToDTO(team domain.Team) dto.Team {
	var tmpTeam dto.Team
	tmpTeam.TeamName = team.TeamName
	tmpTeam.Members = make([]dto.TeamMember, 0, len(team.Members))
	for _, member := range team.Members {
		tmpTeam.Members = append(tmpTeam.Members, dto.TeamMember{
			UserId:   string(member.UserId),
			Username: member.Username,
			IsActive: member.IsActive,
		})
	}
	return tmpTeam
}

func ResponseFormatError(w http.ResponseWriter, httpStatus int, errResponseCode dto.ErrorResponseErrorCode, massage string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	jsonResponse := dto.ErrorResponse{Error: struct {
		Code    dto.ErrorResponseErrorCode `json:"code"`
		Message string                     `json:"message"`
	}{
		Code:    errResponseCode,
		Message: massage,
	},
	}
	_ = json.NewEncoder(w).Encode(jsonResponse)
	return
}
func ResponseFormatOK(w http.ResponseWriter, httpStatus int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	_ = json.NewEncoder(w).Encode(data)
	return

}

func FromUserToUserDTO(user domain.User) dto.User {
	return dto.User{
		IsActive: user.IsActive,
		TeamName: user.TeamName,
		UserId:   string(user.UserId),
		Username: user.Username,
	}
}
