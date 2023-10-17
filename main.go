package main

import (
	"fmt"
	"strconv"
	"strings"
)

var win1251Ru = map[string]uint{
	"А": 0xC0, "Б": 0xC1, "В": 0xC2, "Г": 0xC3, "Д": 0xC4, "Е": 0xC5, "Ё": 0xA8, "Ж": 0xC6, "З": 0xC7, "И": 0xC8, "Й": 0xC9, "К": 0xCA, "Л": 0xCB, "М": 0xCC, "Н": 0xCD, "О": 0xCE, "П": 0xCF,
	"Р": 0xD0, "С": 0xD1, "Т": 0xD2, "У": 0xD3, "Ф": 0xD4, "Х": 0xD5, "Ц": 0xD6, "Ч": 0xD7, "Ш": 0xD8, "Щ": 0xD9, "Ъ": 0xDA, "Ы": 0xDB, "Ь": 0xDC, "Э": 0xDD, "Ю": 0xDE, "Я": 0xDF,
	"а": 0xE0, "б": 0xE1, "в": 0xE2, "г": 0xE3, "д": 0xE4, "е": 0xE5, "ё": 0xB8, "ж": 0xE6, "з": 0xE7, "и": 0xE8, "й": 0xE9, "к": 0xEA, "л": 0xEB, "м": 0xEC, "н": 0xED, "о": 0xEE, "п": 0xEF,
	"р": 0xF0, "с": 0xF1, "т": 0xF2, "у": 0xF3, "ф": 0xF4, "х": 0xF5, "ц": 0xF6, "ч": 0xF7, "ш": 0xF8, "щ": 0xF9, "ъ": 0xFA, "ы": 0xFB, "ь": 0xFC, "э": 0xFD, "ю": 0xFE, "я": 0xFF,
}

type Text string

func (s *Text) getWin1251Codes() []uint {
	var codes []uint
	for _, r := range *s {
		codes = append(codes, win1251Ru[string(r)])
	}
	return codes
}

func getWin1251Text(codes []uint) Text {
	var result string
	for _, code := range codes {
		for k, v := range win1251Ru {
			if code == v {
				result += k
			}
		}
	}
	return Text(result)
}

func convertToBinary(codes []uint) []string {
	var binaryCodes []string
	for _, code := range codes {
		binaryCodes = append(binaryCodes, fmt.Sprintf("%08b", code))
	}
	return binaryCodes
}

func convertToCodes(binaryCodes []string) []uint {
	var result []uint
	for _, code := range binaryCodes {
		i, _ := strconv.ParseUint(code, 2, 64)
		result = append(result, uint(i))
	}
	return result
}

func DecodeGammaCypher(encodedData Text, Salt []uint) Text {
	codes := encodedData.getWin1251Codes()
	fmt.Println(codes)

	binaryCodesText := convertToBinary(codes)
	fmt.Println(binaryCodesText)

	binarySalt := convertToBinary(Salt)
	fmt.Println(binarySalt)

	for len(binaryCodesText) > len(binarySalt) {
		binarySalt = append(binarySalt, binarySalt...)
	}

	var result []string
	for i := 0; i < len(binaryCodesText); i++ {
		result = append(result, XOR(binaryCodesText[i], binarySalt[i]))
	}
	fmt.Println(result)

	codes = convertToCodes(result)
	fmt.Println(codes)

	return getWin1251Text(codes)
}

func XOR(a, b string) string {
	var result string
	if len(b) < len(a) {
		b = b + strings.Repeat("0", len(a)-len(b))
	}
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			result += "0"
		} else {
			result += "1"
		}
	}
	return result
}

func main() {
	var s Text = "ВЧСКЛЭСОШВКВУЬЯ"
	salt := []uint{3, 18, 21, 7, 5, 12}

	plainText := DecodeGammaCypher(s, salt)
	fmt.Println(plainText)
}
