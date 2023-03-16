package distributions

import "github.com/steve-care-software/libs/cryptography/hash"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewDistributionBuilder creates a new distribution builder
func NewDistributionBuilder() DistributionBuilder {
	hashAdapter := hash.NewAdapter()
	return createDistributionBuilder(hashAdapter)
}

// Builder represents the distributions builder
type Builder interface {
	Create() Builder
	WithList(list []Distribution) Builder
	Now() (Distributions, error)
}

// Distributions represents distributions
type Distributions interface {
	Hash() hash.Hash
	Power() uint
	List() []Distribution
}

// DistributionBuilder represents a distribution builder
type DistributionBuilder interface {
	Create() DistributionBuilder
	WithPower(power uint) DistributionBuilder
	WithPubKey(pubKey hash.Hash) DistributionBuilder
	Now() (Distribution, error)
}

// Distribution represents a distribution
type Distribution interface {
	Hash() hash.Hash
	Power() uint
	PubKey() hash.Hash
}
