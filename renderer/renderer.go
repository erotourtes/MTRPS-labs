package renderer

func Render(input string, parser Parser) (string, error) {
	err := parser.Parse()
	if err != nil {
		return "", err
	}
	return input, nil
}
