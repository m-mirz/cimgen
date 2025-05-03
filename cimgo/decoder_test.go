package cimgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"testing"
)

func TestDecode(t *testing.T) {

	var logLevel = new(slog.LevelVar)
	logLevel.Set(slog.LevelInfo)
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: logLevel})))

	entries, err := filepath.Glob("../cgmes_schema/CGMES_3.0.0/IEC61970-600-2_CGMES_3_0_0_RDFS2020_TP.rdf")
	//entries, err := filepath.Glob("../cgmes_schema/CGMES_3.0.0/*.rdf")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Read files:", entries)

	for _, entry := range entries {

		b, err := os.ReadFile(entry)
		if err != nil {
			slog.Any("error", err)
		}

		newMap, err := DecodeToMap(bytes.NewReader(b))
		if err != nil {
			slog.Any("error", err)
		}

		jsonb, err := json.MarshalIndent(newMap, "", "  ")
		if err != nil {
			slog.Any("error", err)
		}
		//fmt.Fprintln(os.Stdout, string(jsonb))

		err = os.WriteFile("go-test.log", jsonb, 0644)
		if err != nil {
			slog.Any("error", err)
		}
	}
}
