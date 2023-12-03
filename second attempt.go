validate-number(){
  re='^[0-9]+$'
  if [[ ! $1 =~ $re ]]; then
    echo "Error: The value must be a number." >&2
    exit 1
import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
)
