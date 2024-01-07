package keeper

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	blittypes "blit/x/blit/types"
	"blit/x/script/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/viper"
)

// SetScript set a specific script in the store from its index
func (k Keeper) SetScript(ctx context.Context, script types.Script) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ScriptKeyPrefix))
	b := k.cdc.MustMarshal(&script)
	store.Set(types.ScriptKey(
		script.Address,
	), b)
}

// GetScript returns a script from its index
func (k Keeper) GetScript(
	ctx context.Context,
	address string,

) (val types.Script, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ScriptKeyPrefix))

	b := store.Get(types.ScriptKey(
		address,
	))
	if b == nil {
		val = types.Script{
			Address: address,
			Code:    "",
			Version: 0,
		}
		k.SetScript(ctx, val)
		return val, false

	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveScript removes a script from the store
func (k Keeper) RemoveScript(
	ctx context.Context,
	address string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ScriptKeyPrefix))
	store.Delete(types.ScriptKey(
		address,
	))
}

// GetAllScript returns all script
func (k Keeper) GetAllScript(ctx context.Context) (list []types.Script) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ScriptKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Script
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) RunWeb(goCtx context.Context, index string, httpreq string) (val string, found bool) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	valFound, found := k.GetScript(ctx, index)
	if !found {
		return "", false
	}
	port, srv, err := k.NewRPCServer(goCtx, valFound.Address)
	if err != nil {
		fmt.Printf("NewRPCServer error: %s\n", err)
		return "", false
	}

	blockInfoJson, err := json.Marshal(ctx.BlockHeader())
	if err != nil {
		log.Fatal(err)
	}

	blitvmPath := viper.GetString(blittypes.FlagBlitVMPath)
	pyenv_root := os.Getenv("PYENV_ROOT")
	pythonExe := filepath.Join(pyenv_root, "versions", "blit-python", "bin", "python")

	cmd := exec.Command(
		pythonExe, filepath.Join(blitvmPath, "blitwsgi.py"),
		port,
		valFound.Address,
		string(blockInfoJson),
		valFound.Code,
		httpreq,
	)

	cmd.Stdin = strings.NewReader(httpreq)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmdErr := cmd.Run()

	if err := srv.Shutdown(context.Background()); err != nil {
		fmt.Printf("shutdown error")
		log.Fatal(err)
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())

	fmt.Println("server stopped")
	fmt.Printf("out: %s\n", outStr)
	fmt.Printf("err: %s\n", errStr)
	if cmdErr != nil {
		return fmt.Sprintf("HTTP/1.1 500\ncontent-type: text/plain\n\n%+v\nout: %s\nerr:%s", cmdErr, outStr, errStr), true
		// log.Fatal(cmdErr)
	}

	return string(outStr[:]), true
}
