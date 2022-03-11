/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright 2022 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package common

import (
	"math/big"
	"unsafe"
)

type MemoryUsage struct {
	Kind   MemoryKind
	Amount uint64
}

type MemoryGauge interface {
	UseMemory(usage MemoryUsage)
}

func NewStringMemoryUsage(length int) MemoryUsage {
	return MemoryUsage{
		Kind:   MemoryKindString,
		Amount: uint64(length) + 1, // +1 to account for empty strings
	}
}

func NewBigIntMemoryUsage(bytes int) MemoryUsage {
	return MemoryUsage{
		Kind:   MemoryKindBigInt,
		Amount: uint64(bytes),
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

const bigIntWordSize = int(unsafe.Sizeof(big.Word(0)))

func BigIntByteLength(v *big.Int) int {
	// NOTE: big.Int.Bits() actually returns bytes:
	// []big.Word, where big.Word = uint
	return len(v.Bits()) * bigIntWordSize
}

func NewPlusBigIntMemoryUsage(a, b *big.Int) MemoryUsage {
	return NewBigIntMemoryUsage(
		max(
			BigIntByteLength(a),
			BigIntByteLength(b),
		) + bigIntWordSize,
	)
}

func NewMulBigIntMemoryUsage(a, b *big.Int) MemoryUsage {
	return NewBigIntMemoryUsage(
		BigIntByteLength(a) +
			BigIntByteLength(b),
	)
}