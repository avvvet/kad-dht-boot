# Kademlia DHT boot node 

what is it ? 

In order your p2p nodes to join your public network , they need at least one node that helps them to find other nodes and start the distrubuted communication. once nodes up and running they don't need the availablity of the boot node. 


## How to run

```shell
$ go run .

2022/06/27 14:58:24 Host ID: QmW3TsL15x5b9bRDVUEodnXyxPaNHzp6JashgYz37CmhpJ
2022/06/27 14:58:24   /ip4/172.20.10.2/tcp/41943/p2p/QmW3TsL15x5b9bRDVUEodnXyxPaNHzp6JashgYz37CmhpJ
2022/06/27 14:58:24   /ip4/127.0.0.1/tcp/41943/p2p/QmW3TsL15x5b9bRDVUEodnXyxPaNHzp6JashgYz37CmhpJ
Kademlia DHT boot node : -->  QmW3TsL15x5b9bRDVUEodnXyxPaNHzp6JashgYz37CmhpJ

## How to use it
copy this node address and add it to your bootstrap nodes list on your joining peer. 
(boot node address from the above example is `/ip4/172.31.91.153/tcp/39049/p2p/QmYZYpGMGd2m9JcUrB9Sdfik7nG2qVTYQ5V6zAVSe6dHTh`)

## server
server which the boot node running, need to have public ip address. make sure when the bood start shows the public ip not the 
private ip. 

if you use server from digitialocean, it will show you the public ip and it works directly. However on AWS EC2 you need to open port number 
of the boot node on the ec2 network security.
 

```

![image](https://user-images.githubusercontent.com/25494022/175945690-53ab8fb3-d2ba-4fba-89f3-5272382e3ddc.png)

## Credit
to open source community - while some tries to make our day dificult and be smart ass, other beautiful mind make our day joyfull !! ❤️
