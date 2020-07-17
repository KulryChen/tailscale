// Copyright (c) 2020 Tailscale Inc & AUTHORS All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dns

func newManager(mconfig ManagerConfig) managerImpl {
	switch {
	case resolvedIsActive() && mconfig.PerDomain:
		if mconfig.Cleanup {
			return newNoopManager(mconfig)
		} else {
			return newResolvedManager(mconfig)
		}
	case nmIsActive():
		if mconfig.Cleanup {
			return newNoopManager(mconfig)
		} else {
			return newNMManager(mconfig)
		}
	case resolvconfIsActive():
		return newResolvconfManager(mconfig)
	default:
		return newDirectManager(mconfig)
	}
}