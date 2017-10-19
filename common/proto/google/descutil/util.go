// Copyright 2015 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package descutil

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/golang/protobuf/proto"
	pb "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

// splitFullName splits package name and service/type name.
func splitFullName(fullName string) (pkg string, name string) {
	lastDot := strings.LastIndex(fullName, ".")
	if lastDot < 0 {
		return "", fullName
	}
	return fullName[:lastDot], fullName[lastDot+1:]
}

////////////////////////////////////////////////////////////////////////////////
// FileDescriptorSet

// FindFile searches for a FileDescriptorProto by name.
func FindFile(s *pb.FileDescriptorSet, name string) int {
	for i, x := range s.GetFile() {
		if x.GetName() == name {
			return i
		}
	}
	return -1
}

// Resolve searches for an object by full name. obj can be one of
// *ServiceDescriptorProto,
// *MethodDescriptorProto,
// *DescriptorProto,
// *FieldDescriptorProto,
// *DescriptorProto,
// *EnumDescriptorProto,
// *EnumValueDescriptorProto or
// nil
//
// For path, see comment in SourceCodeInfo message.
func Resolve(s *pb.FileDescriptorSet, fullName string) (file *pb.FileDescriptorProto, obj interface{}, path []int) {
	if fullName == "" {
		return nil, nil, nil
	}
	pkg, name := splitFullName(fullName)

	// Check top-level objects.
	for _, f := range s.GetFile() {
		if f.GetPackage() == pkg {
			if i := FindServiceForFile(f, name); i != -1 {
				return f, f.Service[i], []int{FileDescriptorProtoServiceTag, i}
			}
			if i := FindMessageForFile(f, name); i != -1 {
				return f, f.MessageType[i], []int{FileDescriptorProtoMessageTag, i}
			}
			if i := FindEnumForFile(f, name); i != -1 {
				return f, f.EnumType[i], []int{FileDescriptorProtoEnumTag, i}
			}
		}
	}

	// Recurse.
	var parent interface{}
	file, parent, path = Resolve(s, pkg)
	switch parent := parent.(type) {

	case *pb.ServiceDescriptorProto:
		if i := FindMethodForService(parent, name); i != -1 {
			return file, parent.Method[i], append(path, ServiceDescriptorProtoMethodTag, i)
		}

	case *pb.DescriptorProto:
		if i := FindMessage(parent, name); i != -1 {
			return file, parent.NestedType[i], append(path, DescriptorProtoNestedTypeTag, i)
		}
		if i := FindEnum(parent, name); i != -1 {
			return file, parent.EnumType[i], append(path, DescriptorProtoEnumTypeTag, i)
		}
		if i := FindField(parent, name); i != -1 {
			return file, parent.Field[i], append(path, DescriptorProtoFieldTag, i)
		}
		if i := FindOneOf(parent, name); i != -1 {
			return file, parent.OneofDecl[i], append(path, DescriptorProtoOneOfTag, i)
		}

	case *pb.EnumDescriptorProto:
		if i := FindEnumValue(parent, name); i != -1 {
			return file, parent.Value[i], append(path, EnumDescriptorProtoValueTag, i)
		}
	}

	return nil, nil, nil
}

// FindService searches for a service by full name.
func FindService(s *pb.FileDescriptorSet, fullName string) (file *pb.FileDescriptorProto, serviceIndex int) {
	pkg, name := splitFullName(fullName)
	for _, f := range s.GetFile() {
		if f.GetPackage() == pkg {
			if i := FindServiceForFile(f, name); i != -1 {
				return f, i
			}
		}
	}
	return nil, -1
}

////////////////////////////////////////////////////////////////////////////////
// FileDescriptorProto

// FindServiceForFile searches for a FileDescriptorProto by name.
func FindServiceForFile(f *pb.FileDescriptorProto, name string) int {
	for i, x := range f.GetService() {
		if x.GetName() == name {
			return i
		}
	}
	return -1
}

// FindMessageForFile searches for a DescriptorProto by name.
func FindMessageForFile(f *pb.FileDescriptorProto, name string) int {
	for i, x := range f.GetMessageType() {
		if x.GetName() == name {
			return i
		}
	}
	return -1
}

// FindEnumForFile searches for an EnumDescriptorProto by name.
func FindEnumForFile(f *pb.FileDescriptorProto, name string) int {
	for i, x := range f.GetEnumType() {
		if x.GetName() == name {
			return i
		}
	}
	return -1
}

////////////////////////////////////////////////////////////////////////////////
// ServiceDescriptorProto

// FindMethodForService searches for a MethodDescriptorProto by name.
func FindMethodForService(s *pb.ServiceDescriptorProto, name string) int {
	for i, x := range s.GetMethod() {
		if x.GetName() == name {
			return i
		}
	}
	return -1
}

////////////////////////////////////////////////////////////////////////////////
// DescriptorProto (a message)

// FindField searches for a FieldDescriptorProto by name.
func FindField(d *pb.DescriptorProto, name string) int {
	for i, x := range d.GetField() {
		if x.GetName() == name {
			return i
		}
	}
	return -1
}

// FindMessage searches for a nested DescriptorProto by name.
func FindMessage(d *pb.DescriptorProto, name string) int {
	for i, x := range d.GetNestedType() {
		if x.GetName() == name {
			return i
		}
	}
	return -1
}

// FindEnum searches for a nested EnumDescriptorProto by name.
func FindEnum(d *pb.DescriptorProto, name string) int {
	for i, x := range d.GetEnumType() {
		if x.GetName() == name {
			return i
		}
	}
	return -1
}

// FindOneOf searches for a nested OneofDescriptorProto by name.
func FindOneOf(d *pb.DescriptorProto, name string) int {
	for i, x := range d.GetOneofDecl() {
		if x.GetName() == name {
			return i
		}
	}
	return -1
}

////////////////////////////////////////////////////////////////////////////////
// FieldDescriptorProto

// Repeated returns true if the field is repeated.
func Repeated(f *pb.FieldDescriptorProto) bool {
	return f.GetLabel() == pb.FieldDescriptorProto_LABEL_REPEATED
}

////////////////////////////////////////////////////////////////////////////////
// EnumDescriptorProto

// FindEnumValue searches for an EnumValueDescriptorProto by name.
func FindEnumValue(e *pb.EnumDescriptorProto, name string) int {
	for i, x := range e.GetValue() {
		if x.GetName() == name {
			return i
		}
	}
	return -1
}

// FindValueByNumber searches for an EnumValueDescriptorProto by number.
func FindValueByNumber(e *pb.EnumDescriptorProto, number int32) int {
	for i, x := range e.GetValue() {
		if x.GetNumber() == number {
			return i
		}
	}
	return -1
}

////////////////////////////////////////////////////////////////////////////////
// SourceCodeInfo

// At returns a pointer to a descriptor proto or its field at the given path.
// The path has same semantics as
// descriptor.SourceCodeInfo_Location.Path. See its comment for explanation.
//
// For example, given a FileDescriptorProto and path [4, 2],
// At will return a pointer to the 2nd top-level message DescriptorProto
// because 4 is FileDescriptorProto.MessageType field tag.
func At(descProto proto.Message, path []int32) (interface{}, error) {
	// invariant: cur's value must implement proto.Message.
	cur := reflect.ValueOf(descProto)
	for len(path) > 0 {
		tag := int(path[0])
		path = path[1:]

		// The tag->field index mapping could be precomputed.
		var prop *proto.Properties
		for _, p := range proto.GetProperties(cur.Type().Elem()).Prop {
			if p.Tag == tag {
				prop = p
				break
			}
		}
		if prop == nil {
			return nil, fmt.Errorf("%T has no tag %d", cur.Interface(), tag)
		}
		next := cur.Elem().FieldByName(prop.Name)

		if next.Kind() == reflect.Slice {
			switch {
			case len(path) == 0:
				return nil, fmt.Errorf("expected element index in path")
			case path[0] < 0, int(path[0]) >= next.Len():
				return nil, fmt.Errorf("element index out of bounds")
			}
			next, path = next.Index(int(path[0])), path[1:]
			if next.Kind() != reflect.Ptr {
				panic("impossible: all repeated fields are slices of pointers")
			}
		}

		if _, ok := next.Interface().(proto.Message); !ok {
			// next is not a descriptor proto, it must be a field,
			// e.g. myFieldDescriptorProto.Name.
			if len(path) > 0 {
				return nil, fmt.Errorf("expected path end, but got %q", path)
			}
			// return the address of the field, not field value.
			return next.Addr().Interface(), nil
		}
		cur = next
	}
	return cur.Interface(), nil
}

// IndexSourceCodeInfo returns a map that maps a pointer to the associated
// source code info, where the pointer points to a descriptor proto or its
// field, e.g. &myFieldDescriptorProto.Name.
//
// IndexSourceCodeInfo can be used to retrieve comments.
func IndexSourceCodeInfo(f *pb.FileDescriptorProto) (map[interface{}]*pb.SourceCodeInfo_Location, error) {
	if f.SourceCodeInfo == nil {
		return nil, nil
	}
	ret := make(map[interface{}]*pb.SourceCodeInfo_Location, len(f.SourceCodeInfo.Location))
	for _, loc := range f.SourceCodeInfo.Location {
		ptr, err := At(f, loc.Path)
		if err != nil {
			return nil, err
		}
		ret[ptr] = loc
	}
	return ret, nil
}
