// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"github.com/ryt-io/ryt-v2/ids"
	"github.com/ryt-io/ryt-v2/utils/logging"
	"github.com/ryt-io/ryt-v2/vms"
)

var (
	// ID this VM should be referenced by
	IDStr = "subnetevm"
	ID    = ids.ID{'s', 'u', 'b', 'n', 'e', 't', 'e', 'v', 'm'}

	_ vms.Factory = (*Factory)(nil)
)

type Factory struct{}

func (*Factory) New(logging.Logger) (interface{}, error) {
	return &VM{}, nil
}
