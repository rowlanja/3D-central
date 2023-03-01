# Title : proof of authotity and optimiztion for sharded networks

# Section 1

## Introduction
Blockchain networks are stil in the infancy with network design. Theres still a lot of optimization to go with computation optimization in clustered blockchain networks. Cluster computing is the process of sharing the computation tasks among multiple computers, and those computers or machines form the cluster. It works on the distributed system with the networks. In this paper we will explore the application of proof of authority to sharded networks to see how much we can optimize the network. Proof-of-Authority (PoA) is a new consensus algorithms family that provides high performance and fault tolerance. In PoA, rights to generate new blocks are awarded to nodes that have proven their authority to do so. To gain this authority and a right to generate new blocks, a node must pass a preliminary authentication. In sharded networks, nodes in the network are partitioned into subsets called committee's, where each committee is assigned a seperate portion of transactions to process. By combining both PoA and sharding we aim to generate a lightweight protocol for sharded networks, and will also introduce a protocol for scaleing the shardedness as more authorised nodes are added


## Plan
1. Create cosmos chain
2. Introduce branch for expermentation
3. Add PoA module
4. Introduce transaction routing module
5. Introduce simulation
