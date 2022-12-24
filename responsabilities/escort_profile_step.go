package responsabilities

import (
	"context"
	"fmt"

	"escort-book-delete-customers/models"
	"escort-book-delete-customers/repositories"
)

type EscortProfileStep struct {
	EscortProfileRepository repositories.IEscortProfileRepository
	Next                    IAccountRemovalStep
}

func (s EscortProfileStep) Execute(ctx context.Context, removals []models.CustomerRemoval) {
	stringOfUserIds := convertUserIdToStringByWrapperChar(removals)
	query := fmt.Sprintf("DELETE FROM profile WHERE escort_id IN (%s)", stringOfUserIds)

	if err := s.EscortProfileRepository.Delete(ctx, query); err != nil {
		logger.Warn(fmt.Sprintf("Error deleting escort profiles: %s", err.Error()))
	}

	s.Next.Execute(ctx, removals)
}
