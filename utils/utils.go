package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func Pretty(src []byte) string {
	var b bytes.Buffer
	json.Indent(&b, src, "", "\t")
	return b.String()
}

func MergeABI(abi [][]ABIElement) []ABIElement {
	if len(abi) == 1 {
		return abi[0]
	}
	m := make(map[string]ABIElement)

	for i := 0; i < len(abi); i++ {
		elements := abi[i]
		for _, element := range elements {
			m[fmt.Sprintf("%s-%s", element.Type, element.Name)] = element
		}
	}

	var merged []ABIElement
	for _, value := range m {
		merged = append(merged, value)
	}

	return merged
}

func ReadABI(path string) [][]ABIElement {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Printf("Scan dir error %s\n", err.Error())
		os.Exit(1)
	}
	var abis [][]ABIElement
	for index, file := range files {
		if !strings.Contains(file.Name(), ".json") {
			continue
		}
		filePath := path + "/" + file.Name()
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Printf("Failed to read file: %s - %s", filePath, err)
			continue
		}
		var abiElement []ABIElement
		err = json.Unmarshal(content, &abiElement)
		if err != nil {
			fmt.Printf("Unmarshal abi file error: %s\n", err.Error())
			os.Exit(1)
		}
		log.Printf("abi %d has %d elements", index, len(abiElement))
		abis = append(abis, abiElement)
	}
	return abis
}

type ABIElement struct {
	Inputs []interface{} `json:"inputs"`
	Name   string        `json:"name"`
	Type   string        `json:"type"`
}
