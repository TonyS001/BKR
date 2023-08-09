#!/bin/bash

# (C) 2016-2023 Ant Group Co.,Ltd.
# SPDX-License-Identifier: Apache-2.0

NUM=$1
SUM=$2

for ((i = 0; i < SUM; i++)); do
{
        id=$(expr $NUM - 1 - $i)
        echo $id
        host1=$(jq ".nodes[$id].PublicIpAddress" nodes.json)
        host=${host1//\"/}
        port=5000
        user='ubuntu'
        key="~/.ssh/aws"
        node="node"$id

        expect <<-END
spawn ssh -oStrictHostKeyChecking=no -i $key $user@$host "cd;cd bkr/script;./stop.sh main"
expect EOF
exit
END
} &
done

wait
