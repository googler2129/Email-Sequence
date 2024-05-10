package router

import (
	"context"
	"github.com/depender/email-sequence-service/constants"
	"github.com/depender/email-sequence-service/internal/controllers/sequence"
	"github.com/depender/email-sequence-service/pkg/config"
	"github.com/depender/email-sequence-service/pkg/db"
	log "github.com/depender/email-sequence-service/pkg/logger"
	"github.com/depender/email-sequence-service/pkg/middleware"
	"github.com/depender/email-sequence-service/pkg/newrelic"
	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
)

func Initialize(ctx context.Context, s *gin.Engine) (err error) {
	//Middleware for newrelic tracking
	s.Use(nrgin.Middleware(newrelic.GetNewrelicApplication().Application))

	//Middleware for adding config to ctx
	s.Use(config.Middleware())

	s.Use(middleware.Pagination(constants.DefaultPerPageLimit))

	s.Use(log.LogMiddleware())

	sequenceRoutes := s.Group("/api/v1/sequences")
	{
		seqCtrl := sequence.SequenceWire(db.GetCluster())

		sequenceRoutes.POST("/", seqCtrl.CreateSequence)
		sequenceRoutes.PUT("/:sequence_id", seqCtrl.UpdateSequence)
		sequenceRoutes.PUT("/:sequence_id/steps/:step_id", seqCtrl.UpdateSequenceStep)
		sequenceRoutes.DELETE("/:sequence_id/steps/:step_id", seqCtrl.DeleteSequenceStep)
	}

	return
}
