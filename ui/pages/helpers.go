package pages

import (
	"fmt"
	"time"
)

var S = fmt.Sprint
var F = fmt.Sprintf

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}
