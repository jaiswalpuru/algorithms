# System Design Basics

![](https://github.com/jaiswalpuru/ALGO_DS_SD/blob/main/system_design/static/img/vs_hs.jpg)

## Vertical Scaling vs Horizontal Scaling
```sh
* Vertical scaling means adding more power to the CPU by increasing its resources such as RAM. This is done
by upgrading the current specs of system or adding more featuers to our current system.

* Horizontal scaling means adding more resources to our existing pool of resources. 
We can achieve this by adding more machines to our current setup.

What is the need for scaling ?

Let's say we design a system, and its performance keeps on getting slow as the number of requests increases, 
hence we need to scale the system.

Vertical scaling doesn't make the system fault tolerant, 
lets say we are running an application on a single server and if that machine goes down the complete 
application will be down for that moment.

Another is solution is to increase the amount of servers, 
this will decrease the amount of requests coming to one server and increasing the throughput.
By doing we need not restart the system afrer adding any server to the existing pool of resources. 
To segregate the requests, equally to each of the application server, we need to add a load balancer.
```
