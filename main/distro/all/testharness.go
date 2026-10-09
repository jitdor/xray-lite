//go:build xraytest

package all

// Test-only: the integration tests in testing/scenarios drive the real binary
// with dokodemo-door as a local entry point. It is not part of the shipped build.
import (
	_ "github.com/xtls/xray-core/proxy/dokodemo"
)
