package conf

import (
	"google.golang.org/protobuf/proto"
)

const (
	strategyRandom     string = "random"
	strategyRoundRobin string = "roundrobin"
)

var strategyConfigLoader = NewJSONConfigLoader(ConfigCreatorCache{
	strategyRandom:     func() interface{} { return new(strategyEmptyConfig) },
	strategyRoundRobin: func() interface{} { return new(strategyEmptyConfig) },
}, "type", "settings")

type strategyEmptyConfig struct{}

func (v *strategyEmptyConfig) Build() (proto.Message, error) {
	return nil, nil
}
