package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
)

type SchemaObjectMap []SchemaObjectMapItem

type SchemaObjectMapItem struct {
	Key   string
	Value SchemaObject
}

func (ms SchemaObjectMap) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	buf.Write([]byte{'{'})
	for i, mi := range ms {
		b, err := json.Marshal(&mi.Value)
		if err != nil {
			return nil, err
		}
		buf.WriteString(fmt.Sprintf("%q:", fmt.Sprintf("%v", mi.Key)))
		buf.Write(b)
		if i < len(ms)-1 {
			buf.Write([]byte{','})
		}
	}
	buf.Write([]byte{'}'})
	return buf.Bytes(), nil
}

func (ms *SchemaObjectMap) UnmarshalJSON(data []byte) error {

	var v map[string]SchemaObject
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	indexes := make([]int, 0)
	indexMap := map[int]string{}
	for key := range v {
		index := bytes.Index(data, []byte(`"`+key+`"`))
		indexes = append(indexes, index)
		indexMap[index] = key
	}

	if len(v) != len(indexMap) {
		panic("unable to unmarshal ordered map")
	}

	keys := make([]string, 0)
	sort.Ints(indexes)
	for _, index := range indexes {
		keys = append(keys, indexMap[index])
	}

	items := make([]SchemaObjectMapItem, 0, len(v))
	for _, key := range keys {
		items = append(items, SchemaObjectMapItem{
			Key:   key,
			Value: v[key],
		})
	}

	*ms = items

	return nil
}

type PathItemObjectMap []PathItemObjectMapItem

type PathItemObjectMapItem struct {
	Key   string
	Value PathItemObject
}

func (ms PathItemObjectMap) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	buf.Write([]byte{'{'})
	for i, mi := range ms {
		b, err := json.Marshal(&mi.Value)
		if err != nil {
			return nil, err
		}
		buf.WriteString(fmt.Sprintf("%q:", fmt.Sprintf("%v", mi.Key)))
		buf.Write(b)
		if i < len(ms)-1 {
			buf.Write([]byte{','})
		}
	}
	buf.Write([]byte{'}'})
	return buf.Bytes(), nil
}

func (ms *PathItemObjectMap) UnmarshalJSON(data []byte) error {

	var v map[string]PathItemObject
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	indexes := make([]int, 0)
	indexMap := map[int]string{}
	for key := range v {
		index := bytes.Index(data, []byte(`"`+key+`"`))
		indexes = append(indexes, index)
		indexMap[index] = key
	}

	if len(v) != len(indexMap) {
		panic("unable to unmarshal ordered map")
	}

	keys := make([]string, 0)
	sort.Ints(indexes)
	for _, index := range indexes {
		keys = append(keys, indexMap[index])
	}

	items := make([]PathItemObjectMapItem, 0, len(v))
	for _, key := range keys {
		items = append(items, PathItemObjectMapItem{
			Key:   key,
			Value: v[key],
		})
	}

	*ms = items

	return nil
}

type ResponseObjectMap []ResponseObjectMapItem

type ResponseObjectMapItem struct {
	Key   string
	Value ResponseObject
}

func (ms ResponseObjectMap) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	buf.Write([]byte{'{'})
	for i, mi := range ms {
		b, err := json.Marshal(&mi.Value)
		if err != nil {
			return nil, err
		}
		buf.WriteString(fmt.Sprintf("%q:", fmt.Sprintf("%v", mi.Key)))
		buf.Write(b)
		if i < len(ms)-1 {
			buf.Write([]byte{','})
		}
	}
	buf.Write([]byte{'}'})
	return buf.Bytes(), nil
}

func (ms *ResponseObjectMap) UnmarshalJSON(data []byte) error {

	var v map[string]ResponseObject
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	indexes := make([]int, 0)
	indexMap := map[int]string{}
	for key := range v {
		index := bytes.Index(data, []byte(`"`+key+`"`))
		indexes = append(indexes, index)
		indexMap[index] = key
	}

	if len(v) != len(indexMap) {
		panic("unable to unmarshal ordered map")
	}

	keys := make([]string, 0)
	sort.Ints(indexes)
	for _, index := range indexes {
		keys = append(keys, indexMap[index])
	}

	items := make([]ResponseObjectMapItem, 0, len(v))
	for _, key := range keys {
		items = append(items, ResponseObjectMapItem{
			Key:   key,
			Value: v[key],
		})
	}

	*ms = items

	return nil
}

type HeaderObjectMap []HeaderObjectMapItem

type HeaderObjectMapItem struct {
	Key   string
	Value HeaderObject
}

func (ms HeaderObjectMap) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	buf.Write([]byte{'{'})
	for i, mi := range ms {
		b, err := json.Marshal(&mi.Value)
		if err != nil {
			return nil, err
		}
		buf.WriteString(fmt.Sprintf("%q:", fmt.Sprintf("%v", mi.Key)))
		buf.Write(b)
		if i < len(ms)-1 {
			buf.Write([]byte{','})
		}
	}
	buf.Write([]byte{'}'})
	return buf.Bytes(), nil
}

func (ms *HeaderObjectMap) UnmarshalJSON(data []byte) error {

	var v map[string]HeaderObject
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	indexes := make([]int, 0)
	indexMap := map[int]string{}
	for key := range v {
		index := bytes.Index(data, []byte(`"`+key+`"`))
		indexes = append(indexes, index)
		indexMap[index] = key
	}

	if len(v) != len(indexMap) {
		panic("unable to unmarshal ordered map")
	}

	keys := make([]string, 0)
	sort.Ints(indexes)
	for _, index := range indexes {
		keys = append(keys, indexMap[index])
	}

	items := make([]HeaderObjectMapItem, 0, len(v))
	for _, key := range keys {
		items = append(items, HeaderObjectMapItem{
			Key:   key,
			Value: v[key],
		})
	}

	*ms = items

	return nil
}

type MediaTypeObjectMap []MediaTypeObjectMapItem

type MediaTypeObjectMapItem struct {
	Key   string
	Value MediaTypeObject
}

func (ms MediaTypeObjectMap) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	buf.Write([]byte{'{'})
	for i, mi := range ms {
		b, err := json.Marshal(&mi.Value)
		if err != nil {
			return nil, err
		}
		buf.WriteString(fmt.Sprintf("%q:", fmt.Sprintf("%v", mi.Key)))
		buf.Write(b)
		if i < len(ms)-1 {
			buf.Write([]byte{','})
		}
	}
	buf.Write([]byte{'}'})
	return buf.Bytes(), nil
}

func (ms *MediaTypeObjectMap) UnmarshalJSON(data []byte) error {

	var v map[string]MediaTypeObject
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	indexes := make([]int, 0)
	indexMap := map[int]string{}
	for key := range v {
		index := bytes.Index(data, []byte(`"`+key+`"`))
		indexes = append(indexes, index)
		indexMap[index] = key
	}

	if len(v) != len(indexMap) {
		panic("unable to unmarshal ordered map")
	}

	keys := make([]string, 0)
	sort.Ints(indexes)
	for _, index := range indexes {
		keys = append(keys, indexMap[index])
	}

	items := make([]MediaTypeObjectMapItem, 0, len(v))
	for _, key := range keys {
		items = append(items, MediaTypeObjectMapItem{
			Key:   key,
			Value: v[key],
		})
	}

	*ms = items

	return nil
}

type SecuritySchemeObjectMap []SecuritySchemeObjectMapItem

type SecuritySchemeObjectMapItem struct {
	Key   string
	Value SecuritySchemeObject
}

func (ms SecuritySchemeObjectMap) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	buf.Write([]byte{'{'})
	for i, mi := range ms {
		b, err := json.Marshal(&mi.Value)
		if err != nil {
			return nil, err
		}
		buf.WriteString(fmt.Sprintf("%q:", fmt.Sprintf("%v", mi.Key)))
		buf.Write(b)
		if i < len(ms)-1 {
			buf.Write([]byte{','})
		}
	}
	buf.Write([]byte{'}'})
	return buf.Bytes(), nil
}

func (ms *SecuritySchemeObjectMap) UnmarshalJSON(data []byte) error {

	var v map[string]SecuritySchemeObject
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	indexes := make([]int, 0)
	indexMap := map[int]string{}
	for key := range v {
		index := bytes.Index(data, []byte(`"`+key+`"`))
		indexes = append(indexes, index)
		indexMap[index] = key
	}

	if len(v) != len(indexMap) {
		panic("unable to unmarshal ordered map")
	}

	keys := make([]string, 0)
	sort.Ints(indexes)
	for _, index := range indexes {
		keys = append(keys, indexMap[index])
	}

	items := make([]SecuritySchemeObjectMapItem, 0, len(v))
	for key, value := range v {
		items = append(items, SecuritySchemeObjectMapItem{
			Key:   key,
			Value: value,
		})
	}

	*ms = items

	return nil
}

type ServerVariableObjectMap []ServerVariableObjectMapItem

type ServerVariableObjectMapItem struct {
	Key   string
	Value ServerVariableObject
}

func (ms ServerVariableObjectMap) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	buf.Write([]byte{'{'})
	for i, mi := range ms {
		b, err := json.Marshal(&mi.Value)
		if err != nil {
			return nil, err
		}
		buf.WriteString(fmt.Sprintf("%q:", fmt.Sprintf("%v", mi.Key)))
		buf.Write(b)
		if i < len(ms)-1 {
			buf.Write([]byte{','})
		}
	}
	buf.Write([]byte{'}'})
	return buf.Bytes(), nil
}

func (ms *ServerVariableObjectMap) UnmarshalJSON(data []byte) error {

	var v map[string]ServerVariableObject
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	indexes := make([]int, 0)
	indexMap := map[int]string{}
	for key := range v {
		index := bytes.Index(data, []byte(`"`+key+`"`))
		indexes = append(indexes, index)
		indexMap[index] = key
	}

	if len(v) != len(indexMap) {
		panic("unable to unmarshal ordered map")
	}

	keys := make([]string, 0)
	sort.Ints(indexes)
	for _, index := range indexes {
		keys = append(keys, indexMap[index])
	}

	items := make([]ServerVariableObjectMapItem, 0, len(v))
	for key, value := range v {
		items = append(items, ServerVariableObjectMapItem{
			Key:   key,
			Value: value,
		})
	}

	*ms = items

	return nil
}
