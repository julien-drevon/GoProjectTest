package core

import (
	"time"
)

type DateTimeNowProvider struct{}

func (provider *DateTimeNowProvider) GetDateTime() (time.Time, error) {
	return time.Now(), nil
}
