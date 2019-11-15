This is about the code that I recently learn from Micro.  
If you interested anout the Micro, you can chick this url(https://github.com/micro-in-cn/tutorials) to learn the Micro.  
Now I'm going to introduce you to the structure of this directory and how to run.  
____
First  
====
you need understand what is the Micro and what is the go-micro.  
  
Micro is a microservices framework (or a collection of tools), it provides various compoents to solve different problems in the microservices
architecture,service monitoring, service discovery, fuse mechanism, load balancing... And you also see [the documents of Micro](https://micro.mu/docs/index.html). If you is Chinese , please see [the documents of Micro](https://micro.mu/docs/cn/index.html) 
  
The go-micro framework provides mechanisms for service discovery, load balancing, synchronous transport, 
asynchronous communication, and event-driven. It attempts to simplify communication between distributed systems, 
allowing us to focus on the development of our own business logic.  
  
This is structure of go-micro:  
  
![alt](https://image-static.segmentfault.com/180/641/1806415878-5c28d8815645a)  
  
The go-micro's compoents :  
* Registry component: The service discovery component, which provides a service discovery mechanism: resolves the service name to the service address. The currently supported registration centers are etcd, zookeeper, dns, gossip, etc.(Tip: In a recent update, the consul isn't use.)
- Selector component: A client-side intelligent load balancing component built on the Registry for the client component to intelligently select the services returned by the Registry.
* Broker component: Publish/subscribe component, the asynchronous communication method based on message middleware between services. By default, http mode is used. Message middleware such as Kafka and RabbitMQ are usually used on the line.
- Transport component: Synchronous communication between services.
* Codec component: Encoding/decoding of messages between services.
- Server component: The service body, which is based on the above Registry/Selector/Transport/Broker component, provides a unified service request entry.
* Client component: Provides access to the microservices client. Similar to the Server component, it also implements lookup services, load balancing, synchronous communication, asynchronous messaging, etc. through the Registry/Selector/Transport/Broker component.  
  
All of the above compoents functions together form a go-micro microservice.  
 
Sencond
====
you need download the tools ( but Micro seems to have included all the tools,so you just need download the micro.But in order to clearly understand the tools we need to get into this project, I will list them below. Note : there may be some tools which is you need,like protoc-gen-gorm..., you need use by yourself.In here I will not say much):  
* protobuf  
```
go get github.com/micro/protoc-gen-micro  
go get github.com/micro/protoc-gen-go
```
- Micro  
```
go get github.com/micro/micro
```
  
好了，夜已深，再不睡觉，大灰狼就来抓人了.日后在更新
