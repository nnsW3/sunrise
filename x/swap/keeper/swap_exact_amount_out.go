package keeper

import (
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	lptypes "github.com/sunriselayer/sunrise/x/liquiditypool/types"
	"github.com/sunriselayer/sunrise/x/swap/types"
)

func (k Keeper) calculateInterfaceFeeExactAmountOut(
	ctx sdk.Context,
	hasInterfaceFee bool,
	amountOutNet math.Int,
) (amountOutGross math.Int, interfaceFee math.Int) {
	if !hasInterfaceFee {
		return amountOutNet, math.ZeroInt()
	}

	params := k.GetParams(ctx)
	// totalAmountOut = amountOut + interfaceFee
	//                = amountOut / (1 - interfaceFeeRate)
	amountOutGross = math.LegacyNewDecFromInt(amountOutNet).Quo(math.LegacyOneDec().Sub(params.InterfaceFeeRate)).TruncateInt()
	interfaceFee = amountOutGross.Sub(amountOutNet)

	return amountOutNet, interfaceFee
}

func (k Keeper) CalculateResultExactAmountOut(
	ctx sdk.Context,
	hasInterfaceFee bool,
	route types.Route,
	amountOut math.Int,
) (result types.RouteResult, interfaceFee math.Int, err error) {
	var (
		amountOutGross math.Int
	)

	amountOutGross, interfaceFee = k.calculateInterfaceFeeExactAmountOut(ctx, hasInterfaceFee, amountOut)

	result, err = k.calculateResultRouteExactAmountOut(ctx, route, amountOutGross)
	if err != nil {
		return result, interfaceFee, err
	}

	return result, interfaceFee, nil
}

func (k Keeper) SwapExactAmountOut(
	ctx sdk.Context,
	sender sdk.AccAddress,
	interfaceProvider string,
	route types.Route,
	maxAmountIn math.Int,
	amountOut math.Int,
) (result types.RouteResult, interfaceFee math.Int, err error) {
	var (
		hasInterfaceFee = interfaceProvider != ""
		totalAmountOut  math.Int
	)

	result, interfaceFee, err = k.CalculateResultExactAmountOut(ctx, hasInterfaceFee, route, totalAmountOut)
	if err != nil {
		return result, interfaceFee, err
	}

	// TODO: swap following the result
	if err := k.swapRouteExactAmountOut(ctx, sender, result); err != nil {
		return result, interfaceFee, err
	}

	if hasInterfaceFee {
		// Validated in ValidateBasic
		addr := sdk.MustAccAddressFromBech32(interfaceProvider)

		// TODO: Deduct interface fee
		_ = addr
	}

	if result.TokenIn.Amount.GT(maxAmountIn) {
		return result, interfaceFee, fmt.Errorf("TODO")
	}

	return result, interfaceFee, nil
}

func generateResultExactAmountOut(denomIn, denomOut string, amountExact, amountResult math.Int) (tokenIn sdk.Coin, tokenOut sdk.Coin) {
	return sdk.NewCoin(denomIn, amountResult), sdk.NewCoin(denomOut, amountExact)
}

func (k Keeper) calculateResultRouteExactAmountOut(
	ctx sdk.Context,
	route types.Route,
	amountOut math.Int,
) (result types.RouteResult, err error) {
	_, result, err = route.InspectRoute(
		amountOut,
		func(denomIn string, denomOut string, pool types.RoutePool, amountExact math.Int) (math.Int, error) {
			return k.calculateResultRoutePoolExactAmountOut(ctx, pool.PoolId, denomIn, denomOut, amountExact)
		},
		generateResultExactAmountOut,
		true,
	)

	return result, err
}

func (k Keeper) calculateResultRoutePoolExactAmountOut(
	ctx sdk.Context,
	poolId uint64,
	denomIn string,
	denomOut string,
	amountOut math.Int,
) (amountIn math.Int, err error) {
	pool, found := k.liquidityPoolKeeper.GetPool(ctx, poolId)
	if !found {
		return math.Int{}, lptypes.ErrPoolNotFound
	}

	// No needs to validate the denom,
	// as liquiditypool side is responsible for ensuring the denom exists in the pool.
	tokenOut := sdk.NewCoin(denomOut, amountOut)

	amountIn, err = k.liquidityPoolKeeper.CalculateResultExactAmountOut(
		ctx,
		pool,
		tokenOut,
		denomIn,
	)
	if err != nil {
		return math.Int{}, err
	}

	return amountIn, nil
}

func (k Keeper) swapRouteExactAmountOut(
	ctx sdk.Context,
	sender sdk.AccAddress,
	result types.RouteResult,
) error {
	switch strategy := result.Strategy.(type) {
	case *types.RouteResult_Pool:
		amountIn, err := k.swapRoutePoolExactAmountOut(
			ctx,
			sender,
			strategy.Pool.PoolId,
			result.TokenIn.Denom,
			result.TokenOut.Denom,
			result.TokenOut.Amount,
		)
		if err != nil {
			return err
		}

		if !amountIn.Equal(result.TokenIn.Amount) {
			return fmt.Errorf("TODO")
		}

		return nil

	case *types.RouteResult_Series:
		for _, r := range strategy.Series.RouteResults {
			if err := k.swapRouteExactAmountOut(ctx, sender, r); err != nil {
				return err
			}
		}

		return nil

	case *types.RouteResult_Parallel:
		for _, r := range strategy.Parallel.RouteResults {
			if err := k.swapRouteExactAmountOut(ctx, sender, r); err != nil {
				return err
			}
		}

		return nil
	}

	return fmt.Errorf("TODO")
}

func (k Keeper) swapRoutePoolExactAmountOut(
	ctx sdk.Context,
	sender sdk.AccAddress,
	poolId uint64,
	denomIn string,
	denomOut string,
	amountOut math.Int,
) (amountIn math.Int, err error) {
	pool, found := k.liquidityPoolKeeper.GetPool(ctx, poolId)
	if !found {
		return math.Int{}, lptypes.ErrPoolNotFound
	}

	// No needs to validate the denom,
	// as liquiditypool side is responsible for ensuring the denom exists in the pool.
	tokenOut := sdk.NewCoin(denomOut, amountOut)

	amountIn, err = k.liquidityPoolKeeper.SwapExactAmountOut(
		ctx,
		sender,
		pool,
		tokenOut,
		denomIn,
	)
	if err != nil {
		return math.Int{}, err
	}

	return amountIn, nil
}
