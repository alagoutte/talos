// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Code generated by "deep-copy -type SeccompProfileSpec -header-file ../../../../hack/boilerplate.txt -o deep_copy.generated.go ."; DO NOT EDIT.

package cri

// DeepCopy generates a deep copy of SeccompProfileSpec.
func (o SeccompProfileSpec) DeepCopy() SeccompProfileSpec {
	var cp SeccompProfileSpec = o
	if o.Value != nil {
		cp.Value = make(map[string]interface{}, len(o.Value))
		for k2, v2 := range o.Value {
			cp.Value[k2] = v2
		}
	}
	return cp
}