package keeper

import (
	"github.com/0xknstntn/fluxo/x/oracle/types"
)

var _ types.QueryServer = Keeper{}
