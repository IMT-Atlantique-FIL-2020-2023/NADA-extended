// Package dotenv implements a koanf.Parser that parses DOTENV bytes as conf maps.
package dotenv

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
	"github.com/knadh/koanf/maps"
)

// DotEnv implements a DotEnv parser.
type DotEnv struct{}

// Parser returns a DOTENV Parser.
func Parser() *DotEnv {
	return &DotEnv{}
}

// Unmarshal parses the given DOTENV bytes.
func (p *DotEnv) Unmarshal(b []byte) (map[string]interface{}, error) {
	// Unmarshal DOTENV from []byte
	r, err := godotenv.Unmarshal(string(b))
	if err != nil {
		return nil, err
	}

	// Convert a map[string]string to a map[string]interface{}
	mp := make(map[string]interface{})
	for k, v := range r {
		k, vV := TransformKeyWithValue(k, v)
		mp[k] = vV
	}
	return maps.Unflatten(mp, "."), err
}

// Marshal marshals the given config map to DOTENV bytes.
func (p *DotEnv) Marshal(o map[string]interface{}) ([]byte, error) {
	// Convert a map[string]interface{} to a map[string]string
	mp := make(map[string]string)
	for k, v := range o {
		mp[k] = fmt.Sprint(v)
	}

	// Unmarshal to string
	out, err := godotenv.Marshal(mp)
	if err != nil {
		return nil, err
	}

	// Convert to []byte and return
	return []byte(out), nil
}

func TransformKeyWithValue(s string, value string) (string, interface{}) {
	key := strings.Replace(strings.ToLower(
		strings.TrimPrefix(s, "NADA_SERVE_")), "_", ".", -1)
	trimed := strings.TrimSuffix(key, "..array")
	if trimed != key {
		return trimed, strings.Split(value, ",")
	}
	return key, value

}
