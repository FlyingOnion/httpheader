package header

import (
	"strconv"
	"time"
)

func parseInt(s string) (interface{}, error) {
	return strconv.Atoi(s)
}

func parseInt8(s string) (interface{}, error) {
	i, err := strconv.ParseInt(s, 10, 8)
	return int8(i), err
}

func parseInt16(s string) (interface{}, error) {
	i, err := strconv.ParseInt(s, 10, 16)
	return int16(i), err
}

func parseInt32(s string) (interface{}, error) {
	i, err := strconv.ParseInt(s, 10, 32)
	return int32(i), err
}

func parseInt64(s string) (interface{}, error) {
	return strconv.ParseInt(s, 10, 64)
}

func parseUint(s string) (interface{}, error) {
	u, err := strconv.ParseUint(s, 10, 64)
	return uint(u), err
}

func parseUint8(s string) (interface{}, error) {
	u, err := strconv.ParseUint(s, 10, 8)
	return uint8(u), err
}

func parseUint16(s string) (interface{}, error) {
	u, err := strconv.ParseUint(s, 10, 16)
	return uint16(u), err
}

func parseUint32(s string) (interface{}, error) {
	u, err := strconv.ParseUint(s, 10, 32)
	return uint32(u), err
}

func parseUint64(s string) (interface{}, error) {
	return strconv.ParseUint(s, 10, 64)
}

func parseBool(s string) (interface{}, error) {
	return strconv.ParseBool(s)
}

func parseString(s string) (interface{}, error) {
	return s, nil
}

var timeFormat string = "2006-01-02 15:04:05"

func parseTime(s string) (interface{}, error) {
	return time.ParseInLocation(timeFormat, s, time.Local)
}

func parseDuration(s string) (interface{}, error) {
	return time.ParseDuration(s)
}

func parseStringSlice(s []string) (interface{}, error) {
	return s, nil
}

func parseIntSlice(ss []string) (interface{}, error) {
	t := make([]int, 0, len(ss))
	for _, s := range ss {
		if i, err := parseInt(s); err == nil {
			t = append(t, i.(int))
		}
	}
	return t, nil
}
