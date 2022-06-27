# Kademlia DHT boot node 

what is it ? 

In order your p2p nodes to join your network , they need at least one node that helps them to find other nodes and start the distrubuted communication. once nodes up and running they don't need the availablity of the boot node. 


## How to run

```shell
$ go run .

2022/06/27 14:58:24 Host ID: QmW3TsL15x5b9bRDVUEodnXyxPaNHzp6JashgYz37CmhpJ
2022/06/27 14:58:24   /ip4/172.20.10.2/tcp/41943/p2p/QmW3TsL15x5b9bRDVUEodnXyxPaNHzp6JashgYz37CmhpJ
2022/06/27 14:58:24   /ip4/127.0.0.1/tcp/41943/p2p/QmW3TsL15x5b9bRDVUEodnXyxPaNHzp6JashgYz37CmhpJ
Kademlia DHT boot node : -->  QmW3TsL15x5b9bRDVUEodnXyxPaNHzp6JashgYz37CmhpJ

## How to use it
copy this node address and add it to your bootstrap nodes list on your joining peer. (thi node address from the above example is `/ip4/172.20.10.2/tcp/41943/p2p/QmW3TsL15x5b9bRDVUEodnXyxPaNHzp6JashgYz37CmhpJ`)

## Credit
to open source community - while some tries to make our day dificult and be smart ass, other beautiful mind make our day joyfull !! 


![image](https://user-images.githubusercontent.com/25494022/175942566-a1365cc2-c171-4002-9e69-e681b9397917.png)
