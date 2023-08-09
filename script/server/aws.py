# (C) 2016-2023 Ant Group Co.,Ltd.
# SPDX-License-Identifier: Apache-2.0

import boto3
import json

regions = ["us-east-2","ap-southeast-1","ap-northeast-1","ca-central-1","eu-central-1"]

total = {}
total["nodes"] = []
clients = {}
clients["nodes"] = []
server_id = 0
client_id = 0
for region in regions:
    ec2 = boto3.client('ec2', region_name=region)
    Tags = [{'Key': 'Name', 'Value': 'Free'}]
    Filter = [
        {
            'Name': 'key-name',
            'Values': [
                'aws',
            ]
        }
    ]
    response = ec2.describe_instances(Filters=Filter)
    instances = []
    for i in range(len(response['Reservations'])):
        instances += response['Reservations'][i]['Instances']

    # --------------------------------
    # --------nodes-----------------
    for i in range(len(instances)):
        status = instances[i]['State']['Name']
        if status != "running":
            continue
        instance = {}
        instance['Id'] = server_id
        server_id += 1
        instance['InstanceId'] = instances[i]['InstanceId']
        instance['InstanceType'] = instances[i]['InstanceType']
        instance['PublicIpAddress'] = instances[i]['PublicIpAddress']
        instance['PrivateIpAddress'] = instances[i]['PrivateIpAddress']
        instance['ServerURL'] = "http://" + instances[i]['PublicIpAddress'] +":6000/client"

        total['nodes'].append(instance)

print("----- begin to load nodes ----")
file = "./nodes.json"
with open(file,"w") as f:
    json.dump(total,f)
print("----- load success ----")

for item in range(len(total['nodes'])):
    total['nodes'][item]['ServerURL'] = "http://" + total['nodes'][item]['PublicIpAddress'] +":6000/client"

print("----- begin to load clients ----")
file = "../client/clients.json"
with open(file,"w") as f:
    json.dump(total,f)
print("----- load success ----")
