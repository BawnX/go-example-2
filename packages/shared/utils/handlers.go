package utils

import "github.com/hashicorp/go-hclog"

type Handlers struct {
	Log        hclog.Logger
	Validation *Validation
}
