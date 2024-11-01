package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"reflect"
	"testing"
)

func TestParseCsv(t *testing.T) {
	f, err := os.Open("test.csv")
	if err != nil {
		log.Fatalln(err)
	}

	rdr := csv.NewReader(f)
	for {
		rec, err := rdr.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatalf("error reading csv: %v", err)
		}
		str := testCsv{}
		if err := str.ParseCSV(rec); err != nil {
			t.Fatalf("error parsing csv record: %v", err)
		}

		expectedStr := testCsv{
			TestStr:     "TestString",
			TestInt32:   1,
			TestInt64:   12345,
			TestFloat32: 1.05,
			TestFloat64: 123456789.2,
		}

		if !reflect.DeepEqual(str, expectedStr) {
			t.Fatalf("expected: %#v, got: %#v", expectedStr, str)
		}
	}
}

func TestParseCSVWithTag(t *testing.T) {
	f, err := os.Open("test.csv")
	if err != nil {
		log.Fatalln(err)
	}

	rdr := csv.NewReader(f)
	for {
		rec, err := rdr.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatalf("error reading csv: %v", err)
		}

		str := testCsvWithTags{Foobar: "foobar"}
		if err := str.ParseCSV(rec); err != nil {
			t.Fatalf("error parsing csv record: %v", err)
		}

		expectedStr := testCsvWithTags{
			TestStr:     "TestString",
			TestInt64:   12345,
			TestFloat32: 1.05,
			Foobar:      "foobar",
		}

		if !reflect.DeepEqual(str, expectedStr) {
			t.Fatalf("expected: %#v, got: %#v", str, expectedStr)
		}
	}
}

func TestParseCSVWithOmitemptyTag(t *testing.T) {
	f, err := os.Open("test_omitempty.csv")
	if err != nil {
		log.Fatalln(err)
	}

	rdr := csv.NewReader(f)
	for {
		rec, err := rdr.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatalf("error reading csv: %v", err)
		}

		str := testCsvWithOmitEmptyTags{}
		if err := str.ParseCSV(rec); err != nil {
			t.Fatalf("error parsing csv record: %v", err)
		}

		expectedStr := testCsvWithOmitEmptyTags{
			TestStr:     "TestString",
			TestInt64:   0,
			TestFloat32: 1.05,
		}

		if !reflect.DeepEqual(str, expectedStr) {
			t.Fatalf("expected: %#v, got: %#v", str, expectedStr)
		}
	}
}
