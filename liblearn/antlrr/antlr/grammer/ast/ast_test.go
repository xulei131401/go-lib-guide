package ast

import (
	"bytes"
	"log"
	"testing"
)

func TestParse(t *testing.T) {
	parser := NewParser()
	expected := `Hello World
Hello gozero
Hello keson
`
	spec := parser.Parse(expected)
	writer := bytes.NewBuffer(nil)
	for _, e := range spec.Statements {
		writer.WriteString(e.HelloExpr.Text)
		writer.WriteString(" ")
		writer.WriteString(e.MsgExpr.Text)
		writer.WriteRune('\n')
	}
	actual := writer.String()
	log.Println(actual)

	if actual != expected {
		t.Fatalf("expected %v got %v\n", expected, actual)
	}
}
