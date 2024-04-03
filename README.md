# Generate Chord Snippets for you Behringer X32

this script will generate a snippet file for each of the chords that you request from it.

It uses FX1-4 (dual pitch shift) to make up to 6 note chords.

## To make it work
Feed a sign wave into Input1 on the desk, double patch to 2-6

Load snippet

The snippet should handle the FX unit settings, the inserts into the channels, and whether the channel is unmuted

# To generate files

create an output directory `output/`

```go
  go run . <list of chords>
  go run . C F G Am Bb7
```