package example

import (
	"github.com/efc2/efc2-agent/collector"
	"github.com/efc2/efc2-agent/common/metric"
	"github.com/efc2/efc2-agent/common/plugin"
)

// NewExample XXX
func NewExample(conf plugin.InitConfig) plugin.Plugin {
	return &Example{}
}

// Example XXX
type Example struct {
}

// Check XXX
func (e *Example) Check(agg metric.Aggregator) error {
	return nil
}

func init() {
	collector.Add("example", NewExample)
}
