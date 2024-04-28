package utils

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/tuanloc1105/go-common-lib/model"
)

func GenerateNewBaseEntity(ctx context.Context) model.BaseEntity {
	var currentUsernameInsertOrUpdateData = GetCurrentUsernameFromContextForInsertOrUpdateDataInDb(ctx)
	return model.BaseEntity{
		Active:    GetPointerOfAnyValue(true),
		UUID:      uuid.New().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: currentUsernameInsertOrUpdateData,
		UpdatedBy: currentUsernameInsertOrUpdateData,
	}
}

func ChangeDataOfBaseEntityForUpdate(ctx context.Context, baseEntity *model.BaseEntity) {
	var currentUsernameInsertOrUpdateData = GetCurrentUsernameFromContextForInsertOrUpdateDataInDb(ctx)
	baseEntity.UpdatedAt = time.Now()
	baseEntity.UpdatedBy = currentUsernameInsertOrUpdateData
}
