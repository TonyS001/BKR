// (C) 2016-2023 Ant Group Co.,Ltd.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"bkr-go/acs/server"
	"bkr-go/crypto/bls"
	"bkr-go/transport/info"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const CONFIG_FILE = "node.json"

func newLogger(id int) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"log/server" + strconv.Itoa(id),
	}
	cfg.Sampling = nil
	cfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	return cfg.Build()
}

func removeLastRune(s string) string {
	r := []rune(s)
	return string(r[:len(r)-1])
}

type Configuration struct {
	Id      uint64 `json:"id"`
	Port    int    `json:"port"`
	Key     string `json:"key_path"`
	Cluster string `json:"cluster"`
}

func main() {
	jsonFile, err := os.Open(CONFIG_FILE)
	if err != nil {
		panic(fmt.Sprint("os.Open: ", err))
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(fmt.Sprint("ioutil.ReadAll: ", err))
	}
	var config Configuration
	json.Unmarshal([]byte(data), &config)

	lg, err := newLogger(int(config.Id))
	if err != nil {
		panic(fmt.Sprintf("newLogger: ", err))
	}
	defer lg.Sync()

	addrs := strings.Split(config.Cluster, ",")
	fmt.Printf("%d %s %d\n", config.Id, addrs, len(addrs))

	bls, err := bls.InitBLS(config.Key, len(addrs), int(len(addrs)/3+1), int(config.Id))
	if err != nil {
		panic(fmt.Sprint("bls.InitBLS: ", err))
	}

	server.InitNode(lg, bls, info.IDType(config.Id), uint64(len(addrs)), config.Port, addrs)
}
