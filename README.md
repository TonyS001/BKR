# BKR

## Getting Started


### Environment Setup
* latest verion of Golang (1.19 will be ok)

* python3+
  ```bash
  # boto3 offers AWS APIs which we can use to access the service of AWS in a shell 
  pip3 install boto3==1.16.0
  ```

* protobuf, the implementation of BKR uses protobuf to serialize messages, refer [this](https://github.com/protocolbuffers/protobuf) to get a pre-built binary. (libprotoc 3.14.0 will be ok)


### Build
```bash
# Download dependencies.
# under BKR/
cd src
go get -v -t -d ./...

# Build Node.
# under BKR/
cd src/acs/server/cmd
go build -o main main.go

# Build Client.
# under BKR/
cd src/client
go build -o client main.go
```

### Generate BLS keys
```bash
# Download dependencies.
# under BKR/
cd src/crypto
go get -v -t -d ./...

# Generate bls keys.
# Default n = 4 (f = 1, n = 3f + 1), t = 2 (t = f + 1)
# under BKR/
cd src/crypto/cmd/bls
go run main.go -n 4 -t 2
```


## Testing

1. Create a cluster of EC2 machines.

2. (if use AWS) Fetch machine information from AWS.
    ```bash
    # under BKR/
    cd script/server
    python3 aws.py
    ```

3. (if use AWS) Generate config file(`node.json`) for every node.
    ```bash
    # under BKR/script/server/
    python3 generate.py
    ```

4. Deliver nodes. 
    ```bash
    # Compress BLS keys.
    # under BKR/script/server/
    ./tarKeys.sh
    
    # Deliver to every node.
    # n is the number of nodes running in the test.
    ./deliverNode.sh n
    ```

5. Run nodes.
   ```bash
   # under BKR/script/server/
   ./beginNode.sh n
   ```

6. Deliver client. (Open another terminal)
   ```bash
   # under BKR/
   cd script/client
   ./deliverClient.sh n
   ```

7. Run client and wait for a period of time.
   ```bash
   # under BKR/script/client/
   ./beginClient.sh n <payload size> <batch size> <running time>
   # example: ./beginClient.sh 4 1000 10000 30
   # wait for <running time>
   ```

8. Copy result from client node.
   ```bash
   # create dirs to store logs
   # under BKR/script/client/
   mkdir log
   ./createDir.sh n 

   # fetch logs from remote machines
   ./copyResult.sh n output
   ```

9. Calculate throughput and latency.
   ```bash
   # under BKR/script/client/
   python3 cal.py n <batch size> log output <running time>
   # example: python3 cal.py 4 10000 log output 30
   ```

1. Stop nodes. (Back to node terminal)  
   ```bash
   # stop node process
   # under BKR/script/server/
   ./stopNode.sh n

   # clear node log files
   ./rmLog.sh n
   ```


## Brief introduction of scripts

### script/server

* aws.py: get machine information from AWS.

* generate.py：generate configuration for every node.

* tarKeys.sh: compress BLS keys.

* deliverNode.sh: deliver node to remote machines.

* beginNode.sh: run node on remote machines.

* stopNode.sh: stop node on remote machines.

* rmLog.sh: remove log file on remote machines.

### script/client
* deliverClient.sh: deliver client to remote machines.

* createDir.sh: create dirs to store client logs.

* copyResult.sh: fetch log files from remote machines.


## recommended parameters in our paper

We give the [4 and 10 nodes] parameters we used to test BKR presented in our paper.

### preparation (do once)
```bash
# under BKR/
cd src
go get -v -t -d ./...
cd acs/server/cmd
go build -o main main.go
cd ../../../client
go build -o client main.go
cd ../crypto
go get -v -t -d ./...
cd ../../script/client/
mkdir log
./createDir.sh 10
```

### fault-free 

#### n = 4 (Figure 4)

```bash
# node terminal
# under BKR/
cd src/crypto/cmd/bls
go run main.go -n 4 -t 2
cd ../../../../script/server
python3 aws.py
python3 generate.py
./tarKeys.sh
./deliverNode.sh 4
./beginNode.sh 4
```

```bash
# client terminal (open another terminal)
# under BKR/
cd script/client
./deliverClient.sh 4
./beginClient.sh 4 1000 1000 30
# please wait for 30 seconds
./copyResult.sh 4 output
python3 cal.py 4 1000 log output 30
```

```bash
# stop processes (back to node terminal)
# under BKR/script/server
./stopNode.sh 4
./rmLog.sh 4
```

To saturate the system and draw a curve, gradually increase the client parameter `batch size`, the third parameter in command `./beginClient.sh [running num] [payload] [batch size] [test time]` and the second parameter in command `python3 cal.py [running num] [batch size] log output [test time]`. **Note that the parameters should be same in these two commands in a test.** 

The peak throughput appears around `batch size = 10000`:

```bash
./beginClient.sh 4 1000 10000 30

python3 cal.py 4 10000 log output 30
```

#### n = 10 (Figure 5)

```bash
# node terminal
# under BKR/
cd src/crypto/cmd/bls
go run main.go -n 10 -t 4
cd ../../../../script/server
python3 aws.py
python3 generate.py
./tarKeys.sh
./deliverNode.sh 10
./beginNode.sh 10
```

```bash
# client terminal (open another terminal)
# under BKR/
cd script/client
./deliverClient.sh 10
./beginClient.sh 10 1000 1000 30
# please wait for 30 seconds
./copyResult.sh 10 output
python3 cal.py 10 1000 log output 30
```

```bash
# stop processes (back to node terminal)
# under BKR/script/server
./stopNode.sh 10
./rmLog.sh 10
```

To saturate the system and draw a curve, gradually increase the client parameter `batch size`, the third parameter in command `./beginClient.sh [running num] [payload] [batch size] [test time]` and the second parameter in command `python3 cal.py [running num] [batch size] log output [test time]`. **Note that the parameters should be same in these two commands in a test.** 

The peak throughput appears around `batch size = 12000`:

```bash
./beginClient.sh 10 1000 12000 30

python3 cal.py 10 12000 log output 30
```

### crash fault

#### n = 4 (Figure 7)

```bash
# node terminal
# under BKR/
cd src/crypto/cmd/bls
go run main.go -n 4 -t 2
cd ../../../../script/server
python3 aws.py
python3 generate.py
./tarKeys.sh
./deliverNode.sh 3
./beginNode.sh 3
```

```bash
# client terminal (open another terminal)
# under BKR/
cd script/client
./deliverClient.sh 3
./beginClient.sh 3 1000 1000 30
# please wait for 30 seconds
./copyResult.sh 3 output
python3 cal.py 3 1000 log output 30
```

```bash
# stop processes (back to node terminal)
# under BKR/script/server
./stopNode.sh 3
./rmLog.sh 3
```

To saturate the system and draw a curve, gradually increase the client parameter `batch size`, the third parameter in command `./beginClient.sh [running num] [payload] [batch size] [test time]` and the second parameter in command `python3 cal.py [running num] [batch size] log output [test time]`. **Note that the parameters should be same in these two commands in a test.** 

The peak throughput appears around `batch size = 10000`:

```bash
./beginClient.sh 3 1000 10000 30

python3 cal.py 3 10000 log output 30
```

#### n = 10 (Figure 8)

```bash
# node terminal
# under BKR/
cd src/crypto/cmd/bls
go run main.go -n 10 -t 4
cd ../../../../script/server
python3 aws.py
python3 generate.py
./tarKeys.sh
./deliverNode.sh 7
./beginNode.sh 7
```

```bash
# client terminal (open another terminal)
# under BKR/
cd script/client
./deliverClient.sh 7
./beginClient.sh 7 1000 1000 30
# please wait for 30 seconds
./copyResult.sh 7 output
python3 cal.py 7 1000 log output 30
```

```bash
# stop processes (back to node terminal)
# under BKR/script/server
./stopNode.sh 7
./rmLog.sh 7
```

To saturate the system and draw a curve, gradually increase the client parameter `batch size`, the third parameter in command `./beginClient.sh [running num] [payload] [batch size] [test time]` and the second parameter in command `python3 cal.py [running num] [batch size] log output [test time]`. **Note that the parameters should be same in these two commands in a test.** 

The peak throughput appears around `batch size = 12000`:

```bash
./beginClient.sh 7 1000 12000 30

python3 cal.py 7 12000 log output 30
```

# License

(C) 2016-2023 Ant Group Co.,Ltd.  
SPDX-License-Identifier: Apache-2.0
