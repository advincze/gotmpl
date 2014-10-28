package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"text/template"

	"flag"
	"fmt"
)

var stdin io.Reader
var stdout io.Writer

func init() {
	stdin = os.Stdin
	stdout = os.Stdout
}

func main() {

	var tmplFile = flag.String("t", "", "the go template file to use")
	var dataFile = flag.String("d", "", "the json data file to use")

	flag.Usage = func() {
		fmt.Println("usage: ")
		fmt.Println("       gotmpl [template] [data]\n")
		fmt.Println("examples: ")
		fmt.Println(`       gotmpl 'hello {{.name}}!' '{"name":"bob"}'`)
		fmt.Println(`       gotmpl -t hello.tmpl '{"name":"bob"}'`)
		fmt.Println("       gotmpl 'hello {{.name}}!' -d data.json")
		fmt.Println("       gotmpl -t hello.tmpl -d data.json")
		fmt.Println("       curl http://time.jsontest.com/ -s | gotmpl 'the time is  {{.time }}.' ")

		flag.PrintDefaults()
	}
	flag.Parse()
	args := flag.Args()
	tmpl := readTemplate(*tmplFile, args)
	if *tmplFile == "" {
		args = args[1:]
	}

	data := readData(*dataFile, args)

	err := tmpl.Execute(stdout, data)
	if err != nil {
		panic(err)
	}

}

func readTemplate(tmplFile string, args []string) *template.Template {
	if tmplFile == "" {
		if len(args) == 0 {
			fmt.Println("error: expected template")
			flag.Usage()
			os.Exit(1)
		}

		return template.Must(template.New("").Parse(args[0]))
	}
	return template.Must(template.ParseFiles(tmplFile))
}

func readData(dataFile string, args []string) map[string]interface{} {
	var data map[string]interface{}

	if dataFile == "" {
		if len(args) > 0 {
			json.Unmarshal([]byte(args[0]), &data)
		} else {
			bytes, err := ioutil.ReadAll(stdin)
			if err != nil {
				panic(err)
			}
			json.Unmarshal(bytes, &data)
		}
	} else {
		f, err := os.Open(dataFile)
		if err != nil {
			panic(err)
		}

		bytes, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}

		json.Unmarshal(bytes, &data)
	}
	return data
}
