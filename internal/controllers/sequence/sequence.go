package sequence

import (
	requests3 "github.com/depender/email-sequence-service/internal/controllers/sequence/requests"
	"github.com/depender/email-sequence-service/internal/domain/interfaces"
	requests2 "github.com/depender/email-sequence-service/internal/sequence/requests"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
	"sync"
)

type Controller struct {
	sequenceSvc interfaces.SequenceService
}

var (
	ctrl     *Controller
	ctrlOnce sync.Once
)

func NewController(sequenceSvc interfaces.SequenceService) *Controller {
	ctrlOnce.Do(func() {
		ctrl = &Controller{
			sequenceSvc: sequenceSvc,
		}
	})

	return ctrl
}

func (c *Controller) CreateSequence(ctx *gin.Context) {
	var ctrlReq *requests3.CreateSeqCtrlReq
	err := ctx.ShouldBindJSON(&ctrlReq)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	svcReq, err := validateCreateSequence(ctrlReq)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	err = c.sequenceSvc.CreateSequence(ctx, svcReq)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, "Success")
}

func (c *Controller) UpdateSequence(ctx *gin.Context) {
	var ctrlReq *requests3.UpdateSeqCtrlReq
	err := ctx.ShouldBindJSON(&ctrlReq)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	sequenceID, err := strconv.ParseUint(ctx.Param("sequence_id"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	svcReq, err := validateUpdateSequence(ctrlReq, sequenceID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	err = c.sequenceSvc.UpdateSequence(ctx, svcReq)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, "Success")
}

func (c *Controller) UpdateSequenceStep(ctx *gin.Context) {
	var ctrlReq *requests3.UpdateSeqStepCtrlReq
	err := ctx.ShouldBindJSON(&ctrlReq)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	sequenceID, err := strconv.ParseUint(ctx.Param("sequence_id"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	stepID, err := strconv.ParseUint(ctx.Param("step_id"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	svcReq, err := validateUpdateSequenceStep(ctrlReq, sequenceID, stepID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	err = c.sequenceSvc.UpdateSequenceStep(ctx, svcReq)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, "Success")
}

func (c *Controller) DeleteSequenceStep(ctx *gin.Context) {
	sequenceID, err := strconv.ParseUint(ctx.Param("sequence_id"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	stepID, err := strconv.ParseUint(ctx.Param("step_id"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	err = c.sequenceSvc.DeleteSequenceStep(ctx, sequenceID, stepID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, "Success")
}

func validateUpdateSequence(req *requests3.UpdateSeqCtrlReq, sequenceID uint64) (svcReq *requests2.UpdateSeqSvcReq, err error) {
	err = validator.New().Struct(req)
	if err != nil {
		return
	}

	svcReq = &requests2.UpdateSeqSvcReq{
		IsOpenTrackingEnabled:  req.IsOpenTrackingEnabled,
		IsClickTrackingEnabled: req.IsClickTrackingEnabled,
		SequenceID:             sequenceID,
		UserID:                 req.UserID,
		UserName:               req.UserName,
	}

	return
}

func validateCreateSequence(req *requests3.CreateSeqCtrlReq) (svcReq *requests2.CreateSeqSvcReq, err error) {
	err = validator.New().Struct(req)
	if err != nil {
		return
	}

	steps := make([]requests2.SequenceStep, 0)

	for _, step := range req.SequenceSteps {
		steps = append(steps, requests2.SequenceStep{
			Subject:     step.Subject,
			Content:     step.Content,
			WaitingDays: step.WaitingDays,
			SerialOrder: step.SerialOrder,
		})
	}

	svcReq = &requests2.CreateSeqSvcReq{
		Name:                   req.Name,
		IsOpenTrackingEnabled:  req.IsOpenTrackingEnabled,
		IsClickTrackingEnabled: req.IsClickTrackingEnabled,
		SequenceSteps:          steps,
		UserID:                 req.UserID,
		UserName:               req.UserName,
	}

	return
}

func validateUpdateSequenceStep(req *requests3.UpdateSeqStepCtrlReq, sequenceID, stepID uint64) (svcReq *requests2.UpdateSeqStepSvcReq, err error) {
	err = validator.New().Struct(req)
	if err != nil {
		return
	}

	svcReq = &requests2.UpdateSeqStepSvcReq{
		Subject:    req.Subject,
		Content:    req.Content,
		SequenceID: sequenceID,
		StepID:     stepID,
		UserID:     req.UserID,
		UserName:   req.UserName,
	}

	return
}
