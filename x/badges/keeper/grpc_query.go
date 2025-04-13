package keeper

import (
	"github.com/bitbadges/badges-module/x/badges/types"
)

var _ types.QueryServer = Keeper{}
