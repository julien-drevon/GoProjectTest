package core

import (
	"time"
)

type DateTimeNowProvider struct{}

func (provider *DateTimeNowProvider) GetDateTime() time.Time {
	return time.Now()
}
