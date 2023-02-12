package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

type Address struct {
	Street       string
	Number       int64
	Complement   string
	Neighborhood string
	City         string
	State        string
}

type Person struct {
	Id       string
	Name     string
	Age      int64
	Document string
	Address  Address
}

func main() {

	uuidGen := uuid.New()

	maicon := Person{
		Id:       uuidGen.String(),
		Name:     "Maicon Keller",
		Age:      39,
		Document: "99999999999",
		Address: Address{
			Street:       "Times Square",
			Number:       999999,
			Complement:   "som Complement",
			Neighborhood: "Some Neighborhood",
			City:         "New York",
			State:        "NY",
		},
	}

	// first example - marshal generate minified json string
	j, _ := json.Marshal(maicon)
	fmt.Println("Minified", string(j))

	// another example - compact minify some json string
	fmt.Println("Normal:", myJson)
	dst := &bytes.Buffer{}
	_ = json.Compact(dst, []byte(myJson))
	fmt.Println("Minified:", dst.String())
}

const myJson = `{
  "Id": "b88a4fff-7371-4cf2-a3b8-65e92746569e",
  "Name": "Maicon Keller",
  "Age": 39,
  "Document": "99999999999",
  "Address": {
    "Street": "Times Square",
    "Number": 999999,
    "Complement": "som Complement",
    "Neighborhood": "Some Neighborhood",
    "City": "New York",
    "State": "NY"
  }
}`
