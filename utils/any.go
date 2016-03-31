package utils

import "fmt"

// Any é o tipo que pode receber qualquer valor - não pode ser chave de maps.
type Any interface{}

// AnyArray é um array que tem elementos de tipos variados.
type AnyArray []Any

// ToString converts a value to its string representation
func ToString(any Any) string {
	switch t := any.(type) {
	case fmt.Stringer:
		return t.String()
	case bool:
		return fmt.Sprintf("%t", t)
	default:
		return fmt.Sprintf("%v", t)
	}
}

// ToInt returns the int form of Any
func ToInt(any Any) (int, error) {
	switch t := any.(type) {
	case int:
		return t, nil
	case int8:
		return int(t), nil
	case int16:
		return int(t), nil
	case int32:
		return int(t), nil
	case int64:
		return int(t), nil
	case float32:
		return int(t), nil
	case float64:
		return int(t), nil
	default:
		return 0, fmt.Errorf("Not a number")
	}
}

// ToInt8 returns the int8 form of Any
func ToInt8(any Any) (int8, error) {
	switch t := any.(type) {
	case int:
		return int8(t), nil
	case int8:
		return int8(t), nil
	case int16:
		return int8(t), nil
	case int32:
		return int8(t), nil
	case int64:
		return int8(t), nil
	case float32:
		return int8(t), nil
	case float64:
		return int8(t), nil
	default:
		return 0, fmt.Errorf("Not a number")
	}
}

// ToInt16 returns the int16 form of Any
func ToInt16(any Any) (int16, error) {
	switch t := any.(type) {
	case int:
		return int16(t), nil
	case int8:
		return int16(t), nil
	case int16:
		return int16(t), nil
	case int32:
		return int16(t), nil
	case int64:
		return int16(t), nil
	case float32:
		return int16(t), nil
	case float64:
		return int16(t), nil
	default:
		return 0, fmt.Errorf("Not a number")
	}
}

// ToInt32 returns the int32 form of Any
func ToInt32(any Any) (int32, error) {
	switch t := any.(type) {
	case int:
		return int32(t), nil
	case int8:
		return int32(t), nil
	case int16:
		return int32(t), nil
	case int32:
		return int32(t), nil
	case int64:
		return int32(t), nil
	case float32:
		return int32(t), nil
	case float64:
		return int32(t), nil
	default:
		return 0, fmt.Errorf("Not a number")
	}
}

// ToInt64 returns the int64 form of Any
func ToInt64(any Any) (int64, error) {
	switch t := any.(type) {
	case int:
		return int64(t), nil
	case int8:
		return int64(t), nil
	case int16:
		return int64(t), nil
	case int32:
		return int64(t), nil
	case int64:
		return int64(t), nil
	case float32:
		return int64(t), nil
	case float64:
		return int64(t), nil
	default:
		return 0, fmt.Errorf("Not a number")
	}
}

// ToFloat32 returns the float32 form of Any
func ToFloat32(any Any) (float32, error) {
	switch t := any.(type) {
	case int:
		return float32(t), nil
	case int8:
		return float32(t), nil
	case int16:
		return float32(t), nil
	case int32:
		return float32(t), nil
	case int64:
		return float32(t), nil
	case float32:
		return float32(t), nil
	case float64:
		return float32(t), nil
	default:
		return 0, fmt.Errorf("Not a number")
	}
}

// ToFloat64 returns the float64 form of Any
func ToFloat64(any Any) (float64, error) {
	switch t := any.(type) {
	case int:
		return float64(t), nil
	case int8:
		return float64(t), nil
	case int16:
		return float64(t), nil
	case int32:
		return float64(t), nil
	case int64:
		return float64(t), nil
	case float32:
		return float64(t), nil
	case float64:
		return float64(t), nil
	default:
		return 0, fmt.Errorf("Not a number")
	}
}
