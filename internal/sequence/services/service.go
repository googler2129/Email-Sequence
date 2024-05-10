package services

import (
	"context"
	"github.com/depender/email-sequence-service/internal/domain/interfaces"
	"github.com/depender/email-sequence-service/internal/domain/models"
	"github.com/depender/email-sequence-service/internal/sequence/requests"
	"github.com/guregu/null/v5"
	"sync"
)

type Service struct {
	repo interfaces.SequenceRepo
}

var (
	svcOnce sync.Once
	svc     *Service
)

func NewService(repo interfaces.SequenceRepo) *Service {
	svcOnce.Do(func() {
		svc = &Service{
			repo: repo,
		}
	})

	return svc
}

func (s *Service) CreateSequence(ctx context.Context, req *requests.CreateSeqSvcReq) error {
	sequenceSteps := make([]*models.SequenceStep, 0)
	for _, step := range req.SequenceSteps {
		sequenceSteps = append(sequenceSteps, &models.SequenceStep{
			Subject:     step.Subject,
			Content:     step.Content,
			WaitingDays: step.WaitingDays,
			SerialOrder: step.SerialOrder,
			CreatedBy:   null.StringFrom(req.UserName),
			CreatedByID: null.StringFrom(req.UserID),
		})
	}

	sequence := &models.Sequence{
		Name:                 req.Name,
		ClickTrackingEnabled: req.IsClickTrackingEnabled,
		OpenTrackingEnabled:  req.IsOpenTrackingEnabled,
		SequenceSteps:        sequenceSteps,
		CreatedBy:            null.StringFrom(req.UserName),
		CreatedByID:          null.StringFrom(req.UserID),
	}

	err := s.repo.CreateSequence(ctx, sequence)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateSequence(ctx context.Context, req *requests.UpdateSeqSvcReq) error {
	sequence, err := s.repo.GetSequence(ctx, map[string]interface{}{
		"id": req.SequenceID,
	})
	if err != nil {
		return err
	}

	err = s.repo.UpdateSequencesByCondition(ctx, map[string]interface{}{
		"id": sequence.ID,
	}, map[string]interface{}{
		"open_tracking_enabled":  req.IsOpenTrackingEnabled,
		"click_tracking_enabled": req.IsClickTrackingEnabled,
		"updated_by":             null.StringFrom(req.UserName),
		"updated_by_id":          null.StringFrom(req.UserID),
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateSequenceStep(ctx context.Context, req *requests.UpdateSeqStepSvcReq) error {
	sequence, err := s.repo.GetSequence(ctx, map[string]interface{}{
		"id": req.SequenceID,
	})
	if err != nil {
		return err
	}

	sequenceStep, err := s.repo.GetSequenceStep(ctx, map[string]interface{}{
		"id":          req.StepID,
		"sequence_id": sequence.ID,
	})
	if err != nil {
		return err
	}

	err = s.repo.UpdateSequenceStepsByCondition(ctx, map[string]interface{}{
		"id": sequenceStep.ID,
	}, map[string]interface{}{
		"subject":       req.Subject,
		"content":       req.Content,
		"updated_by":    null.StringFrom(req.UserName),
		"updated_by_id": null.StringFrom(req.UserID),
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteSequenceStep(ctx context.Context, sequenceID, stepID uint64) error {
	sequence, err := s.repo.GetSequence(ctx, map[string]interface{}{
		"id": sequenceID,
	})
	if err != nil {
		return err
	}

	sequenceStep, err := s.repo.GetSequenceStep(ctx, map[string]interface{}{
		"id":          stepID,
		"sequence_id": sequence.ID,
	})
	if err != nil {
		return err
	}

	err = s.repo.DeleteSequenceStepsByCondition(ctx, map[string]interface{}{
		"id": sequenceStep.ID,
	})
	if err != nil {
		return err
	}

	nextSequenceSteps, err := s.repo.GetSequenceStepsByCondition(ctx, map[string]interface{}{
		"serial_order": sequenceStep.SerialOrder + 1,
	})

	if len(nextSequenceSteps) > 0 {
		nextSequenceStep := nextSequenceSteps[0]
		err = s.repo.UpdateSequenceStepsByCondition(ctx, map[string]interface{}{
			"id": nextSequenceStep.ID,
		}, map[string]interface{}{
			"waiting_days": nextSequenceStep.WaitingDays + sequenceStep.WaitingDays,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
