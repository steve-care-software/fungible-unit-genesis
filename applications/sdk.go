package applications

import (
	genesis "github.com/steve-care-software/fungible-unit-genesis/domain"
)

// Application represents a fungible unit's application
type Application interface {
	List() ([]genesis.Symbol, error)
	Retrieve(symbol genesis.Symbol) (genesis.Genesis, error)
	New(genesis genesis.Genesis) error
	Delete(symbol genesis.Symbol) error
}
