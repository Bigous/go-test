package utils

// Any é o tipo que pode receber qualquer valor - não pode ser chave de maps.
type Any interface{}

// AnyArray é um array que tem elementos de tipos variados.
type AnyArray []Any

// MapStringAny um json.
type MapStringAny map[string]Any

// JSON é a representação em mapa de um json qualquer.
type JSON MapStringAny
