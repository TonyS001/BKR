#!/bin/bash

# (C) 2016-2023 Ant Group Co.,Ltd.
# SPDX-License-Identifier: Apache-2.0

NUM=$1

for(( i = 0 ; i < NUM ; i++)); do
{
    name="log/server$i"
    mkdir $name
} 
done

wait
