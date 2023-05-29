package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// MakeJsonPath returns the string path for the given type and symbol
// data is agent/systems/etc
// symbol is the symbol of the agent/system/etc
func MakeJsonPath(data, symbol string) string {
	return data + "/" + symbol + "/" + symbol + ".json"
}

func JsonPrettyPrint(in []byte) {
	var out bytes.Buffer
	// TODO ignoring error, but probably shouldn't
	json.Indent(&out, in, "", "  ")
	fmt.Println(out.String())
}

func JsonFilePrettyPrint(path string, in interface{}) error {
	// The path includes the filename, and want to exclude that from the MkdirAll call
	err := os.MkdirAll(path[:len(path)-len(path[strings.LastIndex(path, "/"):])], 0755)
	if err != nil {
		return fmt.Errorf("cannot create directory: %v", err)
	}
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("cannot create file: %v", err)
	}
	defer file.Close()

	m, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		return fmt.Errorf("cannot marshal json: %v", err)
	}
	_, err = file.Write(m)
	if err != nil {
		return fmt.Errorf("cannot write json to file: %v", err)
	}
	return nil
}

// JsonReadFile reads a json file into the out interface
// out needs to be a pointer to the struct that the json file represents
func JsonReadFile(path string, out interface{}) (error) {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("cannot open file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(out)
	if err != nil {
		return fmt.Errorf("cannot decode json: %v", err)
	}
	return nil
}
