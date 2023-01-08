package asciiart

import (
	"fmt"
	"os"
	"strings"
)

const (
	SYMBOL_HEIGHT = 8
	FIRST_SYMBOL  = ' '
	LAST_SYMBOL   = '~'
)

type (
	ArtString [SYMBOL_HEIGHT]string
	ArtFont []ArtString
)

/*
reads a file with an ascii graphic fonts representing characters from ' '(space) to '~' and returns a map which keeps the characters
maps's key is the character (type rune), map's element is a value of the type ArtString= [SYMBOL_HEIGHT]string
It will return an error if an error occures during oppening readin given file
*/
func GetArtFont(fileName string) (aFont ArtFont, err error) {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	
	str := string(bytes[1:])
	letters := strings.Split(str, "\n\n") // divide by ascii letters
	for nl, letter := range letters {
		lines := strings.Split(letter, "\n") 

		var aLetter ArtString
		l := copy(aLetter[:], lines)
		if l != SYMBOL_HEIGHT {
			return nil, fmt.Errorf("error reading the character from a file: %s; character: %c", fileName, nl+FIRST_SYMBOL)
		}

		aFont = append(aFont, aLetter)
	}
	
	return
}
