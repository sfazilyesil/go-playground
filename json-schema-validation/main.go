package main

import (
	"encoding/json"
	"fmt"
	"github.com/guregu/null"
	"github.com/xeipuuv/gojsonschema"
	"os"
	"time"
)

type StructToValidate struct {
	S string    `json:"s"`
	I int64     `json:"i"`
	B bool      `json:"b"`
	F float64   `json:"f"`
	T time.Time `json:"t"`

	ES string `json:"es"`
	EI int64  `json:"ei"`

	NS null.String `json:"ns"`
	NI null.Int    `json:"ni"`
	NB null.Bool   `json:"nb"`
	NF null.Float  `json:"nf"`
	NT null.Time   `json:"nt"`

	SS []string `json:"sliceOfString"`
	SI []int64  `json:"sliceOfInt"`

	IS InnerStruct `json:"innerStruct"`

	OmitMe string `json:"-"`
}

type InnerStruct struct {
	X string `json:"x"`
	Y int64  `json:"y"`
}

const schema = `{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "type": "object",
  "required": ["s","i","b","f","t", "innerStruct", "es"],
  "properties": {
    "s": {"type": "string", "minLength": 2, "maxLength": 5},
    "i": {"type": "number", "minimum": 0, "maximum":999},
    "b": {"type": "boolean"},
    "f": {"type": "number", "minimum": 0, "maximum":999.99, "exclusiveMinimum":true},
    "t": {"type": "string", "format":"date-time"},

    "es": {"enum": ["json", "xml", "protobuf", "avro"]},
    "ei": {"enum": [0, 1]},

    "ns": {"type": ["null", "string"], "maxLength":3},
    "ni": {"type": ["null", "number"], "minimum":-1, "maximum":999999},
    "nb": {"type": ["null", "boolean"]},
    "nf": {"type": ["null", "number"], "minimum":0, "maximum":100, "exclusiveMaxsimum":true},
    "nt": {"type": ["null", "string"], "format":"date-time"},

    "sliceOfString": {"type":["array"], "items":{"type":"string"}, "minItems":1, "uniqueItems":true},
    "sliceOfInt": {"type":["null", "array"], "items":{"type":"number"}},

    "innerStruct": {
      "type": "object",
      "properties": {
        "x": {"type": "string"},
        "y": {"type": "number", "minimum": 9}
      }
    }
  }
}`

func main() {
	fmt.Println("---begin---")

	docToValidate := StructToValidate{
		S: "abcde",    // "abcde"
		I: 999,        // 999
		B: true,       // true
		F: 999.99,     // 999.99
		T: time.Now(), // 2017-07-10T22:21:03.936110578+03:00

		ES: "protobuf", // "protobuf"
		EI: 1,          // 1

		NS: null.StringFrom("abc"),    // abc
		NI: null.Int{},                // null
		NB: null.Bool{},               // null
		NF: null.FloatFrom(99.999999), // 99.999999
		NT: null.TimeFrom(time.Now()), // 2017-07-10T22:21:03.936110578+03:00

		SS: []string{"abc", "def"}, // ["abc", "def"]

		IS: InnerStruct{
			X: "aaaaa", // "aaaaa"
			Y: 9,       // 9
		},

		OmitMe: "aaaaaaaaaaaaaaaaaaaaaaa",
	}

	jsonStr, _ := json.Marshal(docToValidate)
	fmt.Println("json doc to validate: \n", string(jsonStr))

	//schemaLoader := gojsonschema.NewStringLoader(schema)
	schemaLoader, err := loadFromSchemaFile("schema.json")
	if err != nil {
		fmt.Println("exiting... err: ", err)
		return
	}

	docLoader := gojsonschema.NewGoLoader(docToValidate)

	res, err := gojsonschema.Validate(schemaLoader, docLoader)
	if err != nil {
		fmt.Println("exiting... err: ", err)
		return
	}

	fmt.Println("isValid: ", res.Valid())
	fmt.Printf("validation errors: %+v \n", res.Errors())

	fmt.Println("---end---")
}

func loadFromSchemaFile(fn string) (gojsonschema.JSONLoader, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return gojsonschema.NewReferenceLoader("file://" + wd + "/src/github.com/sfazilyesil/go-playground/json-schema-validation/" + fn), nil
}
