package utils

import (
	"github.com/google/go-cmp/cmp"
	"github.com/infracloudio/botkube/pkg/config"
)

// Diff provides differences between two objects spec
func Diff(x, y interface{}, updateconfig config.UpdateConfig) string {
	msg := ""
	for _, val := range updateconfig.Fields {
		if val == config.SpecUpdate {
			var r SpecDiffReporter
			cmp.Equal(x, y, cmp.Reporter(&r))
			msg = msg + r.String()
		}
		if val == config.MetadataUpdate {
			var r MetadataDiffReporter
			cmp.Equal(x, y, cmp.Reporter(&r))
			msg = msg + r.String()
		}
		if val == config.StatusUpdate {
			var r StatusDiffReporter
			cmp.Equal(x, y, cmp.Reporter(&r))
			msg = msg + r.String()
		}
	}
	return msg
}
