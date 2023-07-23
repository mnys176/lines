// The lines tool takes text and breaks into multiple lines of a desired
// length. Additionally, one can also frame each line between an
// optional prefix and suffix.
//
// Usage:
//
//	lines [options] <string>
//
// Options:
//
//	--length, -l <length>
//	    Threshold at which to insert a new line. Default is 72
//	    characters.
//
//	--prefix, -p <prefix>
//	    Optional prefix prepended to the start of each line. The final
//	    line length will remain consistent with existing preferences.
//
//	--suffix, -s <suffix>
//	    Optional suffix appended to the end of each line. The final line
//	    length will remain consistent with existing preferences.
//
//	--help, -h
//	    Prints this help message and exits.
//
// The text can be of any length. For large amounts of text, it can be
// convenient to pass it in via stdin, e.g., piping in from a file.
//
// $ cat large-text.txt | lines
//
// The tool reads the input text from stdin before looking at command
// line arguments.
package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/mnys176/usage"
)

type linesError struct {
	err error
}

func (e linesError) Error() string {
	return "lines: " + e.err.Error()
}

var (
	l int
	p string
	s string
)

func init() {
	usage.Init("lines")
	usage.AddArg("<string>")

	lengthOption, _ := usage.NewOption([]string{"--length", "-l"}, "Threshold at which to insert a new line. Default is 72 characters.")
	lengthOption.AddArg("<length>")
	usage.AddOption(lengthOption)

	prefixOption, _ := usage.NewOption([]string{"--prefix", "-p"}, "Optional prefix prepended to the start of each line. The final line length will remain consistent with existing preferences.")
	prefixOption.AddArg("<prefix>")
	usage.AddOption(prefixOption)

	suffixOption, _ := usage.NewOption([]string{"--suffix", "-s"}, "Optional suffix appended to the end of each line. The final line length will remain consistent with existing preferences.")
	suffixOption.AddArg("<suffix>")
	usage.AddOption(suffixOption)

	helpOption, _ := usage.NewOption([]string{"--help", "-h"}, "Prints this help message and exits.")
	usage.AddOption(helpOption)

	flag.IntVar(&l, "length", 72, "")
	flag.IntVar(&l, "l", 72, "")
	flag.StringVar(&p, "prefix", "", "")
	flag.StringVar(&p, "p", "", "")
	flag.StringVar(&s, "suffix", "", "")
	flag.StringVar(&s, "s", "", "")
	flag.Usage = func() {
		fmt.Fprintln(os.Stdout, usage.Usage())
	}
}

func main() {
	flag.Parse()

	args, err := readArgs()
	if err != nil {
		panic(linesError{err})
	}

	if len(args) == 0 {
		err := linesError{errors.New("no string provided")}
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	if len(args) > 1 {
		err := linesError{errors.New("too many arguments")}
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	str := args[0]
	for _, line := range chopEssay(str, l-len(p)-len(s)) {
		fmt.Fprintln(os.Stdout, p+line+s)
	}
}

func chopLine(line string, length int) []string {
	line = strings.TrimSpace(line)
	splitter := regexp.MustCompile(`\s+`)
	words := splitter.Split(line, -1)
	lines := make([]string, 0)

	var b strings.Builder
	for _, w := range words {
		if len(w) > length {
			continue
		}
		if b.Len()+len(w) > length {
			lines = append(lines, strings.TrimSpace(b.String()))
			b.Reset()
		}
		b.WriteString(w + " ")
	}
	if b.Len() == 0 {
		return lines
	}
	return append(lines, strings.TrimSpace(b.String()))
}

func chopParagraph(paragraph string, length int) []string {
	paragraph = strings.TrimSpace(paragraph)
	splitter := regexp.MustCompile(`\n`)
	lines := make([]string, 0)
	for _, l := range splitter.Split(paragraph, -1) {
		if len(l) > 0 {
			lLines := chopLine(l, length)
			lines = append(lines, lLines...)
		}
	}
	return lines
}

func chopEssay(essay string, length int) []string {
	essay = strings.TrimSpace(essay)
	splitter := regexp.MustCompile("\n{2,}")
	lines := make([]string, 0)
	for _, p := range splitter.Split(essay, -1) {
		if len(p) > 0 {
			pLines := chopParagraph(p, length)
			if len(pLines) > 0 {
				pLines = append(pLines, "")
				lines = append(lines, pLines...)
			}
		}
	}
	if len(lines) == 0 {
		return lines
	}
	return lines[:len(lines)-1]
}

func readArgs() ([]string, error) {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}
	args := flag.Args()
	if stat.Size() > 0 {
		// Args should be read from stdin.
		bytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			return nil, err
		}
		args = append(args, string(bytes))
	}
	return args, nil
}
