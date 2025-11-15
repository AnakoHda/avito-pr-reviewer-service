package domain

import (
	"time"
)

type PullRequestStatus string

const (
	PullRequestStatusOPEN   PullRequestStatus = "OPEN"
	PullRequestStatusMERGED PullRequestStatus = "MERGED"
)

type pullRequestId string
type PullRequest struct {
	PullRequestId     pullRequestId
	PullRequestName   string
	AuthorId          userId
	Status            PullRequestStatus
	AssignedReviewers []userId
	CreatedAt         time.Time
	MergedAt          time.Time
}

func NewPullRequest(pullRequestId pullRequestId, pullRequestName string, authorId userId, assignedReviewers []userId) (*PullRequest, error) {
	if err := ValidatePullRequestFields(pullRequestId, pullRequestName, authorId, assignedReviewers); err != nil {
		return nil, err
	}
	return &PullRequest{
		PullRequestId:     pullRequestId,
		PullRequestName:   pullRequestName,
		AuthorId:          authorId,
		Status:            PullRequestStatusOPEN,
		AssignedReviewers: assignedReviewers,
		CreatedAt:         time.Now(),
		MergedAt:          time.Time{},
	}, nil
}
func (pr *PullRequest) Merge() error {
	if pr.Status == PullRequestStatusMERGED {
		return ErrPullRequestMerged
	}

	pr.Status = PullRequestStatusMERGED
	pr.MergedAt = time.Now()
	return nil
}

func (pr *PullRequest) ReplaceReviewer(oldReviewerId, newReviewerId userId) error {
	if pr.Status == PullRequestStatusMERGED {
		return ErrPullRequestMerged
	}
	if newReviewerId == pr.AuthorId {
		return ErrAuthorCannotBeReviewer
	}

	for i, findUserId := range pr.AssignedReviewers {
		if findUserId == oldReviewerId {
			pr.AssignedReviewers[i] = newReviewerId
			return nil
		}
	}

	return ErrNotFoundReviewerInPullRequest
}

func ValidatePullRequestFields(pullRequestId pullRequestId, pullRequestName string, authorId userId, assignedReviewers []userId) error {
	if pullRequestId == "" {
		return ErrEmptyPullRequestID
	}
	if pullRequestName == "" {
		return ErrEmptyPullRequestName
	}
	if authorId == "" {
		return ErrEmptyAuthorID
	}
	if len(assignedReviewers) > 2 {
		return ErrTooManyReviewersAssigned
	}
	for _, reviewer := range assignedReviewers {
		if reviewer == authorId {
			return ErrAuthorCannotBeReviewer
		}
	}
	return nil
}
