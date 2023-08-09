// (C) 2016-2023 Ant Group Co.,Ltd.
// SPDX-License-Identifier: Apache-2.0

package sha256

import "crypto/sha256"

// ComputeHash computes hash of the given raw message
func ComputeHash(raw []byte) ([]byte, error) {
	hash := sha256.New()
	hash.Write(raw)
	return hash.Sum(nil), nil
}
