# (C) 2016-2023 Ant Group Co.,Ltd.
# SPDX-License-Identifier: Apache-2.0

import json
nodes = {}

with open("nodes.json",'r') as f:
    print("-----begin to load----")
    nodes = json.load(f)
    print("-----load success!----")
nodes = nodes['nodes']
ipset = []
cluster = []
for i in range(len(nodes)):
    ipset.append("http://"+nodes[i]['PublicIpAddress']+":6000")

cluster = ','.join(ipset)
key_path = "./crypto"
pk = "./crypto"

for i in range(len(ipset)):

    file = "node%d.json"%(i,)
    data = {}
    data['id'] = i
    data['port'] = 6000 # 6000
    data['address'] = ipset[i]
    data['key_path'] = key_path
    data['pk'] = pk
    data['cluster'] = cluster
    with open(file,'w') as f:
        json.dump(data,f)





