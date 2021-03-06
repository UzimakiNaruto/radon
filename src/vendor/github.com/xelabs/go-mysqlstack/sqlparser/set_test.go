/*
Copyright 2017 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sqlparser

import "strings"
import "testing"

func TestSet(t *testing.T) {
	validSQL := []struct {
		input  string
		output string
	}{
		{
			input:  "SET @@session.s1= 'ON', @@session.s2='OFF'",
			output: "set @@session.s1 = 'ON', @@session.s2 = 'OFF'",
		},

		{
			input:  "SET @@session.radon_stream_fetching= 'OFF'",
			output: "set @@session.radon_stream_fetching = 'OFF'",
		},
		{
			input:  "SET radon_stream_fetching= false",
			output: "set radon_stream_fetching = false",
		},
		{
			input:  "SET SESSION wait_timeout = 2147483",
			output: "set wait_timeout = 2147483",
		},
		{
			input:  "SET NAMES utf8",
			output: "set names = 'utf8'",
		},
	}

	for _, exp := range validSQL {
		sql := strings.TrimSpace(exp.input)
		tree, err := Parse(sql)
		if err != nil {
			t.Errorf("input: %s, err: %v", sql, err)
			continue
		}

		// Walk.
		Walk(func(node SQLNode) (bool, error) {
			return true, nil
		}, tree)

		got := String(tree.(*Set))
		if exp.output != got {
			t.Errorf("want:\n%s\ngot:\n%s", exp.output, got)
		}
	}
}
