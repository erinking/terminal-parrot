package main

import "github.com/nsf/termbox-go"
import "strings"

var frame_index = 0

func reverse(lines []string) []string {
	newLines := make([]string, len(lines))
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		newLines[i], newLines[j] = lines[j], lines[i]
	}
	return newLines
}

// returns true when a cycle is finished
func draw(orientation string) bool {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	lines := strings.Split(frames[frame_index], "\n")

	if orientation == "aussie" {
		lines = reverse(lines)
	}

	for x, line := range lines {
		for y, cell := range line {
			color := colors[frame_index]
			if cell == '$' || cell == '+' {
				color = beak_color
			}
			termbox.SetCell(y, x, cell, color, termbox.ColorDefault)
		}
	}

	termbox.Flush()
	frame_index++
	if frame_index == num_frames {
		frame_index = 0
		return true
	}
	return false
}
