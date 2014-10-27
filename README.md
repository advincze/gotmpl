gotmpl
======

gotmpl is asmall utility to execute go templates on the command line via files, aguments or stdin
  
    go get github.com/advincze/gotmpl

usage:

    gotmpl [template] [data]

examples

    gotmpl 'hello {{.name}}!' '{"name":"bob"}'
    gotmpl -t hello.tmpl '{"name":"bob"}'
    gotmpl 'hello {{.name}}!' -d data.json
    gotmpl -t hello.tmpl -d data.json
    gotmpl -t hello.tmpl < data.json
    curl http://time.jsontest.com/ -s | gotmpl 'the time is  {{.time }}.'
    
