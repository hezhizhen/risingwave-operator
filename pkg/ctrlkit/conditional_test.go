/*
 * Copyright 2022 Singularity Data
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ctrlkit

import "testing"

func Test_If(t *testing.T) {
	var action = NewAction("", nil)

	testcases := map[string]struct {
		condition bool
		action    Action
		expect    Action
	}{
		"if-false-then-nop": {
			condition: false,
			action:    action,
			expect:    Nop,
		},
		"if-true-then-self": {
			condition: true,
			action:    action,
			expect:    action,
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			if tc.expect != If(tc.condition, tc.action) {
				t.Fail()
			}
		})
	}
}
