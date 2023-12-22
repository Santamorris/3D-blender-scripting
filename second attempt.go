validate-number(){
  re='^[0-9]+$'
  if [[ ! $1 =~ $re ]]; then
    echo "Error: The value must be a number." >&2
    exit 1
import (
	"fmt"

//	sdk "github.com/cosmos/cosmos-sdk/types"
//	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
)

	func (suite *IntegrationTestSuite) TestAppleIAPOrderExist() {
	k := suite.k
	ctx := suite.ctx
	require := suite.Require()

	items := createNAppleIAPOrder(k, ctx, 10)
	for _, item := range items {
		require.True(k.HasAppleIAPOrder(ctx, item.PurchaseId))
			items := createNAppleIAPOrder(k, ctx, 10)
	require.Equal(items, k.GetAllAppleIAPOrder(ctx))
	require.Equal(items, k.GetAllAppleIAPOrder(ctx))
	}
}
