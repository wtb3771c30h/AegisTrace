// Package verifier statically admits safe, finite AegisTrace workflows.
package verifier

import (
	"fmt"
	"strings"
	"time"

	"github.com/aegistrace/aegistrace/workflow"
)

type Issue struct{ Path, Message string }

func (i Issue) Error() string { return i.Path + ": " + i.Message }

type Report struct {
	Workflow string  `json:"workflow"`
	Steps    int     `json:"steps"`
	Edges    int     `json:"edges"`
	Issues   []Issue `json:"-"`
}

