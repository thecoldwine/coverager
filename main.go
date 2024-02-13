package main

import (
	"flag"
	"io"
	"log"
	"log/slog"
	"os"
	"strings"
)

var (
	inputFile  string
	inputType  string
	outputFile string
	threshold  float64

	logger *log.Logger
)

func main() {
	flag.StringVar(&inputFile, "input-file", "", "Path to the input file, default stdin.")
	flag.StringVar(&outputFile, "output-file", "", "Path to the output file, default stdout.")
	flag.StringVar(&inputType, "type", "", "Type of the input coverage report, mandatory")
	flag.Float64Var(&threshold, "threshold", 0.0, "Threshold for coverage - if lower, file will be highlighted")

	flag.Parse()

	logger = slog.NewLogLogger(slog.NewTextHandler(os.Stderr, nil), slog.LevelInfo)

	var r io.Reader
	if inputFile == "" {
		r = os.Stdin
	} else {
		f, err := os.Open(inputFile)
		if err != nil {
			logger.Fatalln(err)
		}

		defer f.Close()

		r = f
	}

	var w io.Writer
	if outputFile == "" {
		w = os.Stdout
	} else {
		f, err := os.Open(inputFile)
		if err != nil {
			logger.Fatalln(err)
		}

		defer f.Close()

		w = f
	}

	var report *CoverageReport

	switch strings.ToLower(inputType) {
	case "cobertura":
		panic("not implemented yet")
	case "gofunc":
		report = ParseGoFunc(r)
	case "":
		flag.Usage()
	default:
		logger.Fatalln("input type", inputType, "is not recognised")
	}

	err := RenderReport(report, threshold, w)
	if err != nil {
		logger.Fatalln(err)
	}
}
