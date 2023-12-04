package app

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"cosmossdk.io/depinject"
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	circuitkeeper "cosmossdk.io/x/circuit/keeper"
	evidencekeeper "cosmossdk.io/x/evidence/keeper"
	feegrantkeeper "cosmossdk.io/x/feegrant/keeper"
	upgradekeeper "cosmossdk.io/x/upgrade/keeper"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	nodeservice "github.com/cosmos/cosmos-sdk/client/grpc/node"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	testdata_pulsar "github.com/cosmos/cosmos-sdk/testutil/testdata/testpb"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	consensuskeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	groupkeeper "github.com/cosmos/cosmos-sdk/x/group/keeper"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	capabilitykeeper "github.com/cosmos/ibc-go/modules/capability/keeper"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/keeper"
	icahostkeeper "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host/keeper"
	ibcfeekeeper "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/keeper"
	ibctransferkeeper "github.com/cosmos/ibc-go/v8/modules/apps/transfer/keeper"
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"

	blitmodulekeeper "blit/x/blit/keeper"
	scriptkeeper "blit/x/script/keeper"
	storagemodulekeeper "blit/x/storage/keeper"

	// this line is used by starport scaffolding # stargate/app/moduleImport

	"blit/docs"
)

const (
	AccountAddressPrefix = "blit"
	Name                 = "blit"
)

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string
)

var (
	_ runtime.AppI            = (*App)(nil)
	_ servertypes.Application = (*App)(nil)
)

// App extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type App struct {
	*runtime.App
	legacyAmino       *codec.LegacyAmino
	appCodec          codec.Codec
	txConfig          client.TxConfig
	interfaceRegistry codectypes.InterfaceRegistry

	// keepers
	AccountKeeper         authkeeper.AccountKeeper
	BankKeeper            bankkeeper.Keeper
	StakingKeeper         *stakingkeeper.Keeper
	SlashingKeeper        slashingkeeper.Keeper
	MintKeeper            mintkeeper.Keeper
	DistrKeeper           distrkeeper.Keeper
	GovKeeper             *govkeeper.Keeper
	CrisisKeeper          *crisiskeeper.Keeper
	UpgradeKeeper         *upgradekeeper.Keeper
	ParamsKeeper          paramskeeper.Keeper
	AuthzKeeper           authzkeeper.Keeper
	EvidenceKeeper        evidencekeeper.Keeper
	FeeGrantKeeper        feegrantkeeper.Keeper
	GroupKeeper           groupkeeper.Keeper
	ConsensusParamsKeeper consensuskeeper.Keeper
	CircuitBreakerKeeper  circuitkeeper.Keeper

	// IBC
	IBCKeeper           *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	CapabilityKeeper    *capabilitykeeper.Keeper
	IBCFeeKeeper        ibcfeekeeper.Keeper
	ICAControllerKeeper icacontrollerkeeper.Keeper
	ICAHostKeeper       icahostkeeper.Keeper
	TransferKeeper      ibctransferkeeper.Keeper

	// Scoped IBC
	ScopedIBCKeeper           capabilitykeeper.ScopedKeeper
	ScopedIBCTransferKeeper   capabilitykeeper.ScopedKeeper
	ScopedICAControllerKeeper capabilitykeeper.ScopedKeeper
	ScopedICAHostKeeper       capabilitykeeper.ScopedKeeper

	// this line is used by starport scaffolding # stargate/app/keeperDeclaration
	BlitKeeper    blitmodulekeeper.Keeper
	ScriptKeeper  scriptkeeper.Keeper
	StorageKeeper storagemodulekeeper.Keeper

	// simulation manager
	sm *module.SimulationManager
}

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, "."+Name)
}

// getGovProposalHandlers return the chain proposal handlers.
func getGovProposalHandlers() []govclient.ProposalHandler {
	var govProposalHandlers []govclient.ProposalHandler
	// this line is used by starport scaffolding # stargate/app/govProposalHandlers

	govProposalHandlers = append(govProposalHandlers,
		paramsclient.ProposalHandler,
		// this line is used by starport scaffolding # stargate/app/govProposalHandler
	)

	return govProposalHandlers
}

// AppConfig returns the default app config.
func AppConfig() depinject.Config {
	return depinject.Configs(
		appConfig,
		// Loads the ao config from a YAML file.
		// appconfig.LoadYAML(AppConfigYAML),
		depinject.Supply(
			// supply custom module basics
			map[string]module.AppModuleBasic{
				genutiltypes.ModuleName: genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),
				govtypes.ModuleName:     gov.NewAppModuleBasic(getGovProposalHandlers()),
				// this line is used by starport scaffolding # stargate/appConfig/moduleBasic
			},
		),
	)
}

// New returns a reference to an initialized App.
func New(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) (*App, error) {
	var (
		app        = &App{}
		appBuilder *runtime.AppBuilder

		// merge the AppConfig and other configuration in one config
		appConfig = depinject.Configs(
			AppConfig(),
			depinject.Supply(
				// Supply the application options
				appOpts,
				// Supply with IBC keeper getter for the IBC modules with App Wiring.
				// The IBC Keeper cannot be passed because it has not been initiated yet.
				// Passing the getter, the app IBC Keeper will always be accessible.
				// This needs to be removed after IBC supports App Wiring.
				app.GetIBCKeeper,
				app.GetCapabilityScopedKeeper,
				// Supply the logger
				logger,

				// ADVANCED CONFIGURATION
				//
				// AUTH
				//
				// For providing a custom function required in auth to generate custom account types
				// add it below. By default the auth module uses simulation.RandomGenesisAccounts.
				//
				// authtypes.RandomGenesisAccountsFn(simulation.RandomGenesisAccounts),
				//
				// For providing a custom a base account type add it below.
				// By default the auth module uses authtypes.ProtoBaseAccount().
				//
				// func() sdk.AccountI { return authtypes.ProtoBaseAccount() },
				//
				// For providing a different address codec, add it below.
				// By default the auth module uses a Bech32 address codec,
				// with the prefix defined in the auth module configuration.
				//
				// func() address.Codec { return <- custom address codec type -> }

				//
				// STAKING
				//
				// For provinding a different validator and consensus address codec, add it below.
				// By default the staking module uses the bech32 prefix provided in the auth config,
				// and appends "valoper" and "valcons" for validator and consensus addresses respectively.
				// When providing a custom address codec in auth, custom address codecs must be provided here as well.
				//
				// func() runtime.ValidatorAddressCodec { return <- custom validator address codec type -> }
				// func() runtime.ConsensusAddressCodec { return <- custom consensus address codec type -> }

				//
				// MINT
				//

				// For providing a custom inflation function for x/mint add here your
				// custom function that implements the minttypes.InflationCalculationFn
				// interface.
			),
		)
	)

	if err := depinject.Inject(appConfig,
		&appBuilder,
		&app.appCodec,
		&app.legacyAmino,
		&app.txConfig,
		&app.interfaceRegistry,
		&app.AccountKeeper,
		&app.BankKeeper,
		&app.StakingKeeper,
		&app.SlashingKeeper,
		&app.MintKeeper,
		&app.DistrKeeper,
		&app.GovKeeper,
		&app.CrisisKeeper,
		&app.UpgradeKeeper,
		&app.ParamsKeeper,
		&app.AuthzKeeper,
		&app.EvidenceKeeper,
		&app.FeeGrantKeeper,
		&app.GroupKeeper,
		&app.ConsensusParamsKeeper,
		&app.CircuitBreakerKeeper,
		// this line is used by starport scaffolding # stargate/app/keeperDefinition
		&app.BlitKeeper,
		&app.ScriptKeeper,
		&app.StorageKeeper,
	); err != nil {
		panic(err)
	}

	// Below we could construct and set an application specific mempool and
	// ABCI 1.0 PrepareProposal and ProcessProposal handlers. These defaults are
	// already set in the SDK's BaseApp, this shows an example of how to override
	// them.
	//
	// Example:
	//
	// app.App = appBuilder.Build(...)
	// nonceMempool := mempool.NewSenderNonceMempool()
	// abciPropHandler := NewDefaultProposalHandler(nonceMempool, app.App.BaseApp)
	//
	// app.App.BaseApp.SetMempool(nonceMempool)
	// app.App.BaseApp.SetPrepareProposal(abciPropHandler.PrepareProposalHandler())
	// app.App.BaseApp.SetProcessProposal(abciPropHandler.ProcessProposalHandler())
	//
	// Alternatively, you can construct BaseApp options, append those to
	// baseAppOptions and pass them to the appBuilder.
	//
	// Example:
	//
	// prepareOpt = func(app *baseapp.BaseApp) {
	// 	abciPropHandler := baseapp.NewDefaultProposalHandler(nonceMempool, app)
	// 	app.SetPrepareProposal(abciPropHandler.PrepareProposalHandler())
	// }
	// baseAppOptions = append(baseAppOptions, prepareOpt)
	//
	// create and set vote extension handler
	// voteExtOp := func(bApp *baseapp.BaseApp) {
	// 	voteExtHandler := NewVoteExtensionHandler()
	// 	voteExtHandler.SetHandlers(bApp)
	// }

	app.App = appBuilder.Build(db, traceStore, baseAppOptions...)

	// Register legacy modules
	app.registerIBCModules()
	app.AppRegisterIBCModules()

	// register streaming services
	if err := app.RegisterStreamingServices(appOpts, app.kvStoreKeys()); err != nil {
		return nil, err
	}

	/****  Module Options ****/

	app.ModuleManager.RegisterInvariants(app.CrisisKeeper)

	// add test gRPC service for testing gRPC queries in isolation
	testdata_pulsar.RegisterQueryServer(app.GRPCQueryRouter(), testdata_pulsar.QueryImpl{})

	// create the simulation manager and define the order of the modules for deterministic simulations
	//
	// NOTE: this is not required apps that don't use the simulator for fuzz testing
	// transactions
	overrideModules := map[string]module.AppModuleSimulation{
		authtypes.ModuleName: auth.NewAppModule(app.appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts, app.GetSubspace(authtypes.ModuleName)),
	}
	app.sm = module.NewSimulationManagerFromAppModules(app.ModuleManager.Modules, overrideModules)

	app.sm.RegisterStoreDecoders()

	// RegisterUpgradeHandlers is used for registering any on-chain upgrades.
	// Make sure it's called after `app.ModuleManager` and `app.configurator` are set.
	app.RegisterUpgradeHandlers()

	// A custom InitChainer can be set if extra pre-init-genesis logic is required.
	// By default, when using app wiring enabled module, this is not required.
	// For instance, the upgrade module will set automatically the module version map in its init genesis thanks to app wiring.
	// However, when registering a module manually (i.e. that does not support app wiring), the module version map
	// must be set manually as follow. The upgrade module will de-duplicate the module version map.
	//
	// app.SetInitChainer(func(ctx sdk.Context, req *abci.RequestInitChain) (*abci.ResponseInitChain, error) {
	// 	app.UpgradeKeeper.SetModuleVersionMap(ctx, app.ModuleManager.GetVersionMap())
	// 	return app.App.InitChainer(ctx, req)
	// })

	if err := app.Load(loadLatest); err != nil {
		return nil, err
	}

	return app, nil
}

// LegacyAmino returns App's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) LegacyAmino() *codec.LegacyAmino {
	return app.legacyAmino
}

// AppCodec returns App's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) AppCodec() codec.Codec {
	return app.appCodec
}

// GetKey returns the KVStoreKey for the provided store key.
func (app *App) GetKey(storeKey string) *storetypes.KVStoreKey {
	kvStoreKey, ok := app.UnsafeFindStoreKey(storeKey).(*storetypes.KVStoreKey)
	if !ok {
		return nil
	}
	return kvStoreKey
}

// GetMemKey returns the MemoryStoreKey for the provided store key.
func (app *App) GetMemKey(storeKey string) *storetypes.MemoryStoreKey {
	key, ok := app.UnsafeFindStoreKey(storeKey).(*storetypes.MemoryStoreKey)
	if !ok {
		return nil
	}

	return key
}

// kvStoreKeys returns all the kv store keys registered inside App.
func (app *App) kvStoreKeys() map[string]*storetypes.KVStoreKey {
	keys := make(map[string]*storetypes.KVStoreKey)
	for _, k := range app.GetStoreKeys() {
		if kv, ok := k.(*storetypes.KVStoreKey); ok {
			keys[kv.Name()] = kv
		}
	}

	return keys
}

// GetSubspace returns a param subspace for a given module name.
func (app *App) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// SimulationManager implements the SimulationApp interface.
func (app *App) SimulationManager() *module.SimulationManager {
	return app.sm
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *App) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	nodeservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	app.App.RegisterAPIRoutes(apiSvr, apiConfig)
	// register swagger API in app.go so that other applications can override easily
	if err := server.RegisterSwaggerAPI(apiSvr.ClientCtx, apiSvr.Router, apiConfig.Swagger); err != nil {
		panic(err)
	}

	// register app's OpenAPI routes.
	docs.RegisterOpenAPIService(Name, apiSvr.Router)
	apiSvr.Router.Use(catchAllMiddleware)
}

func catchAllMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a ResponseRecorder to capture the response
		rec := httptest.NewRecorder()

		// Get the host name from the request
		host := r.Host

		// Split the host name into subdomains
		hostParts := strings.Split(host, ".")

		var address string
		// Loop through the subdomains in reverse order to find a valid address
		for i := len(hostParts) - 1; i >= 0; i-- {
			hostItem := hostParts[i]
			_, err := sdk.AccAddressFromBech32(hostItem)
			if err == nil {
				address = hostItem
				break
			}
		}

		//if address == "" {
		//      address, err = DigBlitAddress(host)
		//      if err != nil {
		//              fmt.Println(w, "Failed to dig blit address", err)
		//      }
		//}

		// Let the handler process the request

		if address == "" && r.URL.Path == "/" {
			docs.DocHandler("Blit HTTP API Console")(rec, r)
		} else {
			next.ServeHTTP(rec, r)
		}
		notImplemented := rec.Body.String() == "{\"code\":12,\"message\":\"Not Implemented\",\"details\":[]}"
		if notImplemented {
			newReq := &http.Request{}
			newReq = newReq.WithContext(r.Context())
			// Change the URL of the request using the parsed address
			newReq.URL, _ = url.Parse("//localhost/blit/script/web/" + address)
			newReq.Method = "GET"

			rawRequest, err := getRawRequest(r)
			if err != nil {
				http.Error(w, "Failed to get raw request", http.StatusInternalServerError)
				return
			}

			//environ := ConvertRequestToWSGIEnviron(r)
			//convert to json string
			//jsonEnviron, err := json.Marshal(environ)
			//if err != nil {
			//      http.Error(w, "Failed to json marshal environ", http.StatusInternalServerError)
			//      return
			//}

			values := newReq.URL.Query()
			values.Add("httprequest", string(rawRequest))

			newReq.URL.RawQuery = values.Encode()

			// Use a new ResponseRecorder
			rec = httptest.NewRecorder()

			// Call the handler again with the new request
			next.ServeHTTP(rec, newReq)

			// Parse the JSON response from the fallback URL
			var response struct {
				HTTPResponse string `json:"httpresponse"`
				Message      string `json:"message"`
			}

			if response.Message != "" {
				http.Error(w, response.Message, http.StatusInternalServerError)
				return
			}

			json.NewDecoder(rec.Body).Decode(&response)

			// Decode the base64 encoded response
			decodedResponse, err := base64.StdEncoding.DecodeString(response.HTTPResponse)
			if err != nil {
				http.Error(w, "Failed to decode base64 response", http.StatusInternalServerError)
				return
			}

			// Write the decoded response
			err = WriteRawResponse(decodedResponse, w)
			if err != nil {
				http.Error(w, "Failed to write raw response", http.StatusInternalServerError)
				fmt.Println("Error:", err)
				fmt.Println("Error:", response.Message)
			}
			return
		}

		// Copy the original headers
		for k, v := range rec.Header() {
			w.Header()[k] = v
		}

		// Write the original status code and body
		w.WriteHeader(rec.Code)
		w.Write(rec.Body.Bytes())
	})
}

func WriteRawResponse(rawResponse []byte, w http.ResponseWriter) error {
	// Convert bytes to a buffered reader for easier line-by-line reading
	reader := bufio.NewReader(bytes.NewReader(rawResponse))

	// Read the status line
	statusLine, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read status line: %v, %s", err, rawResponse)
	}
	statusLine = strings.TrimSpace(statusLine) // Remove any trailing whitespace

	// Parse the status line
	parts := strings.SplitN(statusLine, " ", 3)
	if len(parts) < 2 {
		return fmt.Errorf("malformed status line: '%s'", statusLine)
	}
	statusCode, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("invalid status code: %v", err)
	}

	// Set the status code
	w.WriteHeader(statusCode)

	// Read and set headers
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read header line: %v", err)
		}
		line = strings.TrimSpace(line)
		if line == "" {
			break // Headers section has ended
		}
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) != 2 {
			return fmt.Errorf("malformed header: '%s'", line)
		}
		w.Header().Add(parts[0], parts[1])
	}

	// Write the body
	for {
		buffer := make([]byte, 1024)
		n, err := reader.Read(buffer)
		if err != nil && err.Error() != "EOF" {
			return fmt.Errorf("failed to read body: %v", err)
		}
		if n == 0 {
			break
		}
		_, err = w.Write(buffer[:n])
		if err != nil {
			return fmt.Errorf("failed to write body: %v", err)
		}
	}

	return nil
}

func getRawRequest(r *http.Request) (string, error) {
	var buf bytes.Buffer

	// Write the request method, URL, and protocol
	if _, err := fmt.Fprintf(&buf, "%s %s %s\r\n", r.Method, r.URL.RequestURI(), r.Proto); err != nil {
		return "", err
	}

	// Write the headers
	for k, vs := range r.Header {
		for _, v := range vs {
			if _, err := fmt.Fprintf(&buf, "%s: %s\r\n", k, v); err != nil {
				return "", err
			}
		}
	}

	// Write an extra CRLF to indicate the end of headers
	if _, err := fmt.Fprint(&buf, "\r\n"); err != nil {
		return "", err
	}

	// If there's a body, write it to the buffer
	if r.Body != nil {
		bodyBytes := new(bytes.Buffer)
		if _, err := bodyBytes.ReadFrom(r.Body); err != nil {
			return "", err
		}
		if _, err := buf.Write(bodyBytes.Bytes()); err != nil {
			return "", err
		}
		// IMPORTANT: Restore the body to the request object
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes.Bytes()))
	}

	return buf.String(), nil
}

// GetIBCKeeper returns the IBC keeper.
func (app *App) GetIBCKeeper() *ibckeeper.Keeper {
	return app.IBCKeeper
}

// GetCapabilityScopedKeeper returns the capability scoped keeper.
func (app *App) GetCapabilityScopedKeeper(moduleName string) capabilitykeeper.ScopedKeeper {
	return app.CapabilityKeeper.ScopeToModule(moduleName)
}

// GetMaccPerms returns a copy of the module account permissions
//
// NOTE: This is solely to be used for testing purposes.
func GetMaccPerms() map[string][]string {
	dup := make(map[string][]string)
	for _, perms := range moduleAccPerms {
		dup[perms.Account] = perms.Permissions
	}
	return dup
}

// BlockedAddresses returns all the app's blocked account addresses.
func BlockedAddresses() map[string]bool {
	result := make(map[string]bool)
	if len(blockAccAddrs) > 0 {
		for _, addr := range blockAccAddrs {
			result[addr] = true
		}
	} else {
		for addr := range GetMaccPerms() {
			result[addr] = true
		}
	}
	return result
}
