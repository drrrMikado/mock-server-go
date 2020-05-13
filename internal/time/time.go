package time

import (
	time2 "time"
)

type Duration time2.Duration

// UnmarshalText unmarshal text to duration.
func (d *Duration) UnmarshalText(text []byte) error {
	tmp, err := time2.ParseDuration(string(text))
	if err == nil {
		*d = Duration(tmp)
	}
	return err
}
