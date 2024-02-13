package main

import (
	"bufio"
	"io"
	"strings"
)

func ParseGoFunc(r io.Reader) *CoverageReport {
	s := bufio.NewScanner(r)

	report := &CoverageReport{
		Rows: make([]CoverageReportEntry, 0),
	}

	for s.Scan() {
		fields := strings.Fields(s.Text())
		if len(fields) != 3 {
			logger.Fatalln("doesn't look like gofunc output, abort")
		}

		v, err := ParsePercentage(fields[2])
		if err != nil {
			logger.Fatalln("expected percentage, but got something else", err)
		}

		if fields[0] == "total:" {
			report.Total = struct{ Coverage float64 }{Coverage: v}
		} else {
			report.Rows = append(report.Rows, CoverageReportEntry{
				Filename:     fields[0],
				FunctionName: fields[1],
				Coverage:     v,
			})
		}
	}

	return report
}
