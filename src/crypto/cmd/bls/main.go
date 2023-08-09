// (C) 2016-2023 Ant Group Co.,Ltd.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"fmt"

	"bkr-go/crypto/bls"
)

func main() {
	n := flag.Int("n", 4, "number of nodes")
	th := flag.Int("t", 2, "number of shares")
	flag.Parse()
	// Crypto setup
	err := bls.GenerateBlsKey("./", *n, *th)
	if err != nil {
		fmt.Println(err)
	}
	_, _, err = bls.LoadBlsKey("./", *n, *th)
	if err != nil {
		fmt.Println(err)
	}
}
