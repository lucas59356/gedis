package core

const (
	// TypeWhatever Whatever type
	TypeWhatever = 0
	// TypeBool bool
	TypeBool = 1
	// TypeString string
	TypeString = 2
	// TypeInt int
	TypeInt = 10
)

var (
	// Types Dictionary of types. int -> string
	Types = map[int8]string{
		TypeWhatever: "*",
		TypeBool:     "bool",
		TypeString:   "string",
		TypeInt:      "int",
	}
)

// GuessDataType Guess the type number, TypeWhatever if unknown
func GuessDataType(v interface{}) int8 {
	switch v.(type) {
	case bool:
		return TypeBool
	case string:
		return TypeString
	case int:
		return TypeInt
	default:
		return TypeWhatever
	}
}
