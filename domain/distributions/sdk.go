package distributions

import "github.com/steve-care-software/libs/cryptography/hash"

// Builder represents the distributions builder
type Builder interface {
	Create() Builder
	WithList(list []Distribution) Builder
	Now() (Distributions, error)
}

// Distributions represents distributions
type Distributions interface {
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
