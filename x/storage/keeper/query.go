package keeper

import (
	"blit/x/storage/types"
)

var _ types.QueryServer = Keeper{}
