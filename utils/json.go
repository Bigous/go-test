package utils

import (
	"bytes"
	"encoding/json"
)

// JSON é a representação em mapa de um json qualquer.
type JSON map[string]interface{}

// TODO: Performance of JSON.String could be improved if we detect JSON suboject and call the recusion directly instead of using another String transformation.
// func recursionBufferize(j *JSON, b *bytes.Buffer) {
// 	var i int
// 	for k, v := range *j {
// 		if i > 0 {
// 			b.WriteRune('"')
// 		}
// 		b.WriteRune('"')
// 		b.WriteString(k)
// 		b.WriteString("\": ")
// 		switch vv := v.(type) {
// 		case string:
// 			b.WriteRune('"')
// 			b.WriteString(vv)
// 			b.WriteRune('"')
// 		default:
// 			b.WriteString(ToString(v))
// 		}
// 		i++
// 	}
// }

// String returns the string representation of a JSON object
func (thisJson JSON) String() string {
	var (
		buf bytes.Buffer
		i   int
	)
	buf.WriteRune('{')
	for k, v := range thisJson {
		if i > 0 {
			buf.WriteRune(',')
		}
		buf.WriteRune('"')
		buf.WriteString(k)
		buf.WriteString("\": ")
		switch vv := v.(type) {
		case string:
			buf.WriteRune('"')
			buf.WriteString(vv)
			buf.WriteRune('"')
		case map[string]interface{}:
			buf.WriteString(JSON(vv).String())
		default:
			buf.WriteString(ToString(v))
		}
		i++
	}
	buf.WriteRune('}')
	return buf.String()
}

// ToString returns the string representation of a JSON object
func (thisJson *JSON) ToString() string {
	return thisJson.String()
}

// Stringfy returns the string representation of a JSON object
func (thisJson *JSON) Stringfy() string {
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

// ParseJSON try to parse a string as a JSON object
func ParseJSON(str string) (ret JSON, err error) {
	err = json.Unmarshal([]byte(str), &ret)
	if err == nil {
		normalizaJSON(&ret)
	}
	return
}
