/*
Package testschema is a generated package which contains definitions
of structs which represent a YANG schema. The generated schema can be
compressed by a series of transformations (compression was false
in this case).

This package was generated by /usr/local/google/home/wenbli/gocode/src/github.com/openconfig/ygot4/genutil/names.go
using the following YANG input files:
  - yang/testschema.yang
  - yang/refschema.yang

Imported modules were sourced from:
  - ...
*/
package testschema

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/openconfig/goyang/pkg/yang"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"
)

// Binary is a type that is used for fields that have a YANG type of
// binary. It is used such that binary fields can be distinguished from
// leaf-lists of uint8s (which are mapped to []uint8, equivalent to
// []byte in reflection).
type Binary []byte

// YANGEmpty is a type that is used for fields that have a YANG type of
// empty. It is used such that empty fields can be distinguished from boolean fields
// in the generated code.
type YANGEmpty bool

// UnionInt8 is an int8 type assignable to unions of which it is a subtype.
type UnionInt8 int8

// UnionInt16 is an int16 type assignable to unions of which it is a subtype.
type UnionInt16 int16

// UnionInt32 is an int32 type assignable to unions of which it is a subtype.
type UnionInt32 int32

// UnionInt64 is an int64 type assignable to unions of which it is a subtype.
type UnionInt64 int64

// UnionUint8 is a uint8 type assignable to unions of which it is a subtype.
type UnionUint8 uint8

// UnionUint16 is a uint16 type assignable to unions of which it is a subtype.
type UnionUint16 uint16

// UnionUint32 is a uint32 type assignable to unions of which it is a subtype.
type UnionUint32 uint32

// UnionUint64 is a uint64 type assignable to unions of which it is a subtype.
type UnionUint64 uint64

// UnionFloat64 is a float64 type assignable to unions of which it is a subtype.
type UnionFloat64 float64

// UnionString is a string type assignable to unions of which it is a subtype.
type UnionString string

// UnionBool is a bool type assignable to unions of which it is a subtype.
type UnionBool bool

// UnionUnsupported is an interface{} wrapper type for unsupported types. It is
// assignable to unions of which it is a subtype.
type UnionUnsupported struct {
	Value interface{}
}

var (
	SchemaTree map[string]*yang.Entry
	ΛEnumTypes map[string][]reflect.Type
)

func init() {
	var err error
	initΛEnumTypes()
	if SchemaTree, err = UnzipSchema(); err != nil {
		panic("schema error: " + err.Error())
	}
}

// Schema returns the details of the generated schema.
func Schema() (*ytypes.Schema, error) {
	uzp, err := UnzipSchema()
	if err != nil {
		return nil, fmt.Errorf("cannot unzip schema, %v", err)
	}

	return &ytypes.Schema{
		Root:       &Device{},
		SchemaTree: uzp,
		Unmarshal:  Unmarshal,
	}, nil
}

// UnzipSchema unzips the zipped schema and returns a map of yang.Entry nodes,
// keyed by the name of the struct that the yang.Entry describes the schema for.
func UnzipSchema() (map[string]*yang.Entry, error) {
	var schemaTree map[string]*yang.Entry
	var err error
	if schemaTree, err = ygot.GzipToSchema(ySchema); err != nil {
		return nil, fmt.Errorf("could not unzip the schema; %v", err)
	}
	return schemaTree, nil
}

// Unmarshal unmarshals data, which must be RFC7951 JSON format, into
// destStruct, which must be non-nil and the correct GoStruct type. It returns
// an error if the destStruct is not found in the schema or the data cannot be
// unmarshaled. The supplied options (opts) are used to control the behaviour
// of the unmarshal function - for example, determining whether errors are
// thrown for unknown fields in the input JSON.
func Unmarshal(data []byte, destStruct ygot.GoStruct, opts ...ytypes.UnmarshalOpt) error {
	tn := reflect.TypeOf(destStruct).Elem().Name()
	schema, ok := SchemaTree[tn]
	if !ok {
		return fmt.Errorf("could not find schema for type %s", tn)
	}
	var jsonTree interface{}
	if err := json.Unmarshal([]byte(data), &jsonTree); err != nil {
		return err
	}
	return ytypes.Unmarshal(schema, destStruct, jsonTree, opts...)
}

// Device represents the /device YANG schema element.
type Device struct {
	ΛMetadata []ygot.Annotation  `path:"@" ygotAnnotation:"true"`
	Ref       *Refschema_Ref     `path:"ref" module:"refschema"`
	ΛRef      []ygot.Annotation  `path:"@ref" ygotAnnotation:"true"`
	Target    *Testschema_Target `path:"target" module:"testschema"`
	ΛTarget   []ygot.Annotation  `path:"@target" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that Device implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Device) IsYANGGoStruct() {}

// GetOrCreateRef retrieves the value of the Ref field
// or returns the existing field if it already exists.
func (t *Device) GetOrCreateRef() *Refschema_Ref {
	if t.Ref != nil {
		return t.Ref
	}
	t.Ref = &Refschema_Ref{}
	return t.Ref
}

// GetOrCreateTarget retrieves the value of the Target field
// or returns the existing field if it already exists.
func (t *Device) GetOrCreateTarget() *Testschema_Target {
	if t.Target != nil {
		return t.Target
	}
	t.Target = &Testschema_Target{}
	return t.Target
}

// GetRef returns the value of the Ref struct pointer
// from Device. If the receiver or the field Ref is nil, nil
// is returned such that the Get* methods can be safely chained.
func (t *Device) GetRef() *Refschema_Ref {
	if t != nil && t.Ref != nil {
		return t.Ref
	}
	return nil
}

// GetTarget returns the value of the Target struct pointer
// from Device. If the receiver or the field Target is nil, nil
// is returned such that the Get* methods can be safely chained.
func (t *Device) GetTarget() *Testschema_Target {
	if t != nil && t.Target != nil {
		return t.Target
	}
	return nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Device) ΛValidate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["Device"], t, opts...); err != nil {
		return err
	}
	return nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Device) Validate(opts ...ygot.ValidationOption) error {
	return t.ΛValidate(opts...)
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *Device) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }

// ΛBelongingModule returns the name of the module that defines the namespace
// of Device.
func (*Device) ΛBelongingModule() string {
	return ""
}

// Refschema_Ref represents the /refschema/ref YANG schema element.
type Refschema_Ref struct {
	ΛMetadata  []ygot.Annotation                   `path:"@" ygotAnnotation:"true"`
	Reference  map[string]*Refschema_Ref_Reference `path:"reference" module:"refschema"`
	ΛReference []ygot.Annotation                   `path:"@reference" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that Refschema_Ref implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Refschema_Ref) IsYANGGoStruct() {}

// NewReference creates a new entry in the Reference list of the
// Refschema_Ref struct. The keys of the list are populated from the input
// arguments.
func (t *Refschema_Ref) NewReference(Name string) (*Refschema_Ref_Reference, error) {

	// Initialise the list within the receiver struct if it has not already been
	// created.
	if t.Reference == nil {
		t.Reference = make(map[string]*Refschema_Ref_Reference)
	}

	key := Name

	// Ensure that this key has not already been used in the
	// list. Keyed YANG lists do not allow duplicate keys to
	// be created.
	if _, ok := t.Reference[key]; ok {
		return nil, fmt.Errorf("duplicate key %v for list Reference", key)
	}

	t.Reference[key] = &Refschema_Ref_Reference{
		Name: &Name,
	}

	return t.Reference[key], nil
}

// RenameReference renames an entry in the list Reference within
// the Refschema_Ref struct. The entry with key oldK is renamed to newK updating
// the key within the value.
func (t *Refschema_Ref) RenameReference(oldK, newK string) error {
	if _, ok := t.Reference[newK]; ok {
		return fmt.Errorf("key %v already exists in Reference", newK)
	}

	e, ok := t.Reference[oldK]
	if !ok {
		return fmt.Errorf("key %v not found in Reference", oldK)
	}
	e.Name = &newK

	t.Reference[newK] = e
	delete(t.Reference, oldK)
	return nil
}

// GetOrCreateReference retrieves the value with the specified keys from
// the receiver Refschema_Ref. If the entry does not exist, then it is created.
// It returns the existing or new list member.
func (t *Refschema_Ref) GetOrCreateReference(Name string) *Refschema_Ref_Reference {

	key := Name

	if v, ok := t.Reference[key]; ok {
		return v
	}
	// Panic if we receive an error, since we should have retrieved an existing
	// list member. This allows chaining of GetOrCreate methods.
	v, err := t.NewReference(Name)
	if err != nil {
		panic(fmt.Sprintf("GetOrCreateReference got unexpected error: %v", err))
	}
	return v
}

// GetReference retrieves the value with the specified key from
// the Reference map field of Refschema_Ref. If the receiver is nil, or
// the specified key is not present in the list, nil is returned such that Get*
// methods may be safely chained.
func (t *Refschema_Ref) GetReference(Name string) *Refschema_Ref_Reference {

	if t == nil {
		return nil
	}

	key := Name

	if lm, ok := t.Reference[key]; ok {
		return lm
	}
	return nil
}

// AppendReference appends the supplied Refschema_Ref_Reference struct to the
// list Reference of Refschema_Ref. If the key value(s) specified in
// the supplied Refschema_Ref_Reference already exist in the list, an error is
// returned.
func (t *Refschema_Ref) AppendReference(v *Refschema_Ref_Reference) error {
	if v.Name == nil {
		return fmt.Errorf("invalid nil key received for Name")
	}

	key := *v.Name

	// Initialise the list within the receiver struct if it has not already been
	// created.
	if t.Reference == nil {
		t.Reference = make(map[string]*Refschema_Ref_Reference)
	}

	if _, ok := t.Reference[key]; ok {
		return fmt.Errorf("duplicate key for list Reference %v", key)
	}

	t.Reference[key] = v
	return nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Refschema_Ref) ΛValidate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["Refschema_Ref"], t, opts...); err != nil {
		return err
	}
	return nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Refschema_Ref) Validate(opts ...ygot.ValidationOption) error {
	return t.ΛValidate(opts...)
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *Refschema_Ref) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }

// ΛBelongingModule returns the name of the module that defines the namespace
// of Refschema_Ref.
func (*Refschema_Ref) ΛBelongingModule() string {
	return "refschema"
}

// Refschema_Ref_Reference represents the /refschema/ref/reference YANG schema element.
type Refschema_Ref_Reference struct {
	ΛMetadata []ygot.Annotation `path:"@" ygotAnnotation:"true"`
	Name      *string           `path:"name" module:"refschema"`
	ΛName     []ygot.Annotation `path:"@name" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that Refschema_Ref_Reference implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Refschema_Ref_Reference) IsYANGGoStruct() {}

// GetName retrieves the value of the leaf Name from the Refschema_Ref_Reference
// struct. If the field is unset but has a default value in the YANG schema,
// then the default value will be returned.
// Caution should be exercised whilst using this method since when without a
// default value, it will return the Go zero value if the field is explicitly
// unset. If the caller explicitly does not care if Name is set, it can
// safely use t.GetName() to retrieve the value. In the case that the
// caller has different actions based on whether the leaf is set or unset, it
// should use 'if t.Name == nil' before retrieving the leaf's value.
func (t *Refschema_Ref_Reference) GetName() string {
	if t == nil || t.Name == nil {
		return ""
	}
	return *t.Name
}

// ΛListKeyMap returns the keys of the Refschema_Ref_Reference struct, which is a YANG list entry.
func (t *Refschema_Ref_Reference) ΛListKeyMap() (map[string]interface{}, error) {
	if t.Name == nil {
		return nil, fmt.Errorf("nil value for key Name")
	}

	return map[string]interface{}{
		"name": *t.Name,
	}, nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Refschema_Ref_Reference) ΛValidate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["Refschema_Ref_Reference"], t, opts...); err != nil {
		return err
	}
	return nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Refschema_Ref_Reference) Validate(opts ...ygot.ValidationOption) error {
	return t.ΛValidate(opts...)
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *Refschema_Ref_Reference) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }

// ΛBelongingModule returns the name of the module that defines the namespace
// of Refschema_Ref_Reference.
func (*Refschema_Ref_Reference) ΛBelongingModule() string {
	return "refschema"
}

// Testschema_Target represents the /testschema/target YANG schema element.
type Testschema_Target struct {
	ΛMetadata []ygot.Annotation                    `path:"@" ygotAnnotation:"true"`
	Entity    map[string]*Testschema_Target_Entity `path:"entity" module:"testschema"`
	ΛEntity   []ygot.Annotation                    `path:"@entity" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that Testschema_Target implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Testschema_Target) IsYANGGoStruct() {}

// NewEntity creates a new entry in the Entity list of the
// Testschema_Target struct. The keys of the list are populated from the input
// arguments.
func (t *Testschema_Target) NewEntity(Name string) (*Testschema_Target_Entity, error) {

	// Initialise the list within the receiver struct if it has not already been
	// created.
	if t.Entity == nil {
		t.Entity = make(map[string]*Testschema_Target_Entity)
	}

	key := Name

	// Ensure that this key has not already been used in the
	// list. Keyed YANG lists do not allow duplicate keys to
	// be created.
	if _, ok := t.Entity[key]; ok {
		return nil, fmt.Errorf("duplicate key %v for list Entity", key)
	}

	t.Entity[key] = &Testschema_Target_Entity{
		Name: &Name,
	}

	return t.Entity[key], nil
}

// RenameEntity renames an entry in the list Entity within
// the Testschema_Target struct. The entry with key oldK is renamed to newK updating
// the key within the value.
func (t *Testschema_Target) RenameEntity(oldK, newK string) error {
	if _, ok := t.Entity[newK]; ok {
		return fmt.Errorf("key %v already exists in Entity", newK)
	}

	e, ok := t.Entity[oldK]
	if !ok {
		return fmt.Errorf("key %v not found in Entity", oldK)
	}
	e.Name = &newK

	t.Entity[newK] = e
	delete(t.Entity, oldK)
	return nil
}

// GetOrCreateEntity retrieves the value with the specified keys from
// the receiver Testschema_Target. If the entry does not exist, then it is created.
// It returns the existing or new list member.
func (t *Testschema_Target) GetOrCreateEntity(Name string) *Testschema_Target_Entity {

	key := Name

	if v, ok := t.Entity[key]; ok {
		return v
	}
	// Panic if we receive an error, since we should have retrieved an existing
	// list member. This allows chaining of GetOrCreate methods.
	v, err := t.NewEntity(Name)
	if err != nil {
		panic(fmt.Sprintf("GetOrCreateEntity got unexpected error: %v", err))
	}
	return v
}

// GetEntity retrieves the value with the specified key from
// the Entity map field of Testschema_Target. If the receiver is nil, or
// the specified key is not present in the list, nil is returned such that Get*
// methods may be safely chained.
func (t *Testschema_Target) GetEntity(Name string) *Testschema_Target_Entity {

	if t == nil {
		return nil
	}

	key := Name

	if lm, ok := t.Entity[key]; ok {
		return lm
	}
	return nil
}

// AppendEntity appends the supplied Testschema_Target_Entity struct to the
// list Entity of Testschema_Target. If the key value(s) specified in
// the supplied Testschema_Target_Entity already exist in the list, an error is
// returned.
func (t *Testschema_Target) AppendEntity(v *Testschema_Target_Entity) error {
	if v.Name == nil {
		return fmt.Errorf("invalid nil key received for Name")
	}

	key := *v.Name

	// Initialise the list within the receiver struct if it has not already been
	// created.
	if t.Entity == nil {
		t.Entity = make(map[string]*Testschema_Target_Entity)
	}

	if _, ok := t.Entity[key]; ok {
		return fmt.Errorf("duplicate key for list Entity %v", key)
	}

	t.Entity[key] = v
	return nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Testschema_Target) ΛValidate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["Testschema_Target"], t, opts...); err != nil {
		return err
	}
	return nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Testschema_Target) Validate(opts ...ygot.ValidationOption) error {
	return t.ΛValidate(opts...)
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *Testschema_Target) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }

// ΛBelongingModule returns the name of the module that defines the namespace
// of Testschema_Target.
func (*Testschema_Target) ΛBelongingModule() string {
	return "testschema"
}

// Testschema_Target_Entity represents the /testschema/target/entity YANG schema element.
type Testschema_Target_Entity struct {
	ΛMetadata []ygot.Annotation `path:"@" ygotAnnotation:"true"`
	Name      *string           `path:"name" module:"testschema"`
	ΛName     []ygot.Annotation `path:"@name" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that Testschema_Target_Entity implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Testschema_Target_Entity) IsYANGGoStruct() {}

// GetName retrieves the value of the leaf Name from the Testschema_Target_Entity
// struct. If the field is unset but has a default value in the YANG schema,
// then the default value will be returned.
// Caution should be exercised whilst using this method since when without a
// default value, it will return the Go zero value if the field is explicitly
// unset. If the caller explicitly does not care if Name is set, it can
// safely use t.GetName() to retrieve the value. In the case that the
// caller has different actions based on whether the leaf is set or unset, it
// should use 'if t.Name == nil' before retrieving the leaf's value.
func (t *Testschema_Target_Entity) GetName() string {
	if t == nil || t.Name == nil {
		return ""
	}
	return *t.Name
}

// ΛListKeyMap returns the keys of the Testschema_Target_Entity struct, which is a YANG list entry.
func (t *Testschema_Target_Entity) ΛListKeyMap() (map[string]interface{}, error) {
	if t.Name == nil {
		return nil, fmt.Errorf("nil value for key Name")
	}

	return map[string]interface{}{
		"name": *t.Name,
	}, nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Testschema_Target_Entity) ΛValidate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["Testschema_Target_Entity"], t, opts...); err != nil {
		return err
	}
	return nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Testschema_Target_Entity) Validate(opts ...ygot.ValidationOption) error {
	return t.ΛValidate(opts...)
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *Testschema_Target_Entity) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }

// ΛBelongingModule returns the name of the module that defines the namespace
// of Testschema_Target_Entity.
func (*Testschema_Target_Entity) ΛBelongingModule() string {
	return "testschema"
}

var (
	// ySchema is a byte slice contain a gzip compressed representation of the
	// YANG schema from which the Go code was generated. When uncompressed the
	// contents of the byte slice is a JSON document containing an object, keyed
	// on the name of the generated struct, and containing the JSON marshalled
	// contents of a goyang yang.Entry struct, which defines the schema for the
	// fields within the struct.
	ySchema = []byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x58, 0x4d, 0x6f, 0x13, 0x31,
		0x10, 0xbd, 0xef, 0xaf, 0xb0, 0x7c, 0x8e, 0xd4, 0x44, 0x24, 0x4d, 0xd9, 0x5b, 0xa0, 0x45, 0x48,
		0xa5, 0x50, 0x85, 0xdc, 0x23, 0x2b, 0x99, 0x4d, 0x2d, 0x12, 0x6f, 0x64, 0xcf, 0x42, 0x23, 0x94,
		0xff, 0x8e, 0xbc, 0x76, 0x96, 0xfd, 0xb4, 0xd7, 0x01, 0x41, 0x40, 0x3d, 0x54, 0x72, 0xc7, 0xb3,
		0x33, 0x6f, 0xe6, 0x8d, 0xed, 0xa7, 0x7c, 0x8f, 0x08, 0x21, 0x84, 0x7e, 0x64, 0x3b, 0xa0, 0x31,
		0xa1, 0x6b, 0xf8, 0xca, 0x57, 0x40, 0x07, 0xc6, 0x7a, 0xcf, 0xc5, 0x9a, 0xc6, 0x64, 0x64, 0xff,
		0x7d, 0x9b, 0x8a, 0x84, 0x6f, 0x68, 0x4c, 0x86, 0xd6, 0x70, 0xcb, 0x25, 0x8d, 0x89, 0x09, 0x91,
		0x1b, 0x24, 0x24, 0x15, 0x43, 0x25, 0xb6, 0xde, 0x1c, 0x54, 0xb7, 0xaa, 0x09, 0x0a, 0x73, 0x3d,
		0x51, 0xb1, 0xf1, 0x28, 0x21, 0xe1, 0xcf, 0x8d, 0x14, 0xd5, 0x34, 0xaa, 0x96, 0x25, 0xdf, 0xfd,
		0x9c, 0x66, 0x72, 0x05, 0xad, 0x5f, 0x1a, 0x24, 0x70, 0xf8, 0x96, 0x4a, 0x0d, 0x86, 0xee, 0x4d,
		0x92, 0x41, 0xbb, 0xe3, 0x7b, 0xa6, 0x66, 0x72, 0x93, 0xed, 0x40, 0x20, 0x8d, 0x09, 0xca, 0x0c,
		0x3a, 0x1c, 0x4b, 0x5e, 0x1a, 0x53, 0xc3, 0xe9, 0x58, 0xb1, 0x1c, 0x6b, 0x95, 0xd6, 0x5b, 0x5b,
		0x6e, 0x31, 0x48, 0x10, 0xae, 0x5a, 0x4a, 0x0d, 0xb7, 0xae, 0x1d, 0x10, 0xdb, 0xdb, 0xef, 0xa5,
		0xa1, 0x0f, 0x1d, 0xfd, 0x68, 0xe9, 0x4b, 0x4f, 0x30, 0x4d, 0xc1, 0x74, 0xf5, 0xa6, 0xad, 0x9d,
		0xbe, 0x0e, 0x1a, 0xbd, 0x74, 0x16, 0x0e, 0xc2, 0xb4, 0xc9, 0x53, 0xfe, 0xa9, 0x99, 0xb9, 0xb7,
		0xa7, 0x10, 0x4b, 0xee, 0xd0, 0xe3, 0xe6, 0x23, 0x39, 0x84, 0xec, 0x30, 0xd2, 0x43, 0xc9, 0x3f,
		0x7b, 0x08, 0xce, 0x1e, 0x86, 0xe0, 0xa1, 0x70, 0x0f, 0x87, 0x67, 0x48, 0x8a, 0x6c, 0x8b, 0xc3,
		0x1e, 0xc2, 0xfa, 0xbc, 0x05, 0x96, 0x34, 0xaf, 0x57, 0xe7, 0x99, 0x9f, 0xf6, 0xf0, 0x7d, 0x64,
		0xf8, 0xa4, 0xc3, 0x5f, 0xa1, 0x8a, 0x91, 0xc9, 0x0d, 0xa0, 0x5e, 0x81, 0x40, 0x8e, 0x07, 0xbd,
		0xca, 0xc7, 0x30, 0x3a, 0xaf, 0x11, 0x61, 0xe7, 0xe7, 0x1e, 0x0e, 0x9e, 0xb9, 0xa7, 0x1f, 0xb8,
		0xc2, 0x19, 0xa2, 0xe7, 0x9c, 0x3d, 0x70, 0x71, 0xb7, 0x05, 0xcd, 0xa5, 0x72, 0xcf, 0x3c, 0x7d,
		0x60, 0xcf, 0x25, 0xcf, 0xd1, 0xcd, 0x78, 0x7c, 0x3d, 0x1d, 0x8f, 0x87, 0xd3, 0x57, 0xd3, 0xe1,
		0xeb, 0xc9, 0x64, 0x74, 0x3d, 0x9a, 0x38, 0x3e, 0xfe, 0x24, 0xd7, 0x20, 0x61, 0xfd, 0x46, 0xa3,
		0x16, 0xd9, 0x76, 0x1b, 0x54, 0xec, 0x4c, 0x88, 0x14, 0x19, 0xf2, 0x54, 0xb8, 0x6b, 0x51, 0xab,
		0x27, 0xd8, 0xb1, 0xfd, 0x89, 0x24, 0x09, 0x89, 0xb1, 0xe8, 0xd5, 0x95, 0xef, 0xf2, 0x37, 0x11,
		0x50, 0x66, 0x2b, 0xb4, 0x77, 0x0f, 0x9d, 0x9f, 0x02, 0x2c, 0xe7, 0x90, 0xe8, 0x3f, 0x1b, 0x20,
		0xea, 0xc7, 0x9f, 0xfb, 0x41, 0xf3, 0x14, 0xe5, 0x2c, 0xa6, 0xed, 0x41, 0x77, 0x40, 0xaf, 0x02,
		0xfe, 0x09, 0xab, 0x04, 0x89, 0x9a, 0x71, 0xee, 0x16, 0x2b, 0x76, 0xff, 0x0f, 0xe8, 0x15, 0xbc,
		0x40, 0xbd, 0x82, 0xbf, 0x4f, 0xaf, 0x98, 0xdb, 0xc2, 0x2f, 0x56, 0xac, 0xdf, 0x05, 0x28, 0x15,
		0xfc, 0x07, 0x95, 0x0a, 0xbe, 0x28, 0x95, 0x5f, 0x53, 0x2a, 0xf8, 0x1f, 0x2a, 0x15, 0xbc, 0x40,
		0xa5, 0xa2, 0x50, 0x72, 0xb1, 0x09, 0x11, 0x2a, 0x37, 0x2f, 0x02, 0xe3, 0x42, 0x04, 0x06, 0x82,
		0x42, 0xfb, 0x28, 0x5b, 0x35, 0xe8, 0xbc, 0xb4, 0x5b, 0x9e, 0xe9, 0x45, 0x11, 0x61, 0xb9, 0xc8,
		0x23, 0x2c, 0xef, 0x4c, 0x84, 0xbf, 0x21, 0x31, 0x1a, 0xe5, 0xf8, 0x65, 0x46, 0x03, 0x7f, 0xa7,
		0xd4, 0x88, 0x4a, 0x00, 0xbb, 0x80, 0x51, 0xae, 0xde, 0xb1, 0x2f, 0x30, 0x4f, 0xd3, 0xe6, 0x21,
		0xaf, 0x83, 0xa5, 0xe5, 0xad, 0x0a, 0xa6, 0x5b, 0xf3, 0x93, 0x8d, 0x49, 0x18, 0x1d, 0x7f, 0x00,
		0x00, 0x00, 0xff, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x51, 0x1c, 0x29, 0xd1, 0x11, 0x00,
		0x00,
	}
)

// ΛEnumTypes is a map, keyed by a YANG schema path, of the enumerated types that
// correspond with the leaf. The type is represented as a reflect.Type. The naming
// of the map ensures that there are no clashes with valid YANG identifiers.
func initΛEnumTypes() {
	ΛEnumTypes = map[string][]reflect.Type{}
}
