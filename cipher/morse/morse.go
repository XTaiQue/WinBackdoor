package morse

import "strings"

var latin = []string{"a",
	"b",
	"c",
	"d",
	"e",
	"f",
	"g",
	"h",
	"i",
	"j",
	"k",
	"l",
	"m",
	"n",
	"o",
	"p",
	"q",
	"r",
	"s",
	"t",
	"u",
	"v",
	"w",
	"x",
	"y",
	"z",
	"0",
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	" ",
}

var latinA = []string{"A",
	"B",
	"C",
	"D",
	"E",
	"F",
	"G",
	"H",
	"I",
	"J",
	"K",
	"L",
	"M",
	"N",
	"O",
	"P",
	"Q",
	"R",
	"S",
	"T",
	"U",
	"V",
	"W",
	"X",
	"Y",
	"Z",
	"0",
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	" ",
}

var morse = []string{".-",
	"-...",
	"-.-.",
	"-..",
	".",
	"..-.",
	"--.",
	"....",
	"..",
	".---",
	"-.-",
	".-..",
	"--",
	"-.",
	"---",
	".--.",
	"--.-",
	".-.",
	"...",
	"-",
	"..-",
	"...-",
	".--",
	"-..-",
	"-.--",
	"--..",
	"-----",
	".----",
	"..---",
	"...--",
	"....-",
	".....",
	"-....",
	"--...",
	"---..",
	"----.",
	"/",
}

func MorseCipher(lstr string) string {
	var ss string = lstr
	//vv := []string{lstr}
	//vvv := strings.Join(vv, " ")
	for i := 0; i < len(morse); i++ {
		s := strings.ReplaceAll(ss, latin[i], morse[i])
		d := strings.ReplaceAll(s, latinA[i], morse[i])
		ss = d
	}

	vv := []string{ss}
	return strings.Join(vv, " ")
	//return ss
}

func MorseUnCipher(lstr string) string {
	var ss string = lstr
	for i := 0; i < len(morse); i++ {
		s := strings.ReplaceAll(ss, morse[i], latin[i])
		d := strings.ReplaceAll(s, morse[i], latinA[i])
		ss = d
	}
	return ss
}
