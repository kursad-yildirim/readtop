package readtop

import (
	"time"
)

type myTop struct {
	text     []string
	Freeze   map[bool]time.Duration
	HasFocus bool
}

var Top myTop
var normalRefreshDelay, slowRefreshDelay float64 = 0.2, 2
var help, version bool = false, false
var MaxLines int = 50
var InputFileName, tokenToLook, nextTopFirstLine string = "", "", ""
var appVersion string = "read-top v1.1"
