package nginxParser

import (
	"bufio"
	"os"
	"regexp"
	"time"
)

func GetTimes(path string) ([]time.Time, error) {
	fileReader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(fileReader)
	times := make([]time.Time, 0)

	timeRegex, err := regexp.Compile(`.*\[([^\]]*)\].*`)
	if err != nil {
		return nil, err
	}

	_ = timeRegex

	for scanner.Scan() {
		newTimeString := timeRegex.ReplaceAllString(scanner.Text(), `$1`)
		newTime, err := time.Parse("2/Jan/2006:15:04:05 -0700", newTimeString)
		if err != nil {
			continue
		}
		times = append(times, newTime)
	}
	return times, nil
}

func GetRequests(dur time.Duration, path string) (int, error) {

	times, err := GetTimes(path)
	if err != nil {
		return 0, err
	}
	startTime := time.Now().Add(-1 * dur)
	beg := 0
	end := len(times)
	cur := 0
	for beg < end {
		cur = (beg + end) / 2
		after := startTime.After(times[cur])
		if after {
			beg = cur + 1
		} else {
			end = cur
		}
	}
	return len(times) - cur, nil
}
