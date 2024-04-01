package xml2json

import (
	"encoding/json"
	"encoding/xml"
	"eternity/types"
	"fmt"
	"io"
	"os"
	"strings"
)

func CreateJsonFileFromXML(xmlFile string, jsonFile string) error {
	jsonData, err := xml2json(xmlFile)
	if err != nil {
		return err
	}
	err = writeJsonFile(jsonData, jsonFile)
	if err != nil {
		return err
	}
	return nil
}

func xml2json(file string) ([]byte, error) {
	xmlData, err := readXmlFile(file)
	if err != nil {
		fmt.Println("Error reading xml file", file)
		return nil, err
	}
	value, err := parseXML(xmlData)
	if err != nil {
		fmt.Println("Error parsing xml data", file)
		return nil, err
	}
	jsonData, err := json.Marshal(value)
	if err != nil {
		fmt.Println("Error marshalling xml data", file)
		return nil, err
	}
	return jsonData, nil
}

func parseXML(data []byte) (*types.HealthData, error) {
	var value types.HealthData
	err := xml.Unmarshal(data, &value)
	if err != nil {
		fmt.Println("Error decoding xml data")
		return nil, err
	}
	return &value, nil
}

func readXmlFile(file string) ([]byte, error) {
	if strings.HasSuffix(file, ".xml") {
		return readFile(file)
	} else {
		return readFile(file + ".xml")
	}
}

func readFile(filepath string) ([]byte, error) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file", file)
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file", file)
		return nil, err
	}
	return data, nil
}

func writeJsonFile(data []byte, file string) error {
	if strings.HasSuffix(file, ".json") {
		return writeFile(data, file)
	} else {
		return writeFile(data, file+".json")
	}
}

func writeFile(data []byte, file string) error {
	outputFile, err := os.Create(file)
	if err != nil {
		fmt.Println("Error creating file", file)
		return err
	}
	defer outputFile.Close()
	_, err = outputFile.Write(data)
	if err != nil {
		fmt.Println("Error writing data", file)
		return err
	}
	return nil
}
