// Copyright 2017 Google Inc.
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

package simple

import (
	"testing"

	"github.com/golang/protobuf/proto"
	gpb "github.com/openconfig/gnoi"
	bgppb "github.com/openconfig/gnoi/bgp"
)

func TestGNOI(t *testing.T) {
	tests := []struct {
		desc string
		in   proto.Message
		want string
	}{{
		desc: "gpb.Path",
		in: &gpb.Path{
			Origin: "oc",
			Elem:   []*gpb.PathElem{{Name: "interfaces", Key: map[string]string{"name": "Ethernet1/1/0"}}},
		},
		want: "origin: \"oc\"\nelem: <\n  name: \"interfaces\"\n  key: <\n    key: \"name\"\n    value: \"Ethernet1/1/0\"\n  >\n>\n",
	}, {
		desc: "gpb.HashType",
		in: &gpb.HashType{
			Method: gpb.HashType_MD5,
			Hash:   []byte("foo"),
		},
		want: "method: MD5\nhash: \"foo\"\n",
	}, {
		desc: "bgp.ClearBGPNeighborRequest",
		in: &bgppb.ClearBGPNeighborRequest{
			Address:         "foo",
			RoutingInstance: "bar",
			Mode:            bgppb.ClearBGPNeighborRequest_HARD,
		},
		want: "address: \"foo\"\nrouting_instance: \"bar\"\nmode: HARD\n",
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if got := proto.MarshalTextString(tt.in); got != tt.want {
				t.Fatalf("Text Proto output failed: got %q, want %q", got, tt.want)
			}
		})
	}
}
