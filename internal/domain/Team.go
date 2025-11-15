package domain

type Team struct {
	TeamName string
	Members  []User
}

func NewTeam(teamName string, members []User) *Team {
	return &Team{
		TeamName: teamName,
		Members:  members,
	}
}

func (t *Team) ActiveMembers() []User {
	activeMembers := make([]User, 0, len(t.Members))
	for _, member := range t.Members {
		if member.IsActive {
			activeMembers = append(activeMembers, member)
		}
	}
	return activeMembers
}

func (t *Team) NotActiveMembers() []User {
	notActiveMembers := make([]User, 0, len(t.Members))
	for _, member := range t.Members {
		if !member.IsActive {
			notActiveMembers = append(notActiveMembers, member)
		}
	}
	return notActiveMembers
}
