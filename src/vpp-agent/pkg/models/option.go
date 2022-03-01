//  Copyright (c) 2019 Cisco and/or its affiliates.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package models

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"go.ligato.io/vpp-agent/v3/proto/ligato/generic"
)

func (r *LocalRegistry) checkProtoOptions(x interface{}) *LocallyKnownModel {
	p, ok := x.(protoreflect.Message)
	if !ok {
		return nil
	}
	s := proto.GetExtension(p.Interface(), generic.E_Model)
	if spec, ok := s.(*generic.ModelSpec); ok {
		km, err := r.Register(x, ToSpec(spec))
		if err != nil {
			panic(err)
		}
		return km.(*LocallyKnownModel)
	}
	return nil
}
