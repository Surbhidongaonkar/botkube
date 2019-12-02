package utils

import (
	"fmt"
	"strings"

	"github.com/google/go-cmp/cmp"
)

// MetadataDiffReporter is a simple custom reporter that only records differences
// detected in Object Spec during comparison.
type MetadataDiffReporter struct {
	path  cmp.Path
	diffs []string
}

// PushStep custom implements Reporter interface
func (r *MetadataDiffReporter) PushStep(ps cmp.PathStep) {
	r.path = append(r.path, ps)
}

// Report custom implements Reporter interface
func (r *MetadataDiffReporter) Report(rs cmp.Result) {
	if !rs.Equal() {
		vx, vy := r.path.Last().Values()
		path := fmt.Sprintf("%#v", r.path)
		if ok := strings.Contains(path, ".Metadata."); ok {
			r.diffs = append(r.diffs, fmt.Sprintf("%#v:\n\t-: %+v\n\t+: %+v\n", r.path, vx, vy))
		}
	}
}

// PopStep custom implements Reporter interface
func (r *MetadataDiffReporter) PopStep() {
	r.path = r.path[:len(r.path)-1]
}

// String custom implements Reporter interface
func (r *MetadataDiffReporter) String() string {
	return strings.Join(r.diffs, "\n")
}
