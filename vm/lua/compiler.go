package lua

import (
	"strings"
	"errors"
)

type DocCommentParser struct {
	text string // always ends with \0, which doesn't appear elsewhere

	index int
}

func NewDocCommentParser(text string) (*DocCommentParser, error) {
	parser := new(DocCommentParser)
	if strings.Contains(text, "\\0") {
		return nil, errors.New("Text contains character \\0, parse failed")
	}
	parser.text = text + "\\0"
	return parser, nil
}

