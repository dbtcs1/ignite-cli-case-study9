package keeper

import (
	"github.com/username/planet/x/blog/types"
)

var _ types.QueryServer = Keeper{}
