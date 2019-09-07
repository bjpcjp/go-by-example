/* json.go
- built-in JSON support incl to/from custom datatypes.

more info:
http://blog.golang.org/2011/01/json-and-go.html

*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// use two structs to demo encode/decode abilities.

type response1 struct {
	Page   int
	Fruits []string
}

// only exported fields will be encoded/decoded.
// fields must start with Capital Letters to be exported.

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// encoding basic datatypes

	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// encoding slices & maps
	// (to JSON arrays & objects, as expected.)

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// auto-encoding custom datatypes
	// only includes exported fields
	// by default uses same names as JSON keys.

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// using tags on struct field declarations
	// to customize encoded JSON key names.

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// let's look at decoding JSON data into Go values.
	// ex: generic data structure.

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// need to provide a variable for JSON pkg to put
	// the decoded data.

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	// accessing nested data requires
	// a series of conversions.

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// decoding JSON into custom data types.
	// this allows more type-safety into our code.

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
