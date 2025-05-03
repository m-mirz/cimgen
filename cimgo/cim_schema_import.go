package cimgo

import (
	"errors"
	"strings"
)

var (
	ErrValueNotFound = errors.New("value not found")
)

type CIMAttribute struct {
	Id              string
	Label           string
	Comment         string
	IsList          bool
	RDFRange        string
	DataType        string
	Domain          string
	Namespace       string
	Origin          string
	Stereotype      string
	InverseRole     string
	AssociationUsed bool
	RDFType         string
}

type CIMType struct {
	Id            string
	Label         string
	Comment       string
	SuperType     string
	Namespace     string
	SuperTypes    []string
	SubClasses    []string
	Attributes    []map[string]interface{}
	EnumInstances []string
	Origin        string
	Stereotype    string
	RDFType       string
	Category      string
}

func extractResource(obj map[string]interface{}, key string) string {
	if v, ok := obj[key]; ok {
		if m, ok := v.(map[string]interface{}); ok {
			return m["@rdf:resource"].(string)
		}
	}
	return ""
}

func extractStringOrResource(obj interface{}) string {
	switch item := obj.(type) {
	case string:
		return item
	case map[string]interface{}:
		return item["@rdf:resource"].(string)
	case []interface{}:
		for _, m := range item {
			if m, ok := m.(map[string]interface{}); ok {
				return m["@rdf:resource"].(string)
			}
		}
	}
	return ""
}

func extractText(obj map[string]interface{}, key string) string {
	if v, ok := obj[key]; ok {
		if m, ok := v.(map[string]interface{}); ok {
			return m["_"].(string)
		}
	}
	return ""
}

func extractValue(obj map[string]interface{}, key string) string {
	if t, ok := obj[key]; ok {
		return t.(string)
	}
	return ""
}

func extractURIEnd(uri string) string {
	l := strings.Split(uri, "#")
	return l[len(l)-1]
}

func extractURIPath(uri string) string {
	l := strings.Split(uri, "#")
	return l[0]
}

func processClass(classMap map[string]interface{}) CIMType {

	typeId := extractValue(classMap, "@rdf:about")
	label := extractText(classMap, "rdfs:label")
	superType := extractResource(classMap, "rdfs:subClassOf")
	comment := extractText(classMap, "rdfs:comment")
	stereotype := extractStringOrResource(classMap["cims:stereotype"])
	category := extractResource(classMap, "cims:belongsToCategory")
	rdfType := extractResource(classMap, "rdf:type")

	return CIMType{
		Id:         extractURIEnd(typeId),
		Label:      label,
		SuperType:  extractURIEnd(superType),
		Comment:    comment,
		Namespace:  extractURIPath(typeId),
		Stereotype: extractURIEnd(stereotype),
		RDFType:    extractURIEnd(rdfType),
		Category:   extractURIEnd(category),
	}
}

func processProperty(classMap map[string]interface{}) CIMAttribute {
	attrId := extractValue(classMap, "@rdf:about")
	rdfType := extractResource(classMap, "rdf:type")
	comment := extractText(classMap, "rdfs:comment")
	stereotype := extractStringOrResource(classMap["cims:stereotype"])
	label := extractText(classMap, "rdfs:label")
	domain := extractResource(classMap, "rdfs:domain")
	dataType := extractResource(classMap, "cims:dataType")
	rdfRange := extractResource(classMap, "rdfs:range")
	inverseRoleName := extractResource(classMap, "cims:inverseRoleName")

	associationUsed := false
	if extractStringOrResource(classMap["cims:AssociationUsed"]) == "Yes" {
		associationUsed = true
	}

	isList := false
	multiplicity := extractResource(classMap, "cims:multiplicity")
	if multiplicity == "http://iec.ch/TC57/1999/rdf-schema-extensions-19990926#M:0..n" || multiplicity == "http://iec.ch/TC57/1999/rdf-schema-extensions-19990926#M:1..n" {
		isList = true
	}

	return CIMAttribute{
		Id:              extractURIEnd(attrId),
		Comment:         comment,
		Stereotype:      extractURIEnd(stereotype),
		Namespace:       extractURIPath(attrId),
		Label:           label,
		Domain:          extractURIEnd(domain),
		DataType:        extractURIEnd(dataType),
		RDFRange:        extractURIEnd(rdfRange),
		RDFType:         extractURIEnd(rdfType),
		AssociationUsed: associationUsed,
		InverseRole:     extractURIEnd(inverseRoleName),
		IsList:          isList,
	}
}
