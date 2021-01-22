package voter

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/EdvardasDlugauskas/voter/x/voter/keeper"
	"github.com/EdvardasDlugauskas/voter/x/voter/types"
	"github.com/tendermint/tendermint/crypto"
)

func handleMsgCreatePoll(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreatePoll) (*sdk.Result, error) {
	var poll = types.Poll{
		Creator: msg.Creator,
		Title:   msg.Title,
		Options: msg.Options,
	}
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	payment, _ := sdk.ParseCoins("200token")

	k.CoinKeeper.DelegateCoins

	if err := k.CoinKeeper.SendCoins(ctx, poll.Creator, moduleAcct, payment); err != nil {
		return nil, err
	}
	k.CreatePoll(ctx, msg)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
