#!/bin/bash

# (C) 2016-2023 Ant Group Co.,Ltd.
# SPDX-License-Identifier: Apache-2.0

rm -rf crypto
rm crypto.tar.gz
mkdir crypto
cp ../../src/crypto/cmd/bls/tbls_sk* ./crypto/
tar -czf crypto.tar.gz ./crypto
