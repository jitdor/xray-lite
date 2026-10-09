package router

import (
	"context"

	"github.com/xtls/xray-core/common/dice"
)

// RandomStrategy represents a random balancing strategy
type RandomStrategy struct {
	FallbackTag string

	ctx context.Context
}

func (s *RandomStrategy) InjectContext(ctx context.Context) {
	s.ctx = ctx
}

func (s *RandomStrategy) GetPrincipleTarget(strings []string) []string {
	return strings
}

func (s *RandomStrategy) PickOutbound(candidates []string) string {
	count := len(candidates)
	if count == 0 {
		// goes to fallbackTag
		return ""
	}
	return candidates[dice.Roll(count)]
}
