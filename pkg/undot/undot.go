package undot

import "github.com/latiif/ara/pkg/arabic"

var undotted = map[string]string{
	"أ": "ا",
	"إ": "ا",
	"آ": "ا",
	"ب": "ٮ",
	"ت": "ٮ",
	"ث": "ٮ",
	"ج": "ح",
	"خ": "ح",
	"ذ": "د",
	"ز": "ر",
	"ش": "س",
	"ض": "ص",
	"ظ": "ط",
	"غ": "ع",
	"ف": "ڡ",
	"ق": "ٯ",
	"ن": "ں",
	"ي": "ى",
	"ة": "ه",
}

// Undot replaces dotted letters with undotted ones
func Undot(str string) string {
	str = arabic.RemoveTashkeel(str)
	result := ""
	for _, v := range str {

		if muhmal, found := undotted[string(v)]; found {
			result += muhmal
			continue
		}
		result += string(v)
	}
	return result
}
