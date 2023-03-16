package distributions

import "github.com/steve-care-software/libs/cryptography/hash"

type distributions struct {
	hash  hash.Hash
	power uint
	list  []Distribution
}

func createDistributions(
	hash hash.Hash,
	power uint,
	list []Distribution,
) Distributions {
	out := distributions{
		hash:  hash,
		power: power,
		list:  list,
	}

	return &out
}

// Hash returns the hash
func (obj *distributions) Hash() hash.Hash {
	return obj.hash
}

// Power returns the power of a distribution's list
func (obj *distributions) Power() uint {
	return obj.power
}

// List returns the list
func (obj *distributions) List() []Distribution {
	return obj.list
}
