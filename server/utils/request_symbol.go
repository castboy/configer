package utils

import "regexp"

func GetRequestSymbol(symbolName string) string {
	// eg. AUDUSD, AUDUSD20, AUDUSD50, AUDUSD200 ...
	symbReg := regexp.MustCompile("^([A-Z]{6})[0-9]*$")
	match := symbReg.FindStringSubmatch(symbolName)
	if len(match) == 2 {
		return match[1]
	}

	// eg. EUSTX50_200, GER30_200 ...
	symbReg = regexp.MustCompile("^([0-9A-Z]+)_[0-9]+$")
	match = symbReg.FindStringSubmatch(symbolName)
	if len(match) == 2 {
		return match[1]
	}

	// eg. D30EUR20, 200AUD20 ...
	symbReg = regexp.MustCompile("^([0-9A-Z]{3}[A-Z]{3})[0-9]+$")
	match = symbReg.FindStringSubmatch(symbolName)
	if len(match) == 2 {
		return match[1]
	}

	// others, eg. US500, UK100, JPN225 ...
	return symbolName
}
