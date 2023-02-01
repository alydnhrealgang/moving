package utils

import (
	"reflect"
)

func AnyError(array interface{}, predicate func(index int, value interface{}) error) (error, int, interface{}) {
	s := reflect.ValueOf(array)
	for i := 0; i < s.Len(); i++ {
		value := s.Index(i).Interface()
		if err := predicate(i, value); nil != err {
			return err, i, value
		}
	}

	return nil, -1, nil
}

func All(array interface{}, predicate interface{}) bool {
	predicateFunc := reflect.ValueOf(predicate)
	s := reflect.ValueOf(array)
	for i := 0; i < s.Len(); i++ {
		value := s.Index(i)
		if !predicateFunc.Call([]reflect.Value{value})[0].Bool() {
			return false
		}
	}

	return true
}

func EmptyArray(array interface{}) bool {
	return nil == array || reflect.ValueOf(array).Len() == 0
}

func Select(array interface{}, selector interface{}) interface{} {
	selectorFunc := reflect.ValueOf(selector)
	sliceType := reflect.SliceOf(selectorFunc.Type().Out(0))
	s := reflect.ValueOf(array)
	selectedArray := reflect.MakeSlice(sliceType, 0, s.Len())
	for i := 0; i < s.Len(); i++ {
		value := s.Index(i)
		selectedArray = reflect.Append(selectedArray, selectorFunc.Call([]reflect.Value{value})[0])
	}
	return selectedArray.Interface()
}

func SelectWithIndex(array interface{}, selector interface{}) interface{} {
	selectorFunc := reflect.ValueOf(selector)
	sliceType := reflect.SliceOf(selectorFunc.Type().Out(0))
	s := reflect.ValueOf(array)

	selectedArray := reflect.MakeSlice(sliceType, 0, s.Len())
	for i := 0; i < s.Len(); i++ {
		value := s.Index(i)
		selectedArray = reflect.Append(selectedArray, selectorFunc.Call([]reflect.Value{value, reflect.ValueOf(i)})[0])
	}
	return selectedArray.Interface()
}

func Cast(array interface{}, castReceiver interface{}) {
	sliceArray := reflect.ValueOf(array)
	castArray := reflect.ValueOf(castReceiver)
	castType := castArray.Type().Elem()
	for i := 0; i < sliceArray.Len(); i++ {
		castArray.Index(i).Set(sliceArray.Index(i).Convert(castType))
	}
}

func Where(array interface{}, predicate interface{}) interface{} {
	predicateFunc := reflect.ValueOf(predicate)
	s := reflect.ValueOf(array)
	matchedArray := reflect.MakeSlice(s.Type(), 0, s.Len())
	for i := 0; i < s.Len(); i++ {
		value := s.Index(i)
		if predicateFunc.Call([]reflect.Value{value})[0].Bool() {
			matchedArray = reflect.Append(matchedArray, value)
		}
	}
	return matchedArray.Interface()
}

func Any(array interface{}, predicate interface{}) bool {
	predicateFunc := reflect.ValueOf(predicate)
	s := reflect.ValueOf(array)
	for i := 0; i < s.Len(); i++ {
		value := s.Index(i)
		if predicateFunc.Call([]reflect.Value{value})[0].Bool() {
			return true
		}
	}

	return false
}

func FirstOrDefault(array interface{}, predicate interface{}) (interface{}, bool) {
	s := reflect.ValueOf(array)
	if predicate == nil {
		if s.Len() == 0 {
			return nil, false
		}
		return s.Index(0).Interface(), true
	}

	predicateFunc := reflect.ValueOf(predicate)
	for i := 0; i < s.Len(); i++ {
		value := s.Index(i)
		if predicateFunc.Call([]reflect.Value{value})[0].Bool() {
			return value.Interface(), true
		}
	}

	return nil, false
}
