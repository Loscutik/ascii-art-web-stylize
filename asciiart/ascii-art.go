package asciiart

import (
	"fmt"
	"io"
	"strings"
)

/*
turns a string into an ascii graphic string
*/
func StringToArt(str string, afont ArtFont) (aRows ArtString) {
	for _, ch := range str {
		for i := 0; i < SYMBOL_HEIGHT; i++ {
			aRows[i] += afont[ch-FIRST_SYMBOL][i]
		}
	}
	return
}

/*
turns a whole text into one string, which represents ascii art view of the given text
*/
func TextToArt(text string, banner string) (aText string, err error) {
	artFont, err := GetArtFont(banner)
	if err != nil {
		err = fmt.Errorf("cannot get the art font: %s", err)
		return "", err
	}

	strs := strings.Split(strings.ReplaceAll(text, "\r", ""), "\n")
	for _, str := range strs {
		if str == "" {
			aText += "\n"
		} else {
			aStr := StringToArt(str, artFont)
			for i := 0; i < SYMBOL_HEIGHT; i++ {
				aText += aStr[i] + "\n"
			}
		}
	}
	return aText, nil
}

/*
prints an ascii graphic string 
*/
func (aStr ArtString) ArtFprint(w io.Writer) {
	// the empty string must comprise only 1 line
	if aStr[0] == "" {
		fmt.Fprintln(w)
		return
	}
	for _, line := range aStr {
		fmt.Fprint(w, line)
	}
}

/*
checks if string contains only printable ascii symbols
*/
func IsAsciiString(str string) (bool, []rune) {
	res := true

	var notValidRunes []rune
	str = strings.ReplaceAll(strings.ReplaceAll(str, "\r", ""), "\n", "")
	for _, rune := range str {
		if rune < FIRST_SYMBOL || rune > LAST_SYMBOL {
			res = false
			notValidRunes = append(notValidRunes, rune)
			fmt.Println(rune)
		}
	}
	return res, notValidRunes
}

