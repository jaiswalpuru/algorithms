### Load Balancer

![](https://github.com/jaiswalpuru/ALGO_DS_SD/blob/main/system_design/static/img/lb.jpeg)

```sh
A load balacer evenly distributes incomming traffic among web servers that are defined in a load-balanced set.

Users connect to public IP of load balancer directly. With this setup web servers are unreachable directly by the clients anymore.
For security purpose private IP's are used for intercommunication.
Private IP's are reachable with in the same network. LB communicates with web servers through private IP's.

Pros of using LB :
• If server 1 goes offline, all the traffic will be routed to other servers.
• If traffic increases and two servers are not enough, we only need to add more servers and the rest LB can handle.
```