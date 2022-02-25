package timeRanges

import "time"

type TimeRangeT struct {
	S time.Time
	E time.Time
}

func Generate(rangeStart, rangeEnd time.Time, step time.Duration) []TimeRangeT {
	var ranges []TimeRangeT

	for partialStart := rangeStart; partialStart.Before(rangeEnd); partialStart = partialStart.Add(step) {
		partialEnd := partialStart.Add(step)
		if partialEnd.After(rangeEnd) {
			partialEnd = rangeEnd
		}
		ranges = append(ranges, TimeRangeT{S: partialStart, E: partialEnd})
	}

	return ranges
}
