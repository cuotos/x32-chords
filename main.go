package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"slices"
	"text/template"

	"github.com/go-music-theory/music-theory/chord"
	"gopkg.in/music-theory.v0/note"
)

type Model struct {
	Name   string
	Pitch1 int
	Pitch2 int
	Pitch3 int
	Pitch4 int

	Pitch1State bool
	Pitch2State bool
	Pitch3State bool
	Pitch4State bool
}

func NewModel(name string, intervals []int) Model {
	baseNote := -12
	return Model{
		Name:   name,
		Pitch1: baseNote + intervals[0],
		Pitch2: baseNote + intervals[1],
		Pitch3: baseNote + intervals[2],

		Pitch1State: true,
		Pitch2State: true,
		Pitch3State: true,
	}
}

func main() {

	chords := []string{
		"C", "G", "Am", "F",
	}

	for _, requiredChord := range chords {
		c := chord.Of(requiredChord)
		m := NewModel(requiredChord, tonesAsIntervals(c.Tones))

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

func tonesAsIntervals(tones map[chord.Interval]note.Class) []int {
	intervals := []int{}
	for _, v := range tones {
		intervals = append(intervals, int(v)-1)
	}

	slices.Sort(intervals)

	return intervals
}

func renderTemplate(model Model) ([]byte, error) {
	tmpl := template.Must(template.ParseFiles("snippet.snp.tmpl"))

	buf := bytes.NewBuffer([]byte{})

	err := tmpl.Execute(buf, model)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
