package utils

import (
	"bytes"
	"encoding/json"
)

// JSON é a representação em mapa de um json qualquer.
type JSON map[string]interface{}

func writeJSONString(v interface{}, buf *bytes.Buffer) {
	switch vv := v.(type) {
	case string:
		buf.WriteRune('"')
		buf.WriteString(vv)
		buf.WriteRune('"')
	case JSON:
		buf.WriteRune('{')
		i := 0
		for k, o := range vv {
			if i > 0 {
				buf.WriteRune(',')
			}
			writeJSONString(k, buf)
			buf.WriteRune(':')
			writeJSONString(o, buf)
			i++
		}
		buf.WriteRune('}')
	case map[string]interface{}:
		buf.WriteRune('{')
		i := 0
		for k, o := range vv {
			if i > 0 {
				buf.WriteRune(',')
			}
			writeJSONString(k, buf)
			buf.WriteRune(':')
			writeJSONString(o, buf)
			i++
		}
		buf.WriteRune('}')
	case []interface{}:
		buf.WriteRune('[')
		for i, o := range vv {
			if i > 0 {
				buf.WriteRune(',')
			}
			writeJSONString(o, buf)
		}
		buf.WriteRune(']')
	default:
		buf.WriteString(ToString(v))
	}
}

// String returns the string representation of a JSON object
func (thisJson *JSON) String() string {
	var (
		buf bytes.Buffer
	)
	writeJSONString(*thisJson, &buf)
	return buf.String()
}

// ToString returns the string representation of a JSON object
func (thisJson JSON) ToString() string {
	return thisJson.String()
}

// Stringfy returns the string representation of a JSON object
func (thisJson JSON) Stringfy() string {
	return thisJson.String()
}

func normalizaJSON(j *JSON) {
	for k, v := range *j {
		if vv, ok := v.(map[string]interface{}); ok {
			tmp := JSON(vv)
			normalizaJSON(&tmp)
			(*j)[k] = &tmp
		}
	}
}

// ParseJSON try to parse a []byte as a JSON object
func ParseJSON(buffer []byte) (ret JSON, err error) {
	err = json.Unmarshal(buffer, &ret)
	if err == nil {
		normalizaJSON(&ret)
	}
	return
}
