package terminal

import (
	"flag"
	"fmt"
	"os"
)

type Options struct {
	inputPath  string
	outputPath string
	format     string
}

var allowedFormats = map[string]bool{
	"ansi": true,
	"html": true,
}

func (o *Options) GetContent() (string, error) {
	str, err := getContentFromInput(o.inputPath)
	if err != nil {
		return "", fmt.Errorf("can't open file '%s'", o.inputPath)
	}
	return str, nil
}

func (o *Options) Output(content string) error {
	return output(o.outputPath, content)
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

func ExitIfErr(err error) {
	if err == nil {
		return
	}
	fmt.Printf("Error: %s\n", err)
	os.Exit(1)
}
