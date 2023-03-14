package distributions

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type distributionBuilder struct {
	hashAdapter hash.Adapter
	power       uint
	pubKey      hash.Hash
}

func createDistributionBuilder(
	hashAdapter hash.Adapter,
) DistributionBuilder {
	out := distributionBuilder{
		hashAdapter: hashAdapter,
		power:       0,
		pubKey:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *distributionBuilder) Create() DistributionBuilder {
	return createDistributionBuilder(app.hashAdapter)
}

// WithPower adds a power to the builder
func (app *distributionBuilder) WithPower(power uint) DistributionBuilder {
	app.power = power
	return app
}

// WithPubKey adds a pubKey to the builder
func (app *distributionBuilder) WithPubKey(pubKey hash.Hash) DistributionBuilder {
	app.pubKey = pubKey
	return app
}

// Now builds a new Distribution instance
func (app *distributionBuilder) Now() (Distribution, error) {
	if app.power <= 0 {
		return nil, errors.New("the power must be greater than zero (0) in order to build a Distribution instance")
	}

	if app.pubKey == nil {
		return nil, errors.New("the pubKey is mandatory in order to build a Distribution instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(fmt.Sprintf("%d", app.power)),
		app.pubKey.Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createDistribution(*pHash, app.power, app.pubKey), nil
}
