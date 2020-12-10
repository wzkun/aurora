package decode

import (
	"bytes"
	"encoding/json"
	"reflect"

	jspb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	stpb "github.com/gogo/protobuf/types"
)

var (
	// PROTO decoder or encoder
	PROTO = &ProtoIO{}
)

// ProtoIO structer
type ProtoIO struct {
	en jspb.Marshaler
	de jspb.Unmarshaler
}

// fixStruct function
func (c *ProtoIO) fixStruct(st *stpb.Struct) {
	if st.Fields == nil {
		st.Fields = make(map[string]*stpb.Value)
	}
}

// MarshalToString return the table name
func (c *ProtoIO) MarshalToString(pb proto.Message) (string, error) {
	if pb == nil || reflect.ValueOf(pb).IsNil() {
		return "", nil
	}
	return c.en.MarshalToString(pb)
}

// MarshalToJSON return the table name
func (c *ProtoIO) MarshalToJSON(pb proto.Message) (json.RawMessage, error) {
	if pb == nil || reflect.ValueOf(pb).IsNil() {
		return json.RawMessage("{}"), nil
	}

	var buf bytes.Buffer
	if err := c.en.Marshal(&buf, pb); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// PBStructFromJSON return the table name
func (c *ProtoIO) PBStructFromJSON(data json.RawMessage) (*stpb.Struct, error) {
	reader := bytes.NewReader(data)
	st := &stpb.Struct{}
	st.Fields = make(map[string]*stpb.Value)
	err := c.de.Unmarshal(reader, st)
	return st, err
}

// PBStructToJSON return the table name
func (c *ProtoIO) PBStructToJSON(st *stpb.Struct) (json.RawMessage, error) {
	if st == nil {
		return json.RawMessage("{}"), nil
	}

	c.fixStruct(st)
	str, err := c.en.MarshalToString(st)
	return json.RawMessage(str), err
}

// MergePBStruct return the table name
func (c *ProtoIO) MergePBStruct(source, dest *stpb.Struct) error {
	c.fixStruct(source)
	c.fixStruct(dest)

	destFields := dest.Fields
	for k, v := range source.Fields {
		if st, ok := v.GetKind().(*stpb.Value_StructValue); ok {
			d, hit := destFields[k]
			if !hit {
				destFields[k] = v
				continue
			}
			// If exist, replace directly, no need to check the type consistant
			dt, ok := d.GetKind().(*stpb.Value_StructValue)
			if !ok {
				destFields[k] = v
				continue
			}
			err := c.MergePBStruct(st.StructValue, dt.StructValue)
			if err != nil {
				return err
			}
		} else {
			destFields[k] = v
		}
	}
	return nil
}


// PBAnyFromJSON return the table name
func (c *ProtoIO) PBAnyFromJSON(data string) (*stpb.Any, error) {
	reader := bytes.NewReader([]byte(data))
	st := &stpb.Any{}
	st.Value = []byte(data)
	err := c.de.Unmarshal(reader, st)
	return st, err
}

// PBAnyFromStruct return the table name
func (c *ProtoIO) PBAnyFromStruct(value interface{}) (*stpb.Any, error) {
	data,_ := JSON.MarshalToString(value)
	st := &stpb.Any{}
	st.Value = []byte(data)
	return st, nil
}