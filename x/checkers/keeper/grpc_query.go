package keeper

import (
	"github.com/interchainmedia/checkers/x/checkers/types"
)

var _ types.QueryServer = Keeper{}
