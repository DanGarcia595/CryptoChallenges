package Set1

import (
	"fmt"
	"strings"
)

func Keqvparse(plaintext string) (m map[string]string) {
	entries := strings.Split(plaintext, "&")
	m = make(map[string]string)
	for _, element := range entries {
		if strings.Contains(element, "=") {
			kvpair := strings.Split(element, "=")
			key := kvpair[0]
			value := kvpair[1]
			m[key] = value
		}
	}
	for key, value := range m {
		fmt.Println(key, ": ", value)
	}
	return
}
