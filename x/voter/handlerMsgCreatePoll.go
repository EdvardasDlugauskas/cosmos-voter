package voter

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/EdvardasDlugauskas/voter/x/voter/types"
	"github.com/EdvardasDlugauskas/voter/x/voter/keeper"
)

func handleMsgCreatePoll(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreatePoll) (*sdk.Result, error) {
	k.CreatePoll(ctx, msg)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
