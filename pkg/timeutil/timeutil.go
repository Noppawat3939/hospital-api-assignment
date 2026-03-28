package timeutil

import "time"

const YYYYMMDD = "2006-01-02"

func ParseDate(v string) (*time.Time, error) {
	if v == "" {
		return nil, nil
	}

	t, err := time.Parse(YYYYMMDD, v)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
