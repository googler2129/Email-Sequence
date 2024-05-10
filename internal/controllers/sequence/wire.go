//go:build wireinject
// +build wireinject

package sequence

import (
	"github.com/depender/email-sequence-service/pkg/db"
	"github.com/google/wire"
)

func SequenceWire(db *db.Db) *Controller {
	panic(wire.Build(SequenceProviderSet))
}
