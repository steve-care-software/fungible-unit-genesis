package distributions

import "github.com/steve-care-software/libs/cryptography/hash"

type distribution struct {
	hash   hash.Hash
	power  uint
	pubKey hash.Hash
}

func createDistribution(
	hash hash.Hash,
	power uint,
	pubKey hash.Hash,
) Distribution {
	out := distribution{
		hash:   hash,
		power:  power,
		pubKey: pubKey,
	}

	return &out
}

// Hash returns the hash
func (obj *distribution) Hash() hash.Hash {
	return obj.hash
}

// Power returns the power
func (obj *distribution) Power() uint {
	return obj.power
}

// PubKey returns the pubKey
func (obj *distribution) PubKey() hash.Hash {
	return obj.pubKey
}
