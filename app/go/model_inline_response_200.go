package swagger

import (
	"time"
)

type InlineResponse200 struct {

	Timestamp time.Time `json:"timestamp,omitempty"`

	Hostname string `json:"hostname,omitempty"`
}
