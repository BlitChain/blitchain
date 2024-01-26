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

	sb := collections.NewSchemaBuilder(storeService)

	k := Keeper{
		cdc:          cdc,
		storeService: storeService,
		authority:    authority,
		logger:       logger,
		authzKeeper:  authzKeeper,

		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,

		// 		Proposals:              collections.NewMap(sb, types.ProposalsKeyPrefix, "proposals", collections.Uint64Key, codec.CollValue[v1.Proposal](cdc)),
		Tasks:  collections.NewMap(sb, types.TasksKeyPrefix, "tasks", collections.Uint64Key, codec.CollValue[types.Task](cdc)),
		TaskID: collections.NewSequence(sb, types.TaskIDKey, "taskID"),
	}
	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}
	k.Schema = schema
	return k

}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
