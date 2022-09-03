package fluid

import (
	"bytes"
	"strings"
)

func isUpperCase(r rune) bool {
	if strings.ContainsRune("ABCDEFGHIJKLMNOPQRSTUVWXYZ", r) {
		return true
	}
	return false
}

func containsAnyUpperCase(text string) bool {
	l := len(text)

	if l <= 0 {
		return false
	}

	i := 0
	for i < l {
		if isUpperCase(rune(text[i])) {
			return true
		}
		i++
	}

	return false
}

func keep(text string, keepset string) string {
	l := len(text)
	if l <= 0 || len(keepset) <= 0 {
		return text
	}

	keep := bytes.NewBufferString("")
	i := 0
	for i < l {
		r := rune(text[i])
		if strings.ContainsRune(keepset, r) {
			if _, err := keep.WriteRune(r); err != nil {
				panic(err)
			}
		}
		i++
	}

	return keep.String()
}

func caseSensitiveToKebab(caseSensitive string) string {
	l := len(caseSensitive)
	if l <= 1 {
		return strings.ToLower(caseSensitive)
	}
	if !containsAnyUpperCase(caseSensitive[1:]) {
		return strings.ToLower(caseSensitive)
	}

	kebab := bytes.NewBufferString("")
	kebab.WriteRune(rune(caseSensitive[0]))
	i := 1
	for i < l {
		r := rune(caseSensitive[i])
		if isUpperCase(r) {
			kebab.WriteString("-")
		}
		kebab.WriteRune(r)
		i++
	}

	return strings.ToLower(kebab.String())
}

func kebabCase(anyCase string) string {
	l := len(anyCase)

	if l <= 0 {
		return anyCase
	}

	// strip special characters (expect '-' and '_')
	anyCase = keep(anyCase, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_ ")

	if strings.ContainsRune(anyCase, ' ') {
		// process title case
		anyCase = strings.Replace(anyCase, " ", "-", -1)
	} else if containsAnyUpperCase(anyCase) {
		// process pascal & camel case
		anyCase = caseSensitiveToKebab(anyCase)
	} else if strings.ContainsRune(anyCase, '_') {
		// process snake case
		anyCase = strings.Replace(anyCase, "_", "-", -1)
	}

	return strings.Replace(strings.Trim(keep(strings.ToLower(anyCase), "abcdefghijklmnopqrstuvwxyz0123456789-"), "-"), "--", "-", -1)
}

func camelCase(kebab string) (camelCase string) {
	kebab = kebabCase(kebab)

	isToUpper := false
	for _, runeValue := range kebab {
		if isToUpper {
			camelCase += strings.ToUpper(string(runeValue))
			isToUpper = false
		} else {
			if runeValue == '-' {
				isToUpper = true
			} else {
				camelCase += string(runeValue)
			}
		}
	}
	return
}

func pascalCase(kebab string) (pascalCase string) {
	kebab = kebabCase(kebab)

	isToUpper := false
	for i, runeValue := range kebab {
		if isToUpper || i == 0 {
			pascalCase += strings.ToUpper(string(runeValue))
			isToUpper = false
		} else {
			if runeValue == '-' {
				isToUpper = true
			} else {
				pascalCase += string(runeValue)
			}
		}
	}
	return
}

func snakeCase(kebab string) string {
	return strings.Replace(strings.ToLower(kebabCase(kebab)), "-", "_", -1)
}

func titleCase(kebab string) (titleCase string) {
	kebab = kebabCase(kebab)

	isToUpper := false
	for i, runeValue := range kebab {
		if isToUpper || i == 0 {
			titleCase += strings.ToUpper(string(runeValue))
			isToUpper = false
		} else {
			if runeValue == '-' {
				isToUpper = true
				titleCase += " "
			} else {
				titleCase += string(runeValue)
			}
		}
	}
	return
}
