/* -------------------------------------------------------------------------- */
/* Copyright 2002-2019, OpenNebula Project, OpenNebula Systems                */
/*                                                                            */
/* Licensed under the Apache License, Version 2.0 (the "License"); you may    */
/* not use this file except in compliance with the License. You may obtain    */
/* a copy of the License at                                                   */
/*                                                                            */
/* http://www.apache.org/licenses/LICENSE-2.0                                 */
/*                                                                            */
/* Unless required by applicable law or agreed to in writing, software        */
/* distributed under the License is distributed on an "AS IS" BASIS,          */
/* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.   */
/* See the License for the specific language governing permissions and        */
/* limitations under the License.                                             */
/*--------------------------------------------------------------------------- */

package hook

import (
	dyn "github.com/OpenNebula/one/src/oca/go/src/goca/dynamic"
	"github.com/OpenNebula/one/src/oca/go/src/goca/schemas/hook/keys"
)

// Template is a hook template
type Template struct {
	dyn.Template
}

// NewTemplate returns a hook Template structure
func NewTemplate() *Template {
	return &Template{}
}

// Get returns the string value for an hook template keys
func (n *Template) Get(key keys.Template) (string, error) {
	return n.GetStr(string(key))
}

// GetI returns the integer value for a hook template key
func (n *Template) GetI(key keys.Template) (int, error) {
	return n.GetInt(string(key))
}

// Add adds an hook template key, value pair
func (n *Template) Add(key keys.Template, value string) {
	n.AddPair(string(key), value)
}
