package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyGasPerChar = []byte("GasPerChar")
	// TODO: Determine the default value
	DefaultGasPerChar string = "gas_per_char"
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	gasPerChar string,
) Params {
	return Params{
		GasPerChar: gasPerChar,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultGasPerChar,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyGasPerChar, &p.GasPerChar, validateGasPerChar),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateGasPerChar(p.GasPerChar); err != nil {
		return err
	}

	return nil
}

// validateGasPerChar validates the GasPerChar param
func validateGasPerChar(v interface{}) error {
	gasPerChar, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = gasPerChar

	return nil
}
