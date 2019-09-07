/* xml.go
- XML support with the encoding.xml package.
*/

package main

import (
	"encoding/xml"
	"fmt"
)

// this type will be mapped to XML.
// XMLName = name of XML element for this struct.
// id,attr = Id field is an XML attribute (not nested element)

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// Emit XML representing our plant;
	// use MarshalIndent for human-readable output.

	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	// add generic XML header to output

	fmt.Println(xml.Header + string(out))

	// use Unmarshal to parse bytestream with XML
	// into a data structure.

	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}

	// to add generic XML header to output,
	// append it explicitly.

	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	// parent>child>plant field tag tells the encoder
	// to nest all plants under <parent><child>

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
