package types

import fmt "fmt"

const (
	defaultScriptDebugMode   = false
	defaultPublicRestAPIURL  = "http://localhost:1317"
	defaultPublicCometRPCURL = "http://localhost:26657"

	FlagBlitScriptDebugMode   = "blit.script_debug_mode"
	FlagPublicBlitRestAPIURL  = "blit.public_rest_api_url"
	FlagPublicBlitCometRPCURL = "blit.public_comet_rpc_url"
	FlagBlitVMPath            = "blit.experimental_blitvm_path"
)

// BlitConfig is the extra config required for wasm
type BlitConfig struct {
	ScriptDebugMode   bool
	PublicRestAPIURL  string
	PublicCometRPCURL string
	BlitVMPath        string
}

// DefaultBlitConfig returns the default settings for BlitConfig
func DefaultBlitConfig() BlitConfig {
	return BlitConfig{
		ScriptDebugMode:   defaultScriptDebugMode,
		PublicRestAPIURL:  defaultPublicRestAPIURL,
		PublicCometRPCURL: defaultPublicCometRPCURL,
		BlitVMPath:        "./blitvm",
	}
}

// DefaultConfigTemplate toml snippet with default values for app.toml
func DefaultConfigTemplate() string {
	return ConfigTemplate(DefaultBlitConfig())
}

// ConfigTemplate toml snippet for app.toml
func ConfigTemplate(c BlitConfig) string {
	return fmt.Sprintf(`
###############################################################################
###                         Blit                                            ###
###############################################################################

[blit]
# Enable debug mode for smart contracts
script_debug_mode = %v
# Public REST API URL for blitjs
public_rest_api_url = "%s"
# Public Comet RPC URL for blitjs
public_comet_rpc_url = "%s"
# EXPERIMENTAL: Path to blitvm executable
experimental_blitvm_path = "%s"

`, c.ScriptDebugMode, c.PublicRestAPIURL, c.PublicCometRPCURL, c.BlitVMPath)
}
