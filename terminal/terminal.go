package terminal

import (
	"flag"
	"fmt"
	"mainmod/lib/common"
	"mainmod/lib/renderer"
	"os"
)

type Options struct {
	inputPath  string
	outputPath string
	format     string
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

	if _, ok := renderer.MapRenderer[options.format]; !ok {
		return nil, fmt.Errorf("unknown format '%s';\nallowed options are %s", options.format, mapToStr(renderer.MapRenderer))
	}

	if options.outputPath == "" && options.format == "" {
		options.format = renderer.ANSI
	} else if options.outputPath != "" && options.format == "" {
		options.format = renderer.HTML
	}

	return &options, nil
}

func (o *Options) RenderWith(parser common.Parser) (string, error) {
	rdr := renderer.MapRenderer[o.format]
	return rdr(parser)
}

func ExitIfErr(err error) {
	if err == nil {
		return
	}
	fmt.Printf("Error: %s\n", err)
	os.Exit(1)
}
