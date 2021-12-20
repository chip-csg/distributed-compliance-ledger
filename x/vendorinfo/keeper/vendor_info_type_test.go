package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNVendorInfoType(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.VendorInfoType {
	items := make([]types.VendorInfoType, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)
        
		keeper.SetVendorInfoType(ctx, items[i])
	}
	return items
}

func TestVendorInfoTypeGet(t *testing.T) {
	keeper, ctx := keepertest.VendorinfoKeeper(t)
	items := createNVendorInfoType(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetVendorInfoType(ctx,
		    item.Index,
            
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestVendorInfoTypeRemove(t *testing.T) {
	keeper, ctx := keepertest.VendorinfoKeeper(t)
	items := createNVendorInfoType(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveVendorInfoType(ctx,
		    item.Index,
            
		)
		_, found := keeper.GetVendorInfoType(ctx,
		    item.Index,
            
		)
		require.False(t, found)
	}
}

func TestVendorInfoTypeGetAll(t *testing.T) {
	keeper, ctx := keepertest.VendorinfoKeeper(t)
	items := createNVendorInfoType(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllVendorInfoType(ctx))
}
