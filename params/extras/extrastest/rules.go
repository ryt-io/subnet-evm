// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package extrastest

import (
	"github.com/ryt-io/ryt-v2/upgrade"
	"github.com/ryt-io/ryt-v2/upgrade/upgradetest"

	"github.com/ryt-io/subnet-evm/params/extras"
)

func ForkToAvalancheRules(fork upgradetest.Fork) extras.AvalancheRules {
	networkUpgrades := extras.GetNetworkUpgrades(upgradetest.GetConfig(fork))
	return networkUpgrades.GetAvalancheRules(uint64(upgrade.InitiallyActiveTime.Unix()))
}
