// Copyright 2014 Google Inc. All Rights Reserved.
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

package refutil

import (
	"testing"
)

func Test_IsSlice(t *testing.T) {
	sliceInts := []int{}
	if !IsSlice(sliceInts) {
		t.Errorf("Int slice should be detected as a slice")
	}
	sliceIntsPtr := []*int{}
	if !IsSlice(sliceIntsPtr) {
		t.Errorf("Int ptr slice should be detected as a slice")
	}
	notSlice := "hello world"
	if IsSlice(notSlice) {
		t.Errorf("A string is not a slice")
	}
}

func Test_IsSlicePtr(t *testing.T) {
	sliceStrings := []string{}
	if !IsSlicePtr(&sliceStrings) {
		t.Errorf("A string slice ptr should be detected as a slice ptr")
	}
	if IsSlicePtr(sliceStrings) {
		t.Errorf("A string slice should not be detected as a slice ptr")
	}
}

func Test_IsStruct(t *testing.T) {
	src := struct{}{}
	if !IsStruct(src) {
		t.Errorf("A struct should be detected as a struct")
	}
	if IsStruct(&src) {
		t.Errorf("A ptr to struct should not be detected as a struct")
	}
}

func Test_IsStructPtr(t *testing.T) {
	src := struct{}{}
	if !IsStructPtr(&src) {
		t.Errorf("A struct ptr should be detected as a struct ptr")
	}
	if IsStructPtr(src) {
		t.Errorf("A struct should not be detected as a struct ptr")
	}
}
