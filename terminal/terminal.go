package terminal

import (
	"flag"
	"fmt"
	"os"
)

type Options struct {
	inputPath  string
	outputPath string
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

func GetOptions() *Options {
	var options Options
	flag.StringVar(&options.outputPath, "out", "", "Output file")
	flag.Parse()
	options.inputPath = flag.Arg(0)

	return &options
}
