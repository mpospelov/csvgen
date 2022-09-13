// Code generated by qtc from "gen.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line cmd/generator/gen.qtpl:1
package main

//line cmd/generator/gen.qtpl:3
import "strings"

//line cmd/generator/gen.qtpl:5
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line cmd/generator/gen.qtpl:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line cmd/generator/gen.qtpl:6
type Context struct {
	Pkg     string
	Structs []Struct
	Uuid    string
}

//line cmd/generator/gen.qtpl:13
func streamtoLower(qw422016 *qt422016.Writer, str string) {
//line cmd/generator/gen.qtpl:13
	qw422016.N().S(`
`)
//line cmd/generator/gen.qtpl:15
	qw422016.N().S(strings.ToLower(str))
//line cmd/generator/gen.qtpl:16
	qw422016.N().S(`
`)
//line cmd/generator/gen.qtpl:17
}

//line cmd/generator/gen.qtpl:17
func writetoLower(qq422016 qtio422016.Writer, str string) {
//line cmd/generator/gen.qtpl:17
	qw422016 := qt422016.AcquireWriter(qq422016)
//line cmd/generator/gen.qtpl:17
	streamtoLower(qw422016, str)
//line cmd/generator/gen.qtpl:17
	qt422016.ReleaseWriter(qw422016)
//line cmd/generator/gen.qtpl:17
}

//line cmd/generator/gen.qtpl:17
func toLower(str string) string {
//line cmd/generator/gen.qtpl:17
	qb422016 := qt422016.AcquireByteBuffer()
//line cmd/generator/gen.qtpl:17
	writetoLower(qb422016, str)
//line cmd/generator/gen.qtpl:17
	qs422016 := string(qb422016.B)
//line cmd/generator/gen.qtpl:17
	qt422016.ReleaseByteBuffer(qb422016)
//line cmd/generator/gen.qtpl:17
	return qs422016
//line cmd/generator/gen.qtpl:17
}

//line cmd/generator/gen.qtpl:19
func StreamCsvTemplate(qw422016 *qt422016.Writer, ctx Context) {
//line cmd/generator/gen.qtpl:19
	qw422016.N().S(`
package `)
//line cmd/generator/gen.qtpl:20
	qw422016.N().S(ctx.Pkg)
//line cmd/generator/gen.qtpl:20
	qw422016.N().S(`

import (
    "fmt"
    "strconv"
)

`)
//line cmd/generator/gen.qtpl:27
	for _, str := range ctx.Structs {
//line cmd/generator/gen.qtpl:27
		qw422016.N().S(`
func (s *`)
//line cmd/generator/gen.qtpl:28
		qw422016.N().S(str.StructName)
//line cmd/generator/gen.qtpl:28
		qw422016.N().S(`) ParseCSV(rec []string) error {
    `)
//line cmd/generator/gen.qtpl:29
		for _, f := range str.Fields {
//line cmd/generator/gen.qtpl:29
			qw422016.N().S(`
        `)
//line cmd/generator/gen.qtpl:30
			switch f.Type {
//line cmd/generator/gen.qtpl:31
			case "string":
//line cmd/generator/gen.qtpl:31
				qw422016.N().S(`
        s.`)
//line cmd/generator/gen.qtpl:32
				qw422016.N().S(f.Name)
//line cmd/generator/gen.qtpl:32
				qw422016.N().S(` = rec[`)
//line cmd/generator/gen.qtpl:32
				qw422016.N().D(f.Position)
//line cmd/generator/gen.qtpl:32
				qw422016.N().S(`]
        `)
//line cmd/generator/gen.qtpl:33
			case "int32":
//line cmd/generator/gen.qtpl:33
				qw422016.N().S(`
        `)
//line cmd/generator/gen.qtpl:34
				qw422016.N().S(strings.ToLower(f.Name))
//line cmd/generator/gen.qtpl:34
				qw422016.N().S(`, err := strconv.ParseInt(rec[`)
//line cmd/generator/gen.qtpl:34
				qw422016.N().D(f.Position)
//line cmd/generator/gen.qtpl:34
				qw422016.N().S(`], 0, 32)
        if err != nil {
            return fmt.Errorf("error while parsing `)
//line cmd/generator/gen.qtpl:36
				qw422016.N().S(f.Name)
//line cmd/generator/gen.qtpl:36
				qw422016.N().S(` at index: %d", `)
//line cmd/generator/gen.qtpl:36
				qw422016.N().D(f.Position)
//line cmd/generator/gen.qtpl:36
				qw422016.N().S(`)
        }
        s.`)
//line cmd/generator/gen.qtpl:38
				qw422016.N().S(f.Name)
//line cmd/generator/gen.qtpl:38
				qw422016.N().S(` = int32(`)
//line cmd/generator/gen.qtpl:38
				qw422016.N().S(strings.ToLower(f.Name))
//line cmd/generator/gen.qtpl:38
				qw422016.N().S(`)
        `)
//line cmd/generator/gen.qtpl:39
			case "int64":
//line cmd/generator/gen.qtpl:39
				qw422016.N().S(`
        `)
//line cmd/generator/gen.qtpl:40
				qw422016.N().S(strings.ToLower(f.Name))
//line cmd/generator/gen.qtpl:40
				qw422016.N().S(`, err := strconv.ParseInt(rec[`)
//line cmd/generator/gen.qtpl:40
				qw422016.N().D(f.Position)
//line cmd/generator/gen.qtpl:40
				qw422016.N().S(`], 0, 64)
        if err != nil {
            return fmt.Errorf("error while parsing `)
//line cmd/generator/gen.qtpl:42
				qw422016.N().S(f.Name)
//line cmd/generator/gen.qtpl:42
				qw422016.N().S(` at index: %d", `)
//line cmd/generator/gen.qtpl:42
				qw422016.N().D(f.Position)
//line cmd/generator/gen.qtpl:42
				qw422016.N().S(`)
        }
        s.`)
//line cmd/generator/gen.qtpl:44
				qw422016.N().S(f.Name)
//line cmd/generator/gen.qtpl:44
				qw422016.N().S(` = `)
//line cmd/generator/gen.qtpl:44
				qw422016.N().S(strings.ToLower(f.Name))
//line cmd/generator/gen.qtpl:44
				qw422016.N().S(`
        `)
//line cmd/generator/gen.qtpl:45
			case "float32":
//line cmd/generator/gen.qtpl:45
				qw422016.N().S(`
        `)
//line cmd/generator/gen.qtpl:46
				qw422016.N().S(strings.ToLower(f.Name))
//line cmd/generator/gen.qtpl:46
				qw422016.N().S(`, err := strconv.ParseFloat(rec[`)
//line cmd/generator/gen.qtpl:46
				qw422016.N().D(f.Position)
//line cmd/generator/gen.qtpl:46
				qw422016.N().S(`], 32)
        if err != nil {
            return fmt.Errorf("error while parsing `)
//line cmd/generator/gen.qtpl:48
				qw422016.N().S(f.Name)
//line cmd/generator/gen.qtpl:48
				qw422016.N().S(` at index: %d", `)
//line cmd/generator/gen.qtpl:48
				qw422016.N().D(f.Position)
//line cmd/generator/gen.qtpl:48
				qw422016.N().S(`)
        }
        s.`)
//line cmd/generator/gen.qtpl:50
				qw422016.N().S(f.Name)
//line cmd/generator/gen.qtpl:50
				qw422016.N().S(` = float32(`)
//line cmd/generator/gen.qtpl:50
				qw422016.N().S(strings.ToLower(f.Name))
//line cmd/generator/gen.qtpl:50
				qw422016.N().S(`)
        `)
//line cmd/generator/gen.qtpl:51
			case "float64":
//line cmd/generator/gen.qtpl:51
				qw422016.N().S(`
        `)
//line cmd/generator/gen.qtpl:52
				qw422016.N().S(strings.ToLower(f.Name))
//line cmd/generator/gen.qtpl:52
				qw422016.N().S(`, err := strconv.ParseFloat(rec[`)
//line cmd/generator/gen.qtpl:52
				qw422016.N().D(f.Position)
//line cmd/generator/gen.qtpl:52
				qw422016.N().S(`], 64)
        if err != nil {
            return fmt.Errorf("error while parsing `)
//line cmd/generator/gen.qtpl:54
				qw422016.N().S(f.Name)
//line cmd/generator/gen.qtpl:54
				qw422016.N().S(` at index: %d", `)
//line cmd/generator/gen.qtpl:54
				qw422016.N().D(f.Position)
//line cmd/generator/gen.qtpl:54
				qw422016.N().S(`)
        }
        s.`)
//line cmd/generator/gen.qtpl:56
				qw422016.N().S(f.Name)
//line cmd/generator/gen.qtpl:56
				qw422016.N().S(` = `)
//line cmd/generator/gen.qtpl:56
				qw422016.N().S(strings.ToLower(f.Name))
//line cmd/generator/gen.qtpl:56
				qw422016.N().S(`
        `)
//line cmd/generator/gen.qtpl:57
			}
//line cmd/generator/gen.qtpl:57
			qw422016.N().S(`
    `)
//line cmd/generator/gen.qtpl:58
		}
//line cmd/generator/gen.qtpl:58
		qw422016.N().S(`
    return nil
}

`)
//line cmd/generator/gen.qtpl:62
	}
//line cmd/generator/gen.qtpl:62
	qw422016.N().S(`
`)
//line cmd/generator/gen.qtpl:63
}

//line cmd/generator/gen.qtpl:63
func WriteCsvTemplate(qq422016 qtio422016.Writer, ctx Context) {
//line cmd/generator/gen.qtpl:63
	qw422016 := qt422016.AcquireWriter(qq422016)
//line cmd/generator/gen.qtpl:63
	StreamCsvTemplate(qw422016, ctx)
//line cmd/generator/gen.qtpl:63
	qt422016.ReleaseWriter(qw422016)
//line cmd/generator/gen.qtpl:63
}

//line cmd/generator/gen.qtpl:63
func CsvTemplate(ctx Context) string {
//line cmd/generator/gen.qtpl:63
	qb422016 := qt422016.AcquireByteBuffer()
//line cmd/generator/gen.qtpl:63
	WriteCsvTemplate(qb422016, ctx)
//line cmd/generator/gen.qtpl:63
	qs422016 := string(qb422016.B)
//line cmd/generator/gen.qtpl:63
	qt422016.ReleaseByteBuffer(qb422016)
//line cmd/generator/gen.qtpl:63
	return qs422016
//line cmd/generator/gen.qtpl:63
}
