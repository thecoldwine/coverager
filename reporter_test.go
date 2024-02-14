package main

import (
	"bytes"
	"testing"
)

const exp1 string = `| Filename | Function | Coverage |
|----------|----------|----------|
| file1.go | function1.go | ✅ 75.4% |
| file1.go | function2.go | ❌ 24.5% |
| Total | | ❌ 25.7% |
`

func TestReport(t *testing.T) {
	cases :=
		map[string]struct {
			Expected string
			Report   CoverageReport
		}{
			"simple report": {
				Expected: exp1,
				Report: CoverageReport{
					Header: []string{"Filename", "Function", "Coverage"},
					Rows: []CoverageReportEntry{
						{
							Filename:     "file1.go",
							FunctionName: "function1.go",
							Coverage:     75.4,
						},
						{
							Filename:     "file1.go",
							FunctionName: "function2.go",
							Coverage:     24.5,
						},
					},
					Total: struct{ Coverage float64 }{
						Coverage: 25.7,
					},
				},
			},
		}

	for n, c := range cases {
		c := c
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			var buf bytes.Buffer

			err := RenderReport(&c.Report, 70.0, &buf)
			if err != nil {
				t.Fatalf("failure during test %s", err)
			}

			if buf.String() != c.Expected {
				t.Fatalf("expected\n%s but got\n%s", c.Expected, buf.String())
			}
		})
	}
}
