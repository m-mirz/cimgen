package cimgen

import "encoding/xml"

// RDF is the root element of a CIM RDF file.
type RDF struct {
	XMLName      xml.Name      `xml:"RDF"`
	Base         string        `xml:"base,attr"`
	Descriptions []Description `xml:"Description"`
}

// Description represents an rdf:Description element, which can define a class, property, enum, etc.
type Description struct {
	About                string        `xml:"about,attr"`
	Type                 *Resource     `xml:"type"`
	Label                *Label        `xml:"label"`
	Comment              *Comment      `xml:"comment"`
	Stereotype           *Stereotype   `xml:"stereotype"`
	BelongsToCategory    *Resource     `xml:"belongsToCategory"`
	DataType             *Resource     `xml:"dataType"`
	Domain               *Resource     `xml:"domain"`
	Range                *Resource     `xml:"range"`
	SubClassOf           *Resource     `xml:"subClassOf"`
	InverseRoleName      *Resource     `xml:"inverseRoleName"`
	Multiplicity         *Resource     `xml:"multiplicity"`
	AssociationUsed      *TextValue    `xml:"AssociationUsed"`
	IsFixed              *TextValue    `xml:"isFixed"`
	Title                *TextValue    `xml:"title"`
	Keyword              string        `xml:"keyword"`
	VersionIRI           *Resource     `xml:"versionIRI"`
	VersionInfo          *TextValue    `xml:"versionInfo"`
	Namespaces           []xml.Attr    `xml:",any,attr"`
}

// Resource represents an element with an rdf:resource attribute.
type Resource struct {
	Resource string `xml:"resource,attr"`
}

// Label represents an rdfs:label element.
type Label struct {
	Lang string `xml:"lang,attr"`
	Text string `xml:",chardata"`
}

// Comment represents an rdfs:comment element.
type Comment struct {
	Datatype string `xml:"datatype,attr"`
	Text     string `xml:",chardata"`
}

// Stereotype represents a cims:stereotype element.
type Stereotype struct {
	Resource string `xml:"resource,attr"`
	Text     string `xml:",chardata"`
}

// TextValue represents a simple text element like dcat:keyword or dcterms:title.
type TextValue struct {
	Lang string `xml:"lang,attr"`
	Text string `xml:",chardata"`
}