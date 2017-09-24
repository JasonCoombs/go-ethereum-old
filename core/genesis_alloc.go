// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package core

// Constants containing the genesis allocation of built-in genesis blocks.
// Their content is an RLP-encoded list of (address, balance) tuples.
// Use mkalloc.go to create/update them.

const mainnetAllocData = "\xf8@\u07d4\xd6\xf4\xa7\xb1/\xf9\x11\xcfZ\x93>\xe5V>\x8f\xab\u007f!\v\xa9\x8965\u026d\xc5\u07a0\x00\x00\u07d4\xec\x16\xdc&\xcb\u0131\x82\x81\xeb\xfa\xd1?/\xb8\xecQu\xd6\u00c965\u026d\xc5\u07a0\x00\x00"
const testnetAllocData = ""
const rinkebyAllocData = ""
const devAllocData = ""
