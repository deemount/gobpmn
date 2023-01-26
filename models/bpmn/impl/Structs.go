package impl

type CoreID struct {
	ID string `xml:"id,attr,omitempty" json:"id" csv:"ID"`
}

type CoreInnerID struct {
	ID string `xml:",innerxml,omitempty" json:"id" csv:"ID"`
}

type BaseAttributes struct {
	ID   string `xml:"id,attr,omitempty" json:"id" csv:"ID"`
	Name string `xml:"name,attr,omitempty" json:"name,omitempty" csv:"NAME"`
}
