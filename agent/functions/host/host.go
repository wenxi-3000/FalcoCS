package host

import (
	"sync/atomic"

	"github.com/shirou/gopsutil/v3/host"
)

var (
	KernelVersion atomic.Value
)

func init() {
	KernelVersion, _ = host.KernelVersion()
}
