package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"slices"
	"text/template"

	"github.com/go-music-theory/music-theory/chord"
	"golang.org/x/exp/maps"
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

func NewModel(name string, intervals [4]int) Model {
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

	return m
}

func main() {

	chords := []string{
		"C", "Cmin", "C7", "Cmaj7",
	}

	for _, requiredChord := range chords {
		c := chord.Of(requiredChord)
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

func tonesAsIntervals(tones map[chord.Interval]note.Class) [4]int {

	intervals := [4]int{}

	intervalsFromTones := maps.Keys(tones)
	slices.Sort(intervalsFromTones)

	for i, inter := range intervalsFromTones {
		intervals[i] = int(tones[inter]) - 1
	}

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
