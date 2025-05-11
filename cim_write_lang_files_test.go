package cimgen

import (
	"bytes"
	"log/slog"
	"os"
	"path/filepath"
	"testing"
	"text/template"
)

func TestGenerate(t *testing.T) {

	var logLevel = new(slog.LevelVar)
	logLevel.Set(slog.LevelInfo)
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: logLevel})))

	entries, err := filepath.Glob("cgmes_schema/CGMES_3.0.0/IEC61970-600-2_CGMES_3_0_0_RDFS2020_*.rdf")
	if err != nil {
		panic(err)
	}

	allCimTypesMerged := make(map[string]*CIMType, 0)
	allCimEnums := make([]map[string]*CIMEnum, 0)

	for _, entry := range entries {

		b, err := os.ReadFile(entry)
		if err != nil {
			slog.Any("error", err)
		}

		resultMap, err := DecodeToMap(bytes.NewReader(b))
		if err != nil {
			slog.Any("error", err)
		}

		cimTypes, cimEnums, _ := processRDFMap(resultMap)
		allCimTypesMerged = mergeCimTypes(allCimTypesMerged, cimTypes)
		allCimEnums = append(allCimEnums, cimEnums)
	}

	var tmplFile = "cimpy_class_template.tmpl"
	tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}

	var f *os.File
	f, err = os.Create("cim_write_lang_files_test.py")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = tmpl.Execute(f, allCimTypesMerged["ACDCTerminal"])
	if err != nil {
		panic(err)
	}
}
