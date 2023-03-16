package distributions

import (
	"errors"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []Distribution
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithList adds a list to the builder
func (app *builder) WithList(list []Distribution) Builder {
	app.list = list
	return app
}

// Now builds a new Distributions instance
func (app *builder) Now() (Distributions, error) {
	if app.list == nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Distribution in order to build a Distributions instance")
	}

	data := [][]byte{}
	for _, oneDistribution := range app.list {
		data = append(data, oneDistribution.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	power := uint(0)
	for _, oneDistribution := range app.list {
		power += oneDistribution.Power()
	}

	return createDistributions(*pHash, power, app.list), nil
}
