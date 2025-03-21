// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ebpf

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInstructionChainHelperTest(t *testing.T) {
	tests := []struct {
		testName      string
		operations    []Instruction
		expectedError error
	}{
		{
			testName:      "Instruction chain no jumps",
			operations:    []Instruction{Mov64(RegR0, 0), Mul64(RegR0, 10), Mov64(RegR0, RegR1), Exit()},
			expectedError: nil,
		},
		{
			testName:      "Instruction chain with jumps",
			operations:    []Instruction{Mov64(RegR0, 0), JmpGT(RegR0, 0, 4), Mul64(RegR0, 10), JmpLT(RegR0, RegR1, 2), Jmp(1), Mov64(RegR0, RegR1), Exit()},
			expectedError: nil,
		},
		{
			testName:      "Invalid instruction returns error",
			operations:    []Instruction{Mov64(RegR0, uint32(0))},
			expectedError: fmt.Errorf("Nil instruction at index 0, did you pass an unsigned int value?"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			t.Logf("Running test case %s", tc.testName)
			root, err := InstructionSequence(tc.operations...)
			if tc.expectedError != nil {
				if err.Error() != tc.expectedError.Error() {
					t.Fatalf("Want error %v, got %v", tc.expectedError, err)
				}
				return
			}

			if !reflect.DeepEqual(root, tc.operations) {
				t.Errorf("Want instruction array = %v, have %v", tc.operations, root)
			}
		})
	}
}
