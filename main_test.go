package main

import (
	"bytes"

	"os"
	"testing"
)

func TestReadTemplateFromArgs(t *testing.T) {
	tmpl := readTemplate("", []string{"Hello {{ .name }}!"})
	var buf bytes.Buffer
	tmpl.Execute(&buf, map[string]string{"name": "world"})

	exp := "Hello world!"
	if buf.String() != exp {
		t.Errorf("expected '%s' but got '%s'.", exp, buf.String())
	}
}

func TestReadTemplateFromFile(t *testing.T) {
	tmpl := readTemplate("hello.tmpl", nil)
	var buf bytes.Buffer
	tmpl.Execute(&buf, map[string]string{"name": "world"})

	exp := "Hello world!"
	if buf.String() != exp {
		t.Errorf("expected '%s' but got '%s'.", exp, buf.String())
	}
}

func TestReadDataFromFile(t *testing.T) {
	data := readData("adam.json", nil)
	exp := "Adam"
	if data["name"] != exp {
		t.Errorf("expected '%s' but got '%s'.", exp, data["name"])
	}
}

func TestReadDataFromArgs(t *testing.T) {
	data := readData("", []string{`{ "name": "Bob" }`})
	exp := "Bob"
	if data["name"] != exp {
		t.Errorf("expected '%s' but got '%s'.", exp, data["name"])
	}
}

func TestReadDataFromStdIn(t *testing.T) {
	stdin = bytes.NewBufferString(`{ "name": "Bob" }`)
	data := readData("", []string{})
	exp := "Bob"
	if data["name"] != exp {
		t.Errorf("expected '%s' but got '%s'.", exp, data["name"])
	}
}

func TestReadFromFiles(t *testing.T) {
	os.Args = []string{"gotmpl", "-t", "hello.tmpl", "-d", "adam.json"}
	var buf = bytes.NewBuffer(nil)
	stdout = buf
	main()
	exp := "Hello Adam!"
	if buf.String() != exp {
		t.Errorf("expected '%s' but got '%s'.", exp, buf.String())
	}
}
