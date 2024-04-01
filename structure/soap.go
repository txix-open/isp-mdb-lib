package structure

type MdmObject struct {
	ObjectType   string
	RelationType *string
	Attribute    []*Attribute
	Relations    *Relations
}

type Reference struct {
	ObjectType   string
	RelationType *string
	ObjectId     *string
}

type Relations struct {
	Object []*MdmObject
	Ref    []*Reference
}

type Attribute struct {
	Name  *string
	Value []*string
}

type Object struct {
	Name      *string
	Attribute []*Attribute
}

type EntryType struct {
	EntryName *string
	Seq       *int64
	Attribute []*Attribute
	Object    []*Object
}
