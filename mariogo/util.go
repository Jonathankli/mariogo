package mariogo

import (
	"encoding/json"
	"fmt"
	"os"
)

func Dump(data any) {
	s, _ := json.MarshalIndent(data, "", "    ")
	fmt.Println(string(s))
}

func Getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
