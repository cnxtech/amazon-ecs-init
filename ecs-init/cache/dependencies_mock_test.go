// Copyright 2015-2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.
//
// Source: dependencies.go in package cache
// Automatically generated by MockGen. DO NOT EDIT!

package cache

import (
	io "io"
	http "net/http"
	os "os"

	gomock "github.com/golang/mock/gomock"
)

// Mock of httpGetter interface
type MockhttpGetter struct {
	ctrl     *gomock.Controller
	recorder *_MockhttpGetterRecorder
}

// Recorder for MockhttpGetter (not exported)
type _MockhttpGetterRecorder struct {
	mock *MockhttpGetter
}

func NewMockhttpGetter(ctrl *gomock.Controller) *MockhttpGetter {
	mock := &MockhttpGetter{ctrl: ctrl}
	mock.recorder = &_MockhttpGetterRecorder{mock}
	return mock
}

func (_m *MockhttpGetter) EXPECT() *_MockhttpGetterRecorder {
	return _m.recorder
}

func (_m *MockhttpGetter) Get(url string) (*http.Response, error) {
	ret := _m.ctrl.Call(_m, "Get", url)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockhttpGetterRecorder) Get(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0)
}

// Mock of fileSystem interface
type MockfileSystem struct {
	ctrl     *gomock.Controller
	recorder *_MockfileSystemRecorder
}

// Recorder for MockfileSystem (not exported)
type _MockfileSystemRecorder struct {
	mock *MockfileSystem
}

func NewMockfileSystem(ctrl *gomock.Controller) *MockfileSystem {
	mock := &MockfileSystem{ctrl: ctrl}
	mock.recorder = &_MockfileSystemRecorder{mock}
	return mock
}

func (_m *MockfileSystem) EXPECT() *_MockfileSystemRecorder {
	return _m.recorder
}

func (_m *MockfileSystem) MkdirAll(path string, perm os.FileMode) error {
	ret := _m.ctrl.Call(_m, "MkdirAll", path, perm)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockfileSystemRecorder) MkdirAll(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MkdirAll", arg0, arg1)
}

func (_m *MockfileSystem) TempFile(dir string, prefix string) (*os.File, error) {
	ret := _m.ctrl.Call(_m, "TempFile", dir, prefix)
	ret0, _ := ret[0].(*os.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockfileSystemRecorder) TempFile(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TempFile", arg0, arg1)
}

func (_m *MockfileSystem) Remove(path string) {
	_m.ctrl.Call(_m, "Remove", path)
}

func (_mr *_MockfileSystemRecorder) Remove(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Remove", arg0)
}

func (_m *MockfileSystem) TeeReader(r io.Reader, w io.Writer) io.Reader {
	ret := _m.ctrl.Call(_m, "TeeReader", r, w)
	ret0, _ := ret[0].(io.Reader)
	return ret0
}

func (_mr *_MockfileSystemRecorder) TeeReader(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TeeReader", arg0, arg1)
}

func (_m *MockfileSystem) Copy(dst io.Writer, src io.Reader) (int64, error) {
	ret := _m.ctrl.Call(_m, "Copy", dst, src)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockfileSystemRecorder) Copy(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Copy", arg0, arg1)
}

func (_m *MockfileSystem) Rename(oldpath string, newpath string) error {
	ret := _m.ctrl.Call(_m, "Rename", oldpath, newpath)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockfileSystemRecorder) Rename(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Rename", arg0, arg1)
}

func (_m *MockfileSystem) ReadAll(r io.Reader) ([]byte, error) {
	ret := _m.ctrl.Call(_m, "ReadAll", r)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockfileSystemRecorder) ReadAll(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReadAll", arg0)
}

func (_m *MockfileSystem) Open(name string) (io.ReadCloser, error) {
	ret := _m.ctrl.Call(_m, "Open", name)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockfileSystemRecorder) Open(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Open", arg0)
}

func (_m *MockfileSystem) Stat(name string) (fileSizeInfo, error) {
	ret := _m.ctrl.Call(_m, "Stat", name)
	ret0, _ := ret[0].(fileSizeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockfileSystemRecorder) Stat(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stat", arg0)
}

func (_m *MockfileSystem) Base(path string) string {
	ret := _m.ctrl.Call(_m, "Base", path)
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockfileSystemRecorder) Base(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Base", arg0)
}

func (_m *MockfileSystem) WriteFile(filename string, data []byte, perm os.FileMode) error {
	ret := _m.ctrl.Call(_m, "WriteFile", filename, data, perm)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockfileSystemRecorder) WriteFile(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WriteFile", arg0, arg1, arg2)
}

// Mock of fileSizeInfo interface
type MockfileSizeInfo struct {
	ctrl     *gomock.Controller
	recorder *_MockfileSizeInfoRecorder
}

// Recorder for MockfileSizeInfo (not exported)
type _MockfileSizeInfoRecorder struct {
	mock *MockfileSizeInfo
}

func NewMockfileSizeInfo(ctrl *gomock.Controller) *MockfileSizeInfo {
	mock := &MockfileSizeInfo{ctrl: ctrl}
	mock.recorder = &_MockfileSizeInfoRecorder{mock}
	return mock
}

func (_m *MockfileSizeInfo) EXPECT() *_MockfileSizeInfoRecorder {
	return _m.recorder
}

func (_m *MockfileSizeInfo) Size() int64 {
	ret := _m.ctrl.Call(_m, "Size")
	ret0, _ := ret[0].(int64)
	return ret0
}

func (_mr *_MockfileSizeInfoRecorder) Size() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Size")
}
