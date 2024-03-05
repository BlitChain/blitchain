package keeper

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"cosmossdk.io/core/store"
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authzKeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	"github.com/cosmos/cosmos-sdk/x/group/errors"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"

	blittypes "blit/x/blit/types"
	"blit/x/script/types"

	"github.com/gorilla/rpc"
	rpcjson "github.com/gorilla/rpc/json"
	"google.golang.org/grpc"
)

type (
	Keeper struct {
		cdc          codec.Codec
		storeService store.KVStoreService
		logger       log.Logger

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority    string
		Router       *baseapp.MsgServiceRouter
		Querier      *baseapp.GRPCQueryRouter
		AuthzKeeper  authzKeeper.Keeper
		accKeeper    types.AccountKeeper
		currentDepth int
	}
)

func NewKeeper(
	cdc codec.Codec,
	storeService store.KVStoreService,
	logger log.Logger,
	router *baseapp.MsgServiceRouter,
	querier *baseapp.GRPCQueryRouter,
	authzKeeper authzKeeper.Keeper,
	accKeeper types.AccountKeeper,
	authority string,

) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		logger:       logger,
		authority:    authority,
		Router:       router,
		Querier:      querier,
		AuthzKeeper:  authzKeeper,
		accKeeper:    accKeeper,
		currentDepth: 0,
	}
}

type EvalScriptContext struct {
	Messages      []sdk.Msg
	CallerAddress string
	ScriptAddress string
	FunctionName  string
	Kwargs        string
	Code          string
	ExtraCode     string
}

type EvalScriptResponse struct {
	Response string
}

type RpcService struct {
	k             *Keeper
	m             *msgServer
	goCtx         context.Context
	ScriptAddress string
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (keeper Keeper) MsgServiceRouter() *baseapp.MsgServiceRouter {
	return keeper.Router
}

// return the script keeper's grpc query router
func (keeper Keeper) GRPCQueryRouter() *baseapp.GRPCQueryRouter {
	return keeper.Querier
}

func (k Keeper) ServiceDesc(ctx sdk.Context, name string) (serviceDesc *grpc.ServiceDesc) {

	availableSeviceNames := map[string]bool{
		"blit.blit.Query":                                          true,
		"blit.script.Query":                                        true,
		"blit.storage.Query":                                       true,
		"cosmos.auth.v1beta1.Query":                                true,
		"cosmos.authz.v1beta1.Query":                               true,
		"cosmos.bank.v1beta1.Query":                                true,
		"cosmos.consensus.v1.Query":                                true,
		"cosmos.distribution.v1beta1.Query":                        true,
		"cosmos.evidence.v1beta1.Query":                            true,
		"cosmos.feegrant.v1beta1.Query":                            true,
		"cosmos.gov.v1.Query":                                      true,
		"cosmos.gov.v1beta1.Query":                                 true,
		"cosmos.group.v1.Query":                                    true,
		"cosmos.mint.v1beta1.Query":                                true,
		"cosmos.params.v1beta1.Query":                              true,
		"cosmos.slashing.v1beta1.Query":                            true,
		"cosmos.staking.v1beta1.Query":                             true,
		"cosmos.upgrade.v1beta1.Query":                             true,
		"ibc.applications.fee.v1.Query":                            true,
		"ibc.applications.interchain_accounts.controller.v1.Query": true,
		"ibc.applications.interchain_accounts.host.v1.Query":       true,
		"ibc.applications.transfer.v1.Query":                       true,
		"ibc.core.channel.v1.Query":                                true,
		"ibc.core.client.v1.Query":                                 true,
		"ibc.core.connection.v1.Query":                             true,
		// No no
		"cosmos.app.v1alpha1.Query":                        false,
		"cosmos.autocli.v1.Query":                          false,
		"cosmos.base.node.v1beta1.Service":                 false,
		"cosmos.base.reflection.v1beta1.ReflectionService": false,
		"cosmos.base.tendermint.v1beta1.Service":           false,
		"cosmos.reflection.v1.ReflectionService":           false,
		"cosmos.tx.v1beta1.Service":                        false,
		"testpb.Query":                                     false,
	}
	// Get the value and type of the qrt instance
	value := reflect.ValueOf(k.Querier).Elem()

	// Get the private field serviceData
	field := value.FieldByName("serviceData")

	// Make the private field accessible
	field = reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()

	// Get the length of the serviceData slice
	length := field.Len()

	// log the length of serviceData
	k.Logger(ctx).Info("ServiceDesc", "length", length)

	// Loop through the serviceData slice
	for i := 0; i < length; i++ {
		// Get the serviceDesc field from the sd struct
		serviceDescField := field.Index(i).FieldByName("serviceDesc")

		// Make the private field accessible
		serviceDescField = reflect.NewAt(serviceDescField.Type(), unsafe.Pointer(serviceDescField.UnsafeAddr())).Elem()

		// Get the serviceDesc value
		sd := serviceDescField.Interface().(*grpc.ServiceDesc)
		// log the serviceDesc.ServiceName
		//k.Logger(ctx).Info("ServiceDesc", "sd.ServiceName", sd.ServiceName)

		// Check if the serviceDesc's ServiceName is a prefix of the name
		if strings.HasPrefix(name, "/"+sd.ServiceName) && availableSeviceNames[sd.ServiceName] {
			// Return the serviceDesc as a *grpc.ServiceDesc
			serviceDesc = sd
		}
	}
	return serviceDesc
}

func (k Keeper) HandleMsg(ctx sdk.Context, jsonMsg string, signer string) (res string, err error) {
	ctx, Write := ctx.CacheContext()

	defer func() {
		if r := recover(); r != nil {
			err = errorsmod.Wrapf(types.MsgError, "Error in HandleMsg: %v", r)
			k.Logger(ctx).Error("Error in HandleMsg", "error", r)
			res = string("oh no")
		} else {
			Write()
		}

	}()

	var anyJSON json.RawMessage
	err = json.Unmarshal([]byte(jsonMsg), &anyJSON)
	if err != nil {
		return "", errorsmod.Wrapf(types.MsgError, "could not unmarshal JSON: %v", err)
	}

	var m sdk.Msg
	err = k.cdc.UnmarshalInterfaceJSON(anyJSON, &m)
	if err != nil {
		return "", errorsmod.Wrapf(types.MsgError, "could not unmarshal message: %v", err)
	}

	err = k.ensureMsgAuthZ([]sdk.Msg{m}, signer)
	if err != nil {
		return "", err
	}
	// Validate the message
	// err = m.ValidateBasic()
	// if err != nil {
	//	return "", errorsmod.Wrap(err, "validation failed")
	// }

	k.Logger(ctx).Info("Run", "msg", m)

	handler := k.Router.Handler(m)
	if handler == nil {
		return "", errorsmod.Wrapf(types.MsgError, "no message handler found for %q", sdk.MsgTypeURL(m))
	}

	r, err := handler(ctx, m)
	if err != nil {
		return "", errorsmod.Wrapf(err, "handler error")
	}

	if r == nil {
		return "", errorsmod.Wrapf(types.MsgError, "no result for %q", sdk.MsgTypeURL(m))
	}

	ctx.EventManager().EmitEvents(r.GetEvents())
	// marshal the result
	resbz, err := k.cdc.MarshalJSON(r)
	if err != nil {
		return "", errorsmod.Wrapf(types.MsgError, "could not marshal result: %v", err)
	}

	res = string(resbz)

	Write()
	// return the result
	return res, err
}

func (k Keeper) HandleQuery(ctx sdk.Context, method string, json_args string) (string, error) {

	// create a new cached context
	ctx, _ = ctx.CacheContext()

	k.Logger(ctx).Info("HandleQuery", "method", method, "json_args", json_args)
	// get the handler for the method
	if k.Querier == nil {
		k.Logger(ctx).Error("Querier not found")

		return "", errorsmod.Wrapf(types.QueryError, "no querier defined for module %s", types.ModuleName)
	}
	handler := k.Querier.Route(method)
	if handler == nil {
		k.Logger(ctx).Error("Handler not found", "method", method)
		return "", errorsmod.Wrapf(types.QueryError, "no query handler found for %s", method)
	}

	// access the k.Querier private serviceDesc array
	desc := k.ServiceDesc(ctx, method)
	if desc == nil {
		k.Logger(ctx).Error("ServiceDesc not found", "method", method)
		return "", fmt.Errorf("ServiceDesc not found for method: %s", method)
	}

	k.Logger(ctx).Info("name", "method", method, "desc.ServiceName", desc.ServiceName)
	// check if the method is a prefix of the desc.ServiceName
	if strings.HasPrefix(method, "/"+desc.ServiceName) {
		handlerType := desc.HandlerType
		handlerElemType := reflect.TypeOf(handlerType).Elem()

		for i := 0; i < handlerElemType.NumMethod(); i++ {
			methodInfo := handlerElemType.Method(i)
			k.Logger(ctx).Info("methodInfo.Name", "methodInfo.Name", methodInfo.Name)

			if fmt.Sprintf("/%s/%s", desc.ServiceName, methodInfo.Name) == method {
				inputType := methodInfo.Type.In(1)
				input := reflect.New(inputType.Elem()).Interface()

				// check if the input is a proto.Message
				protoInput, ok := input.(codec.ProtoMarshaler)
				if !ok {
					k.Logger(ctx).Error("Input is not a valid proto.Message", "input", input)
					return "", fmt.Errorf("input is not a valid proto.Message")
				}
				// Unmarshal json_args to protoInput
				err := k.cdc.UnmarshalJSON([]byte(json_args), protoInput)

				if err != nil {
					exampleJson, _ := k.cdc.MarshalJSON(protoInput)
					k.Logger(ctx).Error("Error unmarshalling json_args", "error", err)
					return "", fmt.Errorf("error unmarshalling json_args[%s]: %v into %s", json_args, err, string(exampleJson))
				}

				// Marshal protoInput to bytes.
				protoInputBytes, err := k.cdc.Marshal(protoInput)
				if err != nil {
					return "", errorsmod.Wrapf(types.ErrRun, "could not marshal protoInput: %#v", err)
				}
				// Create the request
				request := abci.RequestQuery{Data: protoInputBytes}

				// Execute the handler
				response, err := handler(ctx, &request)
				if err != nil {
					k.Logger(ctx).Error("Error executing handler", "error", err)
					return "", fmt.Errorf("error executing handler %#v error: %v", protoInput, err)
				}

				// Unmarshal response to proto
				responseType := methodInfo.Type.Out(0)
				responseProto := reflect.New(responseType.Elem()).Interface().(codec.ProtoMarshaler)
				err = k.cdc.Unmarshal(response.Value, responseProto)

				if err != nil {
					k.Logger(ctx).Error("Error unmarshalling response", "error", err)
					return "", fmt.Errorf("error unmarshalling response: %v", err)
				}

				// Marshal response to JSON
				responseJSON, err := k.cdc.MarshalJSON(responseProto)
				if err != nil {
					k.Logger(ctx).Error("Error marshalling response", "error", err)
					return "", fmt.Errorf("error marshalling response: %v", err)
				}

				k.Logger(ctx).Info("Response", "response", string(responseJSON))

				return string(responseJSON), nil
			}
		}
	}

	k.Logger(ctx).Error("Method not found", "method", method)
	return "", fmt.Errorf("method not found: %s", method)
}

// ensureMsgAuthZ checks that if a message requires signers that all of them
// are equal to the given signer/grantee calling the script.
func (k Keeper) ensureMsgAuthZ(msgs []sdk.Msg, signer string) error {
	// In practice, GetSigners() should return a non-empty array without
	// duplicates.
	address, err := k.accKeeper.AddressCodec().StringToBytes(signer)
	if err != nil {
		return err
	}

	for i := range msgs {

		signers, _, err := k.cdc.GetMsgV1Signers(msgs[i])
		if err != nil {
			return err
		}
		for _, acct := range signers {
			if !bytes.Equal(address, acct) {
				return errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "msg [%d] does not have authorization; expected [%s], got [%s]", i, signer, acct)
			}
		}
	}
	return nil
}

type ResultException struct {
	Msg    string `json:"msg"`
	Lineno int    `json:"lineno"`
	Col    int    `json:"col"`
}
type EvalResult struct {
	NodesCalled uint64          `json:"nodes_called"`
	CumSize     uint64          `json:"cumsize"`
	Exception   ResultException `json:"exception"`
}

func (k Keeper) NewRPCServer(goCtx context.Context, address string) (string, *http.Server, error) {
	s := rpc.NewServer()
	s.RegisterCodec(rpcjson.NewCodec(), "application/json")
	s.RegisterCodec(rpcjson.NewCodec(), "application/json;charset=UTF-8")
	rpcservice := new(RpcService)
	rpcservice.k = &k
	rpcservice.m = &msgServer{Keeper: k}
	rpcservice.goCtx = goCtx
	rpcservice.ScriptAddress = address

	s.RegisterService(rpcservice, "")
	r := mux.NewRouter()
	r.Handle("/rpc", s)
	srv := &http.Server{Handler: r}

	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		panic(err)
	}

	//fmt.Println("Using port:", listener.Addr().(*net.TCPAddr).Port)

	go func() {
		//fmt.Println("start ListenAndServe")
		srv.Serve(listener)
		//fmt.Println("end ListenAndServe")
	}()
	//fmt.Println("running dysvm")
	port := strconv.Itoa(listener.Addr().(*net.TCPAddr).Port)
	return port, srv, nil
}

func (k Keeper) EvalScript(goCtx context.Context, scriptCtx *EvalScriptContext, raiseRunErr bool) (resp *EvalScriptResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	cachedCtx, Write := ctx.CacheContext()
	defer func() {
		if r := recover(); r != nil {
			err = handleRunRecovery(r, ctx)
		}

		if err == nil {
			Write()
		} else {
			if raiseRunErr == true {
				resp = nil
			} else {
				resp = &EvalScriptResponse{Response: string(err.Error())}
				err = nil
			}
		}

		k.Logger(ctx).Info(fmt.Sprintf("last defer, resp: %+v  err: %+v", resp, err))
	}()

	resp, err = k.evalScript(cachedCtx, scriptCtx, raiseRunErr)
	ctx.GasMeter().ConsumeGas(cachedCtx.GasMeter().GasConsumed(), "EvalScript gas")
	return
}

func (k Keeper) evalScript(goCtx context.Context, scriptCtx *EvalScriptContext, raiseRunErr bool) (*EvalScriptResponse, error) {
	k.currentDepth += 1

	if k.currentDepth > 3 {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "Maximum Run recusion depth")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	//resolvedIndex := k.nameskeeper.ResolveIndex(ctx, scriptCtx.Index)
	//if resolvedIndex == "" {
	//	return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("Script at address %v not set or expired", scriptCtx.Index))
	//}
	resolvedAddress := scriptCtx.ScriptAddress

	// Run attached messages
	if err := k.ensureMsgAuthZ(scriptCtx.Messages, scriptCtx.CallerAddress); err != nil {
		return nil, err
	}
	// print the messages
	jsonMsgs := "["
	jsonMsgResults := "["

	for i, msg := range scriptCtx.Messages {

		if i > 0 {
			jsonMsgs += ","
			jsonMsgResults += ","
		}
		msgbz, err := k.cdc.MarshalInterfaceJSON(msg)
		if err != nil {
			fmt.Println("error:", err)
			return nil, errorsmod.Wrapf(err, "failed to marshal attached message at index %d: %s", i, msg)
		}

		jsonMsgs += string(msgbz)

		handler := k.Router.Handler(msg)
		if handler == nil {
			return nil, errorsmod.Wrapf(errors.ErrInvalid, "no message handler found for attached message %q at index %d", sdk.MsgTypeURL(msg), i)
		}
		r, err := handler(ctx, msg)
		if err != nil {
			return nil, errorsmod.Wrapf(err, "attached message %s at position %d", sdk.MsgTypeURL(msg), i)
		}
		// Handler should always return non-nil sdk.Result.
		if r == nil {
			return nil, fmt.Errorf("got nil sdk.Result for message %q at position %d", msg, i)
		}

		jsonMsgResults += string(k.cdc.MustMarshalJSON(r))

		// emit the events from the dispatched actions
		events := r.Events
		sdkEvents := make([]sdk.Event, 0, len(events))
		for _, event := range events {
			e := event
			e.Attributes = append(e.Attributes, abci.EventAttribute{Key: "attached_message_index", Value: strconv.Itoa(i)})

			sdkEvents = append(sdkEvents, sdk.Event(e))
		}

		ctx.EventManager().EmitEvents(sdkEvents)

	}

	jsonMsgs += "]"
	// join results with commas
	jsonMsgResults += "]"

	valFound, isFound := k.GetScript(ctx, resolvedAddress)

	if !isFound {
		if scriptCtx.ScriptAddress == scriptCtx.CallerAddress {
			valFound, isFound = k.GetScript(ctx, resolvedAddress)
			if !isFound {
				// This shouldn't happen
				fmt.Println(fmt.Sprintf("Script at address really %v not set", resolvedAddress))
				return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Script at address really %v not set", resolvedAddress))
			}
		}
	}
	if scriptCtx.CallerAddress != valFound.Address {
		if scriptCtx.ExtraCode != "" {
			return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("ExtraLines forbidden: Script.Address %v != Sender %v", valFound.Address, scriptCtx.CallerAddress))
		}
	}

	fmt.Println(fmt.Sprintf("Running script code: %v", valFound.Code))
	fmt.Println(fmt.Sprintf("Extra code: %v", scriptCtx.ExtraCode))

	port, srv, err := k.NewRPCServer(goCtx, valFound.Address)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	defer func() {
		fmt.Println(fmt.Sprintf("Elapsed time %s", time.Since(now)))
		k.currentDepth -= 1
	}()

	blockInfoJson, err := json.Marshal(ctx.BlockHeader())
	if err != nil {
		return nil, err
	}

	pyenv_root := os.Getenv("PYENV_ROOT")
	pythonExe := filepath.Join(pyenv_root, "versions", "blit-python", "bin", "python")

	blitvmPath := viper.GetString(blittypes.FlagBlitVMPath)
	out, runErr := exec.Command(
		pythonExe, filepath.Join(blitvmPath, "blitvm_server.py"),
		port,
		scriptCtx.CallerAddress,
		valFound.Address,
		string(blockInfoJson),
		scriptCtx.FunctionName,
		scriptCtx.Kwargs,
		scriptCtx.ExtraCode,
		valFound.Code,
		jsonMsgs,
		jsonMsgResults,
	).CombinedOutput()

	temp := strings.Split(string(out), "\n")
	response := string(temp[len(temp)-1])

	fmt.Println("worker finished")
	fmt.Printf("Output: %s\n", out)
	fmt.Printf("runErr: %s\n", runErr)
	fmt.Printf("Response: %s\n", response)
	evalResult := EvalResult{}
	err1 := json.Unmarshal([]byte(response), &evalResult)
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	// ctx.GasMeter().ConsumeGas(evalResult.NodesCalled, "nodes_called")
	fmt.Printf("ctx.GasMeter().GasConsumed(): %v\n", ctx.GasMeter().GasConsumed())
	fmt.Printf("cumsize: %v\n", evalResult.CumSize)
	fmt.Printf("nodes called: %v\n", evalResult.NodesCalled)
	fmt.Printf("currentDepth: %v\n", k.currentDepth)
	if err := srv.Shutdown(context.Background()); err != nil {
		fmt.Printf("shutdown error")
		panic(err) // failure/timeout shutting down the server gracefully
	}
	fmt.Println("server stopped")

	ctx.EventManager().EmitEvent(
		sdk.NewEvent("blit.script.MsgRun",
			sdk.NewAttribute("caller_address", scriptCtx.CallerAddress),
			sdk.NewAttribute("script_address", valFound.Address),
			sdk.NewAttribute("function_name", scriptCtx.FunctionName),
			sdk.NewAttribute("kwargs", scriptCtx.Kwargs),
			sdk.NewAttribute("messages", jsonMsgs),
		),
	)
	if (runErr != nil) && (raiseRunErr == true) {
		fmt.Printf("Output: %s\n", out)
		return nil, errorsmod.Wrapf(types.ScriptError, "%s on col: %v line:%v \nOutput:\n%s", evalResult.Exception.Msg, evalResult.Exception.Lineno, evalResult.Exception.Col, response)
	}

	return &EvalScriptResponse{Response: response}, nil
}

func handleRunRecovery(r interface{}, sdkCtx sdk.Context) error {
	switch r := r.(type) {
	case storetypes.ErrorOutOfGas:
		return errorsmod.Wrapf(sdkerrors.ErrOutOfGas,
			"ErrorOutofGas script out of gas in location: %v; gasWanted: %d, gasUsed: %d",
			r.Descriptor, sdkCtx.GasMeter().Limit(), sdkCtx.GasMeter().GasConsumed(),
		)

	default:
		return sdkerrors.ErrPanic.Wrapf("script recovered: %v", r)
	}
}
