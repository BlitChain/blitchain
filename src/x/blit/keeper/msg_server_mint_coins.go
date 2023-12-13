package keeper

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"blit/x/blit/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
)

const COIN_GAS_COST = 1

const reDnmString = `blit/([a-z0-9-_.]+)(/[a-zA-Z0-9-_+]+)*`

var reDnm = regexp.MustCompile(fmt.Sprintf(`^%s$`, reDnmString))

func ValidateDenomOwner(denom string) (string, error) {
	if !reDnm.MatchString(denom) {
		return "", fmt.Errorf("invalid denom (%s) must match: %s", denom, reDnmString)
	}
	// split the denom into its parts
	parts := strings.Split(denom, "/")
	return parts[1], nil
}

func (k msgServer) MintCoins(goCtx context.Context, msg *types.MsgMintCoins) (*types.MsgMintCoinsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, err := ValidateDenomOwner(msg.Amount.Denom)
	if err != nil {
		return nil, err
	}

	ownerAddress, err := sdk.AccAddressFromBech32(owner)
	if err != nil {
		return nil, err
	}

	if msg.Grantee != owner {

		// Prevent looping
		originalGrantee := sdk.MustAccAddressFromBech32(msg.Grantee)
		msg.Grantee = ownerAddress.String()
		execMsg := authz.NewMsgExec(
			originalGrantee,
			[]sdk.Msg{msg},
		)

		msgExecResp, err := k.authzKeeper.Exec(ctx, &execMsg)
		if err != nil {
			return nil, err
		}

		var execResp types.MsgMintCoinsResponse

		data := msgExecResp.Results[0]
		k.cdc.MustUnmarshal(data, &execResp)
		return &execResp, nil
	}

	gasToUse := msg.Amount.Amount.MulRaw(COIN_GAS_COST).Uint64()
	ctx.GasMeter().ConsumeGas(gasToUse, "MintCoins")

	err = k.bankKeeper.MintCoins(ctx, "transfer", sdk.NewCoins(msg.Amount))

	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, "transfer", ownerAddress, sdk.NewCoins(msg.Amount))

	if err != nil {
		return nil, err
	}
	return &types.MsgMintCoinsResponse{}, nil
}
