package responsabilities

import (
	"context"
	"fmt"
	"strings"

	"escort-book-delete-customers/models"
	"escort-book-delete-customers/repositories"

	log "github.com/inconshreveable/log15"
)

type CustomerProfileStep struct {
	CustomerProfileRepository repositories.ICustomerProfileRepository
	Next                      IAccountRemovalStep
}

var logger = log.New("reponsibilities")

func (s CustomerProfileStep) Execute(ctx context.Context, removals []models.CustomerRemoval) {
	stringOfUserIds := convertUserIdToStringByWrapperChar(removals)
	query := fmt.Sprintf("DELETE FROM profile WHERE customer_id IN (%s)", stringOfUserIds)

	if err := s.CustomerProfileRepository.Delete(ctx, query); err != nil {
		logger.Warn(fmt.Sprintf("Error deleting customer profiles: %s", err.Error()))
	}

	s.Next.Execute(ctx, removals)
}

func convertUserIdToStringByWrapperChar(elements []models.CustomerRemoval) string {
	elementsWrappedByChar := []string{}

	for _, value := range elements {
		elementsWrappedByChar = append(elementsWrappedByChar, fmt.Sprintf("'%s'", value.UserId.Hex()))
	}

	return strings.Join(elementsWrappedByChar, ",")
}
