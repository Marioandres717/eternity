package xml2json

import (
	"encoding/json"
	"encoding/xml"
	"eternity/types"
	"fmt"
	"io"
	"os"
)

func CreateJsonFileFromXML(xmlFile string, jsonFile string) error {
	jsonData, err := xml2json(xmlFile)
	if err != nil {
		return err
	}
	err = writeFile(jsonData, jsonFile)
	if err != nil {
		return err
	}
	return nil
}

func xml2json(file string) ([]byte, error) {
	xmlData, err := readFile(file)
	if err != nil {
		fmt.Println("Error reading xml file", file)
		return nil, err
	}
	var value types.HealthData
	err = xml.Unmarshal(xmlData, &value)
	if err != nil {
		fmt.Println("Error decoding xml file", file)
		return nil, err
	}
	jsonData, err := json.Marshal(value)
	if err != nil {
		fmt.Println("Error marshalling xml data", file)
		return nil, err
	}
	return jsonData, nil
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

func writeFile(data []byte, file string) error {
	outputFile, err := os.Create(file)
	if err != nil {
		fmt.Println("Error creating json file", file)
		return err
	}
	defer outputFile.Close()
	_, err = outputFile.Write(data)
	if err != nil {
		fmt.Println("Error writing json data", file)
		return err
	}
	return nil
}
