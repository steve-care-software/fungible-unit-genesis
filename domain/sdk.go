package genesis

import (
	"time"

	"github.com/steve-care-software/fungible-unit-genesis/domain/distributions"
	"github.com/steve-care-software/libs/cryptography/hash"
)

// Symbol represents the symbol type
type Symbol [4]byte

// Builder represents a genesis builder
type Builder interface {
	Create() Builder
	WithSymbol(symbol Symbol) Builder
	WithVestingPeriod(vestingPeriod uint) Builder
	WithDistributions(distributions distributions.Distributions) Builder
	Now() (Genesis, error)
}

// Genesis represents the genesis instance
type Genesis interface {
	Hash() hash.Hash
	Symbol() Symbol
	VestingPeriod() time.Duration
	Distributions() distributions.Distributions
	TotalSupply() uint
}
