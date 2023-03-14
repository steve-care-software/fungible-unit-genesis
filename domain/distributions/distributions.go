package distributions

import "github.com/steve-care-software/libs/cryptography/hash"

type distributions struct {
	hash hash.Hash
	list []Distribution
}

func createDistributions(
	hash hash.Hash,
	list []Distribution,
) Distributions {
	out := distributions{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *distributions) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *distributions) List() []Distribution {
	return obj.list
}
