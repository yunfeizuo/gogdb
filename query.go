package main

import (
	"fmt"
	"reflect"
)

type QueryEdge struct {
	Source   QueryVertex
	Target   QueryVertex
	Name     string
	Type     reflect.Type
	Resolver func(id int64, edge QueryEdge, args []interface{}) (interface{}, error)
}

type QueryVertex struct {
	Id    int64
	Edges []QueryEdge
	Kind  int
}

const (
	Scalar = 1
	Object = 2
	Array  = 3
)

// ResolveEdge Resolve a query edge recursively. The values will be retrieved from
// storage. Depends on the type of the edge value, the return value could be a scalar,
// an object or an array (cannot be map or anything else!).
func ResolveEdge(id int64, edge QueryEdge) reflect.Value {
	redis := Store{}
	key := fmt.Sprintf("%v.%v", id, edge.Name)

	if edge.Target.Kind == Scalar {
		v := redis.Get(key)
		return reflect.ValueOf(v)
	}

	if edge.Target.Kind == Object {
		id := redis.Get(key).(int64)
		objVal := reflect.New(edge.Type)
		for _, e := range edge.Target.Edges {
			var m reflect.Value
			if e.Resolver != nil {
				// TODO: handle error
				v, _ := e.Resolver(id, e, nil)
				m = reflect.ValueOf(v)
			} else {
				m = ResolveEdge(id, e)
			}

			// setField(&objVal, e.Name, &m)
			s := objVal.Elem()
			if s.Kind() != reflect.Struct {
				panic("Not sturct")
			}
			f := s.FieldByName(e.Name)
			if !f.IsValid() {
				panic("Field not found")
			}
			if !f.CanSet() {
				panic("Field can not set")
			}
			// TODO: check type
			// if f.Kind() == reflect.Int
			// v := reflect.ValueOf(m)

			f.Set(m)
		}
		return objVal
	}

	if edge.Target.Kind == Array {
		return reflect.Value{}
	}
	return reflect.Value{}
}

func setField(objVal *reflect.Value, name string, value *reflect.Value) {
	s := objVal.Elem()
	if s.Kind() != reflect.Struct {
		panic("Not sturct")
	}
	f := s.FieldByName(name)
	if !f.IsValid() {
		panic("Field not found")
	}
	if !f.CanSet() {
		panic("Field can not set")
	}
	// TODO: check type
	// if f.Kind() == reflect.Int
	v := reflect.ValueOf(value)

	fmt.Println(objVal, s, f, name, v)
	f.Set(v)
}
