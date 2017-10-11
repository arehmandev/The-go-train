package house

import (
	"strings"
)

const testVersion = 1

// Song - returns the song
func Song() string {

	return expectedSong

}

// Verse - return a verse
func Verse(input int) string {
	verses := strings.Split(expectedSong, "\n\n")
	return verses[input-1]
}
