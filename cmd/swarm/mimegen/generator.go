// Copyright 2018 The Energi Core Authors
// Copyright 2018 The go-ethereum Authors
// This file is part of Energi Core.
//
// Energi Core is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Energi Core is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Energi Core. If not, see <http://www.gnu.org/licenses/>.

package main

// Standard "mime" package rely on system-settings, see mime.osInitMime
// Swarm will run on many OS/Platform/Docker and must behave similar
// This command generates code to add common mime types based on mime.types file
//
// mime.types file provided by mailcap, which follow https://www.iana.org/assignments/media-types/media-types.xhtml
//
// Get last version of mime.types file by:
// docker run --rm -v $(pwd):/tmp alpine:edge /bin/sh -c "apk add -U mailcap; mv /etc/mime.types /tmp"

import (
	"bufio"
	"bytes"
	"flag"
	"html/template"
	"io/ioutil"
	"strings"

	"log"
)

var (
	typesFlag   = flag.String("types", "", "Input mime.types file")
	packageFlag = flag.String("package", "", "Golang package in output file")
	outFlag     = flag.String("out", "", "Output file name for the generated mime types")
)

type mime struct {
	Name string
	Exts []string
}

type templateParams struct {
	PackageName string
	Mimes       []mime
}

func main() {
	// Parse and ensure all needed inputs are specified
	flag.Parse()
	if *typesFlag == "" {
		log.Fatalf("--types is required")
	}
	if *packageFlag == "" {
		log.Fatalf("--types is required")
	}
	if *outFlag == "" {
		log.Fatalf("--out is required")
	}

	params := templateParams{
		PackageName: *packageFlag,
	}

	types, err := ioutil.ReadFile(*typesFlag)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(types))
	for scanner.Scan() {
		txt := scanner.Text()
		if strings.HasPrefix(txt, "#") || len(txt) == 0 {
			continue
		}
		parts := strings.Fields(txt)
		if len(parts) == 1 {
			continue
		}
		params.Mimes = append(params.Mimes, mime{parts[0], parts[1:]})
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	result := bytes.NewBuffer([]byte{})

	if err := template.Must(template.New("_").Parse(tpl)).Execute(result, params); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(*outFlag, result.Bytes(), 0600); err != nil {
		log.Fatal(err)
	}
}

var tpl = `// Code generated by github.com/ethereum/go-ethereum/cmd/swarm/mimegen. DO NOT EDIT.

package {{ .PackageName }}

import "mime"
func init() {
	var mimeTypes = map[string]string{
{{- range .Mimes -}}
	{{ $name := .Name -}}
	{{- range .Exts }}
		".{{ . }}": "{{ $name | html }}",
	{{- end }}
{{- end }}
	}
	for ext, name := range mimeTypes {
		if err := mime.AddExtensionType(ext, name); err != nil {
			panic(err)
		}
	}
}
`