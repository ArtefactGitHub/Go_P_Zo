package platform

import (
	"crypto/rand"
	"io"
	"time"

	ulid "github.com/oklog/ulid/v2"
)

// https://zenn.dev/emiksk/articles/e2716c0af75eea
type Identifier struct {
	identifier string
}

type IdentifierGenerator interface {
	Generate() Identifier
}

var defaultGenerator IdentifierGenerator

func init() {
	defaultGenerator = newULIDGenerator(rand.Reader)
}

func GenerateIdentifier() Identifier {
	return defaultGenerator.Generate()
}

func NewIdentifier(id string) Identifier {
	return Identifier{
		identifier: id,
	}
}

func (i Identifier) Value() string {
	return i.identifier
}

func (i Identifier) Equal(other Identifier) bool {
	return i.identifier == other.identifier
}

type ULIDGenerator struct {
	entropy *ulid.MonotonicEntropy
}

func newULIDGenerator(reader io.Reader) *ULIDGenerator {
	return &ULIDGenerator{
		entropy: ulid.Monotonic(reader, 0),
	}
}

func (g *ULIDGenerator) Generate() Identifier {
	id := ulid.MustNew(ulid.Timestamp(time.Now()), g.entropy)

	return Identifier{
		identifier: id.String(),
	}
}
