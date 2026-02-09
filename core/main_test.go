// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package core

import (
	"os"
	"testing"

	"github.com/ryt-io/libevm/log"
	"go.uber.org/goleak"

	"github.com/ryt-io/subnet-evm/params"
	"github.com/ryt-io/subnet-evm/plugin/evm/customtypes"
)

// TestMain uses goleak to verify tests in this package do not leak unexpected
// goroutines.
func TestMain(m *testing.M) {
	RegisterExtras()

	customtypes.Register()
	params.RegisterExtras()

	// May of these tests are likely to fail due to `log.Crit` in goroutines.
	log.SetDefault(log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stderr, log.LevelCrit, true)))

	opts := []goleak.Option{
		// No good way to shut down these goroutines:
		goleak.IgnoreTopFunction("github.com/ryt-io/subnet-evm/core/state/snapshot.(*diskLayer).generate"),
		goleak.IgnoreTopFunction("github.com/ryt-io/libevm/core.(*txSenderCacher).cache"),
		goleak.IgnoreTopFunction("github.com/ryt-io/libevm/metrics.(*meterArbiter).tick"),
		goleak.IgnoreTopFunction("github.com/ryt-io/ryt-v2/vms/evm/metrics.(*meterArbiter).tick"),
		goleak.IgnoreTopFunction("github.com/syndtr/goleveldb/leveldb.(*DB).mpoolDrain"),
	}
	goleak.VerifyTestMain(m, opts...)
}
