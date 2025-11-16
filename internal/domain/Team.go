package domain

type Team struct {
	TeamName string
	Members  []User
}

func NewTeam(teamName string, members []User) (*Team, error) {
	if err := ValidateTeamName(teamName); err != nil {
		return nil, err
	}
	return &Team{
		TeamName: teamName,
		Members:  members,
	}, nil
}

func ValidateTeamName(teamName string) error {
	if teamName == "" {
		return ErrEmptyTeamName
	}
	return nil
}
