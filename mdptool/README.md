# Markdown preview tool

## Testing this tool

Try the tool by providing this Markdown file as input:

```shell
go​​ ​​run​​ ​​main.go​​ ​​-file​​ ​​README.md
```

Given the input the tool will create a `README.md.html`  file as output

The tool also supports providing a template file 

```shell
go run main.go -file README.md -t template.html.tmpl
```

In your browser you will get:
![Sample browser output](resources/sample-template-2023-12-03%2008-41-16.png)

## packages used 

[blackfriday v2](https://github.com/russross/blackfriday/tree/v2)

[bluemonday](https://github.com/microcosm-cc/bluemonday)


NB:  The `testdata` directory has a special meaning in Go tooling that’s ignored by the Go build tool when compiling your program to ensure that testing artifacts don’t end up in your build.

ref: [golden files](https://softwareengineering.stackexchange.com/questions/358786/what-are-golden-files)
