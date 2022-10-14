package heartbeat

import (
	"agent/functions/host"
	"agent/proto"
	"time"
)

func getAgentStat(now time.Time) {
	rec := &proto.Record{
		DataType:  1000,
		Timestamp: now.Unix(),
		Data: &proto.Payload{
			Fields: map[string]string{},
		},
	}
	rec.Data.Fields["kernel_version"] = host.KernelVersion
}
