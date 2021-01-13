// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package container

// JSONSerializer provides JSON serialization
type JSONSerializer interface {
	// ToJSON outputs the JSON representation of container's elements.
	ToJSON() ([]byte, error)
}

// JSONDeserializer provides JSON deserialization
type JSONDeserializer interface {
	// FromJSON populates container's elements from the input JSON representation.
	FromJSON([]byte) error
}
