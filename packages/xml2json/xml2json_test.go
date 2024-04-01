package xml2json

import (
	"testing"
)

func TestXml2json(t *testing.T) {
	error := CreateJsonFileFromXML("../../data/export.xml", "../../data/export.json")
	if error != nil {
		t.Error("Expected Xml2json not to return an error")
	}
}
