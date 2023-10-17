package keeper

import (
	"blit/x/script/types"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetScript set a specific script in the store from its index
func (k Keeper) SetScript(ctx sdk.Context, script types.Script) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ScriptKeyPrefix))
	b := k.cdc.MustMarshal(&script)
	store.Set(types.ScriptKey(
		script.Address,
	), b)
}

// GetScript returns a script from its index
func (k Keeper) GetScript(
	ctx sdk.Context,
	address string,

) (val types.Script, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ScriptKeyPrefix))

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
		return val, true

	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveScript removes a script from the store
func (k Keeper) RemoveScript(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ScriptKeyPrefix))
	store.Delete(types.ScriptKey(
		index,
	))
}

// GetAllScript returns all script
func (k Keeper) GetAllScript(ctx sdk.Context) (list []types.Script) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ScriptKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

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

	cmd := exec.Command(
		"python3",
		"./blitvm/blitwsgi.py",
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
