package keeper

import (
	"testing"

	"github.com/EZStaking/baobab/v13/x/oracle/types"
	"github.com/stretchr/testify/require"
)

func TestAddTrackingPriceHistoryProposal(t *testing.T) {
	ctx, keepers := CreateTestInput(t, false)
	govKeeper, oracleKeeper := keepers.GovKeeper, keepers.OracleKeeper

	var TwapTrackingList types.DenomList
	params := types.DefaultParams()
	params.TwapTrackingList = TwapTrackingList
	oracleKeeper.SetParams(ctx, params)

	params = oracleKeeper.GetParams(ctx)
	require.Equal(t, 0, len(params.TwapTrackingList))

	trackingList := types.DenomList{
		{
			BaseDenom:   types.Baobabdenom,
			SymbolDenom: types.JunoSymbol,
			Exponent:    types.JunoExponent,
		}, // Already in Whitelist (Default params)
	}

	src := types.AddTrackingPriceHistoryProposalFixture(func(p *types.AddTrackingPriceHistoryProposal) {
		p.TrackingList = trackingList
	})

	// submit proposal
	submitedProposal, err := govKeeper.SubmitProposal(ctx, src)
	require.NoError(t, err)
	// execute proposal
	handler := govKeeper.Router().GetRoute(submitedProposal.ProposalRoute())
	err = handler(ctx, submitedProposal.GetContent())
	require.NoError(t, err)

	params = oracleKeeper.GetParams(ctx)
	require.Equal(t, 1, len(params.TwapTrackingList))
	require.Equal(t, params.TwapTrackingList, trackingList)
}

func TestAddTrackingPriceHistoryWithWhitelistProposal(t *testing.T) {
	ctx, keepers := CreateTestInput(t, false)
	govKeeper, oracleKeeper := keepers.GovKeeper, keepers.OracleKeeper

	var emptyDenomList types.DenomList
	params := types.DefaultParams()
	params.Whitelist = emptyDenomList
	params.TwapTrackingList = emptyDenomList
	oracleKeeper.SetParams(ctx, params)

	params = oracleKeeper.GetParams(ctx)
	require.Equal(t, 0, len(params.Whitelist))
	require.Equal(t, 0, len(params.TwapTrackingList))

	trackingList := types.DenomList{
		{
			BaseDenom:   types.Baobabdenom,
			SymbolDenom: types.JunoSymbol,
			Exponent:    types.JunoExponent,
		},
		{
			BaseDenom:   types.AtomDenom,
			SymbolDenom: types.AtomSymbol,
			Exponent:    types.AtomExponent,
		},
	}

	src := types.AddTrackingPriceHistoryWithWhitelistProposalFixture(func(p *types.AddTrackingPriceHistoryWithWhitelistProposal) {
		p.TrackingList = trackingList
	})

	submittedProposal, err := govKeeper.SubmitProposal(ctx, src)
	require.NoError(t, err)

	// execute proposal
	handler := govKeeper.Router().GetRoute(submittedProposal.ProposalRoute())
	err = handler(ctx, submittedProposal.GetContent())
	require.NoError(t, err)

	params = oracleKeeper.GetParams(ctx)
	require.Equal(t, params.Whitelist, trackingList)
	require.Equal(t, params.TwapTrackingList, trackingList)
}

func TestRemoveTrackingPriceHistoryProposal(t *testing.T) {
	ctx, keepers := CreateTestInput(t, false)
	govKeeper, oracleKeeper := keepers.GovKeeper, keepers.OracleKeeper

	params := oracleKeeper.GetParams(ctx)
	require.Equal(t, 2, len(params.Whitelist))
	require.Equal(t, 2, len(params.TwapTrackingList))

	trackingList := types.DenomList{
		{
			BaseDenom:   types.Baobabdenom,
			SymbolDenom: types.JunoSymbol,
			Exponent:    types.JunoExponent,
		},
		{
			BaseDenom:   types.AtomDenom,
			SymbolDenom: types.AtomSymbol,
			Exponent:    types.AtomExponent,
		},
	}

	src := types.RemoveTrackingPriceHistoryProposalFixture(func(p *types.RemoveTrackingPriceHistoryProposal) {
		p.RemoveTwapList = types.DenomList{trackingList[0]}
	})

	submittedProposal, err := govKeeper.SubmitProposal(ctx, src)
	require.NoError(t, err)
	handler := govKeeper.Router().GetRoute(submittedProposal.ProposalRoute())
	err = handler(ctx, submittedProposal.GetContent())
	require.NoError(t, err)

	params = oracleKeeper.GetParams(ctx)
	require.Equal(t, params.Whitelist, trackingList)
	require.Equal(t, params.TwapTrackingList, types.DenomList{trackingList[1]})
}
