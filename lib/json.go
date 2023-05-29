package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func JsonPrettyPrint(in []byte) {
	var out bytes.Buffer
	// TODO ignoring error, but probably shouldn't
	json.Indent(&out, in, "", "  ")
	fmt.Println(out.String())
}

func JsonFilePrettyPrint(path string, in interface{}) error {
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

func JsonReadFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("cannot open file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var out bytes.Buffer
	err = decoder.Decode(&out)
	if err != nil {
		return nil, fmt.Errorf("cannot decode json: %v", err)
	}
	return out.Bytes(), nil
}
