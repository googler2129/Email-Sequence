package interfaces

import (
	"context"
	"github.com/depender/email-sequence-service/internal/domain/models"
	"github.com/depender/email-sequence-service/internal/sequence/requests"
	"github.com/gin-gonic/gin"
)

type (
	SequenceController interface {
		CreateSequence(ctx *gin.Context)
		UpdateSequence(ctx *gin.Context)
		UpdateSequenceStep(ctx *gin.Context)
		DeleteSequenceStep(ctx *gin.Context)
	}

	SequenceService interface {
		CreateSequence(ctx context.Context, req *requests.CreateSeqSvcReq) error
		UpdateSequence(ctx context.Context, req *requests.UpdateSeqSvcReq) error
		UpdateSequenceStep(ctx context.Context, req *requests.UpdateSeqStepSvcReq) error
		DeleteSequenceStep(ctx context.Context, sequenceID, stepID uint64) error
	}

	SequenceRepo interface {
		CreateSequence(ctx context.Context, sequence *models.Sequence) error
		GetSequence(ctx context.Context, conditions map[string]interface{}) (sequence *models.Sequence, err error)
		GetSequenceStep(ctx context.Context, conditions map[string]interface{}) (step *models.SequenceStep, err error)
		GetSequenceStepsByCondition(ctx context.Context, conditions map[string]interface{}) (steps []*models.SequenceStep, err error)
		UpdateSequencesByCondition(ctx context.Context, conditions map[string]interface{}, updates map[string]interface{}) error
		UpdateSequenceStepsByCondition(ctx context.Context, conditions map[string]interface{}, updates map[string]interface{}) error
		DeleteSequenceStepsByCondition(ctx context.Context, conditions map[string]interface{}) error
	}
)
