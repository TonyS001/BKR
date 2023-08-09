#!/bin/bash

# (C) 2016-2023 Ant Group Co.,Ltd.
# SPDX-License-Identifier: Apache-2.0

ps -u `whoami` | grep main | awk '{system("kill -9 "$1)}'
