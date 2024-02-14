package main

import (
	"fmt"
	"io"
	"text/template"
)

type CoverageReportEntry struct {
	Filename     string
	FunctionName string
	Coverage     float64
}

type CoverageReport struct {
	Header []string
	Rows   []CoverageReportEntry
	Total  struct {
		Coverage float64
	}
}

const templ = `{{ range .Header }}| {{ . }} {{ end }}|
|----------|----------|----------|{{ range .Rows }}
| {{ .Filename }} | {{ .FunctionName }} | {{ fmtCoverage .Coverage }} |{{ end }}
| Total | | {{ fmtCoverage .Total.Coverage }} |
`

func RenderReport(report *CoverageReport, threshold float64, w io.Writer) error {
	customFunctions := map[string]any{
		"fmtCoverage": func(v float64) string {
			if v < threshold {
				return fmt.Sprintf("❌ %.1f%%", v)
			}

			return fmt.Sprintf("✅ %.1f%%", v)
		},
	}

	t, err := template.New("templ1").Funcs(customFunctions).Parse(templ)
	if err != nil {
		logger.Fatalln(err)
	}

	t.Execute(w, report)

	return nil
}
