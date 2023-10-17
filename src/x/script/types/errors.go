package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/script module sentinel errors
var (
	ErrRun                          = sdkerrors.Register(ModuleName, 1100, "Run error")
	ErrorExtraLinesOrExternalScript = sdkerrors.Register(ModuleName, 1102, "cannot use extralines")
	ScriptError                     = sdkerrors.Register(ModuleName, 1103, "Exception in Script")
	MsgError                        = sdkerrors.Register(ModuleName, 1104, "Error in chain Msg call")
	QueryError                      = sdkerrors.Register(ModuleName, 1105, "Error in chain Query call")
	ErrInvalidSigner                = sdkerrors.Register(ModuleName, 1106, "expected gov account as only signer for proposal message")
)
