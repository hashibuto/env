package env

import (
	"fmt"
	"strconv"
	"time"
)

func ParseDuration(value string) (time.Duration, error) {
	modifier := value[len(value)-1]
	data := value[0 : len(value)-1]
	val, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		return 0, err
	}

	switch modifier {
	case 'h':
		return time.Hour * time.Duration(val), nil
	case 'm':
		return time.Minute * time.Duration(val), nil
	case 's':
		return time.Second * time.Duration(val), nil
	default:
		return 0, fmt.Errorf("duration must end in h, m, s")
	}
}
