package vfmt

import (
	"fmt"
	"github.com/vua/vfmt/internal"
	"io"
	"os"
)

func RegisterStyle(name ,style string){
	internal.CustomMap[name] = func(s string) string {
		return Sprintf("@[%s::%s]", s,style)
	}
}

// Sprint is the same as fmt.
func Sprint(a ...interface{}) string {
	text := fmt.Sprint(a...)
	parsed := internal.Parse(text)
	return parsed
}

// Fprint is the same as fmt.
func Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	text := Sprint(a...)
	return fmt.Fprint(w, text)
}

// Print is the same as fmt.
func Print(a ...interface{}) (n int, err error) {
	return Fprint(os.Stdout, a...)
}

// Sprintln is the same as fmt.
func Sprintln(a ...interface{}) string {
	return Sprint(a...) + "\n"
}

// Fprintln is the same as fmt.
func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	text := Sprintln(a...)
	return fmt.Fprint(w, text)
}

// Println is the same as fmt.
func Println(a ...interface{}) (n int, err error) {
	return Fprintln(os.Stdout, a...)
}

// Sprintf is the same as fmt.
func Sprintf(format string, a ...interface{}) string {
	text := fmt.Sprintf(format, a...)
	parsed := internal.Parse(text)
	return parsed
}

// Fprintf is the same as fmt.
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	text := Sprintf(format, a...)
	return fmt.Fprint(w, text)
}

// Printf is the same as fmt.
func Printf(format string, a ...interface{}) (n int, err error) {
	return Fprintf(os.Stdout, format, a...)
}

// Fatalf is the same as fmt.
func Fatalf(format string, a ...interface{}) {
	Printf(format, a...)
	os.Exit(1)
}

// Fatal is the same as fmt.
func Fatal(a ...interface{}) {
	Print(a...)
	os.Exit(1)
}

