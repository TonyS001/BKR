#!/bin/bash

# (C) 2016-2023 Ant Group Co.,Ltd.
# SPDX-License-Identifier: Apache-2.0

NUM=$1

for ((i = 0; i < NUM; i++)); do
{
        host1=$(jq ".nodes[$i].PublicIpAddress" nodes.json)
        host=${host1//\"/}
        port=5000
        user='ubuntu'
        key="~/.ssh/aws"
        id=$i
        node="node"$id

        expect -c "
set timeout -1
spawn scp -i $key $user@$host:bkr/i ./log/server$id/
expect 100%
exit
"

        expect -c "
set timeout -1
spawn scp -i $key $user@$host:bkr/log/server$id ./log/server$id/
expect 100%
exit
"
} &
done

wait
