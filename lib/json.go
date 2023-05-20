package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func JsonPrettyPrint(in []byte) {
	var out bytes.Buffer
	// TODO ignoring error, but probably shouldn't
	json.Indent(&out, in, "", "  ")
	fmt.Println(out.String())
}
