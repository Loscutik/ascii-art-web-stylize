package asciiart

import (
	"fmt"
	"testing"
)

func TestTextToArt(t *testing.T) {
	str := "+y"
	res := "                \n   _            \n _| |_   _   _  \n|_   _| | | | | \n  |_|   | |_| | \n         \\__, | \n         __/ /  \n        |___/   \n"
	aStr,err := TextToArt(str, "../banners/standard.txt")
	if err!=nil || aStr != res {
		fmt.Println("want:")
		fmt.Println(res)
		fmt.Println("result is:")
		fmt.Println(aStr)
		t.Fatal("!=res")
	}
}
