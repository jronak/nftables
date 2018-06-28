// Copyright 2018 Google LLC. All Rights Reserved.
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

package expr

import (
	"fmt"

	"github.com/google/nftables/binaryutil"
	"github.com/mdlayher/netlink"
	"golang.org/x/sys/unix"
)

type Immediate struct {
	Register uint32
	Data     []byte
}

func (e *Immediate) marshal() ([]byte, error) {
	immData, err := netlink.MarshalAttributes([]netlink.Attribute{
		{Type: unix.NFTA_DATA_VALUE, Data: e.Data},
	})
	if err != nil {
		return nil, err
	}

	data, err := netlink.MarshalAttributes([]netlink.Attribute{
		{Type: unix.NFTA_IMMEDIATE_DREG, Data: binaryutil.BigEndian.PutUint32(e.Register)},
		{Type: unix.NLA_F_NESTED | unix.NFTA_IMMEDIATE_DATA, Data: immData},
	})
	if err != nil {
		return nil, err
	}
	return netlink.MarshalAttributes([]netlink.Attribute{
		{Type: unix.NFTA_EXPR_NAME, Data: []byte("immediate\x00")},
		{Type: unix.NLA_F_NESTED | unix.NFTA_EXPR_DATA, Data: data},
	})
}

func (e *Immediate) unmarshal(data []byte) error {
	return fmt.Errorf("not yet implemented")
}