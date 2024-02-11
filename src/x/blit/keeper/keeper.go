package keeper

import (
	"blit/x/blit/types"
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/collections/indexes"
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/baseapp"
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
		Tasks *collections.IndexedMap[collections.Pair[sdk.AccAddress, uint64], types.Task, TasksIndexes]

		// TaskID is a counter for tasks. It tracks the next task ID to be issued.
		TaskID collections.Sequence

		FutureTasks collections.Map[string, types.FutureTask]

		Router *baseapp.MsgServiceRouter
	}
)

type TasksIndexes struct {
	Id *indexes.ReversePair[sdk.AccAddress, uint64, types.Task]
}

func (t TasksIndexes) IndexesList() []collections.Index[collections.Pair[sdk.AccAddress, uint64], types.Task] {
	return []collections.Index[collections.Pair[sdk.AccAddress, uint64], types.Task]{t.Id}
}

func newTasksIndexes(sb *collections.SchemaBuilder) TasksIndexes {
	return TasksIndexes{
		Id: indexes.NewReversePair[types.Task](
			sb, types.TaskAddressPrefix, "task_by_address_index",
			collections.PairKeyCodec(sdk.AccAddressKey, collections.Uint64Key),
			indexes.WithReversePairUncheckedValue(),
		),
	}
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,
	authzKeeper authzKeeper.Keeper,
	accountKeeper types.AccountKeeper,
	bankKeeper bankkeeper.Keeper,
	router *baseapp.MsgServiceRouter,

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

		Tasks: collections.NewIndexedMap(sb, types.TasksKeyPrefix, "tasks", collections.PairKeyCodec(sdk.AccAddressKey, collections.Uint64Key), codec.CollValue[types.Task](cdc), newTasksIndexes(sb)),

		TaskID: collections.NewSequence(sb, types.TaskIDKey, "taskID"),

		FutureTasks: collections.NewMap(sb, types.FutureTasksKeyPrefix, "futureTasks", collections.StringKey, codec.CollValue[types.FutureTask](cdc)),

		Router: router,
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
