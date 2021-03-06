// Code generated by mockery v1.0.0. DO NOT EDIT.

package fs_mocks

import http "net/http"
import mock "github.com/stretchr/testify/mock"
import os "os"

// FileReader is an autogenerated mock type for the FileReader type
type FileReader struct {
	mock.Mock
}

// Open provides a mock function with given fields: name
func (_m *FileReader) Open(name string) (http.File, error) {
	ret := _m.Called(name)

	var r0 http.File
	if rf, ok := ret.Get(0).(func(string) http.File); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Read provides a mock function with given fields: path
func (_m *FileReader) Read(path string) ([]byte, error) {
	ret := _m.Called(path)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Stat provides a mock function with given fields: path
func (_m *FileReader) Stat(path string) (os.FileInfo, error) {
	ret := _m.Called(path)

	var r0 os.FileInfo
	if rf, ok := ret.Get(0).(func(string) os.FileInfo); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(os.FileInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
