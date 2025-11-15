package domain

type userId string
type User struct {
	UserId   userId
	Username string
	TeamName string
	IsActive bool
}

func NewUser(userId userId, username, teamName string, active bool) (*User, error) {
	if err := ValidateUserFields(userId, username, teamName); err != nil {
		return nil, err
	}

	return &User{
		UserId:   userId,
		Username: username,
		TeamName: teamName,
		IsActive: active,
	}, nil
}

func (u *User) UpdateUser(username, teamName string, active bool) {
	u.Username = username
	u.TeamName = teamName
	u.IsActive = active
}

func ValidateUserFields(userId userId, username, teamName string) error {
	if userId == "" {
		return ErrEmptyUserID
	}
	if username == "" {
		return ErrEmptyUsername
	}
	if teamName == "" {
		return ErrEmptyTeamName
	}
	return nil
}
