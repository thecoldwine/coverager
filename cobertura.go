package main

import (
	"encoding/xml"
	"io"
)

type coberturaTotal struct {
	coberturaCoverageEntity

	LinesCovered    int64 `xml:"lines-covered,attr"`
	LinesValid      int64 `xml:"lines-valid,attr"`
	BranchesCovered int64 `xml:"branches-covered,attr"`
	BranchesValid   int64 `xml:"branches-valid,attr"`
}

type coberturaCoverageEntity struct {
	Name       *string `xml:"name,attr"`
	LineRate   float64 `xml:"line-rate,attr"`
	BranchRate float64 `xml:"branch-rate,attr"`
	Complexity int64   `xml:"complexity,attr"`
}

// ideally it should be a stream, not the in-memory structure,
// but I am lazy
func ParseCoberturaReport(r io.Reader) *CoverageReport {
	decoder := xml.NewDecoder(r)

	report := &CoverageReport{
		Header: []string{"Class", "Method", "Coverage"},
		Rows:   make([]CoverageReportEntry, 0),
	}

	currentClass := ""
	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		switch se := t.(type) {
		case xml.StartElement:
			switch se.Name.Local {
			case "coverage":
				var cov coberturaTotal
				decoder.DecodeElement(&cov, &se)
				report.Total = struct{ Coverage float64 }{Coverage: cov.LineRate * 100}
			case "class":
				var class coberturaCoverageEntity
				decoder.DecodeElement(&class, &se)
				currentClass = *class.Name
			case "method":
				var method coberturaCoverageEntity
				decoder.DecodeElement(&method, &se)
				report.Rows = append(report.Rows, CoverageReportEntry{
					Filename:     currentClass,
					FunctionName: *method.Name,
					Coverage:     method.LineRate * 100,
				})
			}
		default:
		}
	}

	return report
}
