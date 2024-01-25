package keeper

import (
	"blit/x/blit/types"
	"fmt"
	"time"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authzKeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string

		authzKeeper authzKeeper.Keeper
		// the bank keeper
		bankKeeper bankkeeper.Keeper
		// the account keeper
		accountKeeper types.AccountKeeper

		Schema collections.Schema
		// Tasks key: taskID | value: Task
		Tasks collections.Map[uint64, types.Task]
		// TaskResults key: taskID | value: TaskResult
		TaskResults collections.Map[uint64, types.TaskResult]
		// PendingTasks key: scheduledTime+taskID | value: taskID
		PendingTasks collections.Map[collections.Pair[time.Time, uint64], uint64]
		// TaskID is a counter for tasks. It tracks the next task ID to be issued.
		TaskID collections.Sequence
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,
	authzKeeper authzKeeper.Keeper,
	accountKeeper types.AccountKeeper,
	bankKeeper bankkeeper.Keeper,

) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		authority:    authority,
		logger:       logger,
		authzKeeper:  authzKeeper,

		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
	}
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
