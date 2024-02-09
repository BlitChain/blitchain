package keeper

import (
	"blit/x/services/types"
)

var _ types.QueryServer = Keeper{}
