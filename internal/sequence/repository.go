package sequence

import (
	"context"
	"github.com/depender/email-sequence-service/internal/domain/models"
	"github.com/depender/email-sequence-service/pkg/db"
	"sync"
)

type Repository struct {
	Db *db.Db
}

var (
	repoOnce sync.Once
	repo     *Repository
)

func NewRepository(db *db.Db) *Repository {
	repoOnce.Do(func() {
		repo = &Repository{
			Db: db,
		}
	})

	return repo
}

func (r *Repository) CreateSequence(ctx context.Context, sequence *models.Sequence) error {
	result := r.Db.Create(sequence)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *Repository) GetSequence(ctx context.Context, conditions map[string]interface{}) (sequence *models.Sequence, err error) {
	result := r.Db.Model(&models.Sequence{}).Where(conditions).First(&sequence)
	if result.Error != nil {
		err = result.Error
		return
	}

	return
}

func (r *Repository) GetSequenceStep(ctx context.Context, conditions map[string]interface{}) (step *models.SequenceStep, err error) {
	result := r.Db.Model(&models.SequenceStep{}).Where(conditions).First(&step)
	if result.Error != nil {
		err = result.Error
		return
	}

	return
}

func (r *Repository) GetSequenceStepsByCondition(ctx context.Context, conditions map[string]interface{}) (steps []*models.SequenceStep, err error) {
	result := r.Db.Model(&models.SequenceStep{}).Where(conditions).Find(&steps)
	if result.Error != nil {
		err = result.Error
		return
	}

	return
}

func (r *Repository) UpdateSequencesByCondition(ctx context.Context, conditions map[string]interface{}, updates map[string]interface{}) error {
	result := r.Db.Model(&models.Sequence{}).Where(conditions).Updates(updates)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *Repository) UpdateSequenceStepsByCondition(ctx context.Context, conditions map[string]interface{}, updates map[string]interface{}) error {
	result := r.Db.Model(&models.SequenceStep{}).Where(conditions).Updates(updates)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *Repository) DeleteSequenceStepsByCondition(ctx context.Context, conditions map[string]interface{}) error {
	result := r.Db.Where(conditions).Delete(&models.SequenceStep{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
