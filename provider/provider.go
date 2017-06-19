// Package provider isolates model and tool effects behind a narrow interface.
package provider

import (
	"context"
	"fmt"
	"sort"
	"sync"
)

type Request struct {
	RunID     string         `json:"run_id"`
