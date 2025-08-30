package source

import (
	"fmt"
	"regexp"
	"strconv"
)

var (
	ErrParse = fmt.Errorf("no match")
)

var (
	DefaultParse = Parse
	DefaultRegex = Regex
)

// Regex matches the following pattern:
//
//	123_name.up.ext
//	123_name.down.ext
var Regex = regexp.MustCompile(`^([0-9]+)[_-]([^.]*)(\.(` + string(Down) + `|` + string(Up) + `))?\.sql$`)

// Parse returns Migration for matching Regex pattern.
func Parse(raw string) (*Migration, error) {
	m := Regex.FindStringSubmatch(raw)
	if len(m) > 0 {
		versionUint64, err := strconv.ParseUint(m[1], 10, 64)
		if err != nil {
			return nil, err
		}
		if m[4] == "" {
			m[4] = string(Up)
		}
		return &Migration{
			Version:    uint(versionUint64),
			Identifier: m[2],
			Direction:  Direction(m[4]),
			Raw:        raw,
		}, nil
	}
	return nil, ErrParse
}
