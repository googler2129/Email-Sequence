package sequence

import (
	"github.com/depender/email-sequence-service/internal/domain/interfaces"
	"github.com/depender/email-sequence-service/internal/sequence"
	"github.com/depender/email-sequence-service/internal/sequence/services"
	"github.com/google/wire"
)

var SequenceProviderSet wire.ProviderSet = wire.NewSet(
	NewController,
	services.NewService,
	sequence.NewRepository,

	// bind each one of the interfaces
	wire.Bind(new(interfaces.SequenceController), new(*Controller)),
	wire.Bind(new(interfaces.SequenceService), new(*services.Service)),
	wire.Bind(new(interfaces.SequenceRepo), new(*sequence.Repository)),
)
