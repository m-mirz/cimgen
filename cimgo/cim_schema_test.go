package cimgo

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"os"
	"strings"
	"testing"
)

func TestSchemaImport(t *testing.T) {

	var logLevel = new(slog.LevelVar)
	logLevel.Set(slog.LevelInfo)
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: logLevel})))

	b, err := os.ReadFile("../cgmes_schema/CGMES_3.0.0/IEC61970-600-2_CGMES_3_0_0_RDFS2020_TP.rdf")
	if err != nil {
		slog.Any("error", err)
	}

	resultMap, err := DecodeToMap(bytes.NewReader(b))
	if err != nil {
		slog.Any("error", err)
	}

	f, err := os.Create("go-test.log")
	if err != nil {
		slog.Any("error", err)
	}
	defer f.Close()

	rdfMap := resultMap["rdf:RDF"].(map[string]interface{})
	descriptions := rdfMap["rdf:Description"].([]map[string]interface{})
	cimTypes := make([]CIMType, 0)
	cimAttributes := make([]CIMAttribute, 0)

	for _, v := range descriptions {
		objType := extractResource(v, "rdf:type")

		if strings.Contains(objType, "http://www.w3.org/2000/01/rdf-schema#Class") {
			classInfo := processClass(v)
			cimTypes = append(cimTypes, classInfo)

			jsonb, err := json.MarshalIndent(classInfo, "", "  ")
			if err != nil {
				slog.Any("error", err)
			}
			f.WriteString(string(jsonb) + "\n")
		}

		if strings.Contains(objType, "http://www.w3.org/1999/02/22-rdf-syntax-ns#Property") {
			propInfo := processProperty(v)
			cimAttributes = append(cimAttributes, propInfo)

			jsonb, err := json.MarshalIndent(propInfo, "", "  ")
			if err != nil {
				slog.Any("error", err)
			}
			f.WriteString(string(jsonb) + "\n")
		}
	}
}
