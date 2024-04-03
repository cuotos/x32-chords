package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"slices"
	"text/template"

	"golang.org/x/exp/maps"
	"gopkg.in/music-theory.v0/chord"
	"gopkg.in/music-theory.v0/note"

	_ "embed"
)

var (
	//go:embed snippet.snp.tmpl
	snippetTemplate string
)

type Model struct {
	Name   string
	Pitch1 int
	Pitch2 int
	Pitch3 int
	Pitch4 int
	Pitch5 int
	Pitch6 int

	Pitch1State bool
	Pitch2State bool
	Pitch3State bool
	Pitch4State bool
	Pitch5State bool
	Pitch6State bool
}

func NewModel(name string, intervals [6]int) Model {
	baseNote := -12

	m := Model{
		Name:        name,
		Pitch1:      baseNote + intervals[0],
		Pitch1State: true,
	}

	if intervals[1] > 0 {
		m.Pitch2 = baseNote + intervals[1]
		m.Pitch2State = true
	}

	if intervals[2] > 0 {
		m.Pitch3 = baseNote + intervals[2]
		m.Pitch3State = true
	}

	if intervals[3] > 0 {
		m.Pitch4 = baseNote + intervals[3]
		m.Pitch4State = true
	}

	if intervals[4] > 0 {
		m.Pitch5 = baseNote + intervals[4]
		m.Pitch5State = true
	}

	if intervals[5] > 0 {
		m.Pitch6 = baseNote + intervals[5]
		m.Pitch6State = true
	}

	return m
}

func main() {
	for _, requiredChord := range os.Args[1:] {
		c := chord.Of(requiredChord)
		fmt.Println(c)
		m := NewModel(requiredChord, tonesAsIntervals(c.Tones))

		fmt.Println(m)

		output, err := renderTemplate(m)
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(fmt.Sprintf("output/%s.snp", m.Name), output, 0700)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func tonesAsIntervals(tones map[chord.Interval]note.Class) [6]int {

	intervals := [6]int{}

	values := maps.Values(tones)
	slices.Sort(values)

	for i, v := range values {
		intervals[i] = int(v) - 1
	}
	return intervals
}

func renderTemplate(model Model) ([]byte, error) {
	tmpl := template.Must(template.New("snippet").Parse(snippetTemplate))

	buf := bytes.NewBuffer([]byte{})

	err := tmpl.Execute(buf, model)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
