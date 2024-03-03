package terminal

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Options struct {
	inputPath  string
	outputPath string
	format     string
}

func (o *Options) GetContent() string {
	str, err := getContentFromInput(o.inputPath)
	if err != nil {
		fmt.Printf("Error: Can't open file '%s'", o.inputPath)
		os.Exit(1)
	}
	return str
}

func (o *Options) Output(content string) {
	err := output(o.outputPath, content)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}

var allowedFormats = map[string]bool{
	"ansi": true,
	"html": true,
}

func mapToStr[T any](m map[string]T) string {
	s := new(strings.Builder)
	for k := range m {
		s.WriteString(fmt.Sprintf("'%s'", k))
		s.WriteString(", ")
	}
	return s.String()
}

func GetOptions() (*Options, error) {
	var options Options
	flag.StringVar(&options.outputPath, "out", "", "Output file")
	flag.StringVar(&options.format, "format", "ansi", "Output file")
	flag.Parse()
	options.inputPath = flag.Arg(0)

	if _, ok := allowedFormats[options.format]; !ok {
		return nil, fmt.Errorf("unknown format '%s';\nallowed options are %s", options.format, mapToStr(allowedFormats))
	}

	return &options, nil
}

func ExitWithError(err error) {
	fmt.Printf("Error: %s\n", err)
	os.Exit(1)
}
