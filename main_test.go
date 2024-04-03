package main

import (
	"testing"

	"github.com/go-music-theory/music-theory/chord"
	"github.com/stretchr/testify/assert"
	"gopkg.in/music-theory.v0/note"
	"gopkg.in/stretchr/testify.v1/require"
)

func TestCChord(t *testing.T) {
	expected := Model{
		Name:        "C",
		Pitch1:      -12,
		Pitch2:      -8,
		Pitch3:      -5,
		Pitch1State: true,
		Pitch2State: true,
		Pitch3State: true,
		Pitch4State: false,
	}

	c := chord.Of("C")

	actual := NewModel(c.Root.String(note.No), tonesAsIntervals(c.Tones))

	assert.Equal(t, expected, actual)
}

func TestDChord(t *testing.T) {
	expected := Model{
		Name:        "D",
		Pitch1:      -10,
		Pitch2:      -6,
		Pitch3:      -3,
		Pitch1State: true,
		Pitch2State: true,
		Pitch3State: true,
		Pitch4State: false,
	}

	c := chord.Of("D")

	actual := NewModel(c.Root.String(note.No), tonesAsIntervals(c.Tones))

	assert.Equal(t, expected, actual)
}

func TestAmChord(t *testing.T) {
	expected := Model{
		Name:        "Am",
		Pitch1:      -12,
		Pitch2:      -8,
		Pitch3:      -3,
		Pitch1State: true,
		Pitch2State: true,
		Pitch3State: true,
		Pitch4State: false,
	}

	c := chord.Of("Am")

	actual := NewModel("Am", tonesAsIntervals(c.Tones))

	assert.Equal(t, expected, actual)
}

func TestRenderTemplateC(t *testing.T) {
	expected := `#4.0# "T" 123008 -1 0 0 1                                                                                                      
/fx/1 PIT2
/fx/1/source INS INS
/fx/1/par -12 0 5.0 100 -100 100 -8 0 6.9 100 100 15k8 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
/fx/2 PIT2
/fx/2/source INS INS
/fx/2/par -5 0 5.0 100 -100 100 -1 0 6.9 100 100 15k8 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
/fx/3 PIT2
/fx/3/source INS INS
/fx/3/par 0 0 5.0 100 -100 100 0 0 6.9 100 100 15k8 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
/fx/4 PIT2
/fx/4/source INS INS
/fx/4/par 0 0 5.0 100 -100 100 0 0 6.9 100 100 15k8 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
/ch/01/mix/on ON
/ch/02/mix/on ON
/ch/03/mix/on OFF
/ch/04/mix/on OFF
/ch/05/mix/on OFF
/ch/06/mix/on OFF
/ch/07/mix/on OFF
/ch/08/mix/on OFF
`

	m := Model{
		Name:        "T",
		Pitch1:      -12,
		Pitch2:      -8,
		Pitch3:      -5,
		Pitch4:      -1,
		Pitch1State: true,
		Pitch2State: true,
		Pitch3State: false,
		Pitch4State: false,
	}

	actual, err := renderTemplate(m)
	require.NoError(t, err)

	assert.Equal(t, expected, string(actual))
}
