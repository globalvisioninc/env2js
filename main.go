package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/joho/godotenv"
)

// EnvVar represents a environment variable (key/value pair)
type EnvVar struct {
	Name  string
	Value string
}

// Template for rendering the JavaScript config file
var tpl = `window._env_ = {
{{- $length := len . }}
{{- range $i, $var := . }}
  {{ $var.Name }}: "{{ $var.Value | js }}"{{if more $i $length }},{{end}}
{{- end}}
};
`

// Custom template function to check if, in a loop, there are other remaining items
var fns = template.FuncMap{
	"more": func(index int, length int) bool {
		return index < length-1
	},
}

// Renders the template above with the passed list of variables and outputs to the specified file
func renderDotEnv(vars []EnvVar, dstEnvJs string) error {
	envConfig, err := os.Create(dstEnvJs)
	if err != nil {
		return errors.New("Unable to create destination file: " + dstEnvJs)
	}
	defer envConfig.Close()

	t := template.Must(template.New("vars").Funcs(fns).Parse(tpl))
	err = t.Execute(envConfig, vars)
	if err != nil {
		return err
	}
	return nil
}

// Custom flag for list of environment files
type envFiles []string

func (i *envFiles) String() string {
	return strings.Join(*i, " ")
}

func (i *envFiles) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func main() {
	// Parse arguments
	var envPrefix = flag.String("prefix", "REACT_APP_", "Prefix of environment variables to include")
	var jsFile = flag.String("out", "", "Path to output the JS file with variables")
	var envFilesArr envFiles
	flag.Var(&envFilesArr, "env", "(Optional) Environnment file to parse. You can specify this flag multiple times.")

	flag.Parse()

	if *jsFile == "" {
		fmt.Println("'out' argument is required")
		os.Exit(1)
	}

	// Load environment files
	for _, file := range envFilesArr {
		// For now, we don't care if this returns an error, we load only if they are present in the CWD
		err := godotenv.Load(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// Loop through environment variables and keep only those that start with "REACT_APP_"
	var appVars = make([]EnvVar, 0)
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if !strings.HasPrefix(pair[0], *envPrefix) {
			continue
		}
		appVars = append(appVars, EnvVar{Name: pair[0], Value: pair[1]})
	}

	err := renderDotEnv(appVars, *jsFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
