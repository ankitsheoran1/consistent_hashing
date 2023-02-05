# consistent_hashing
consistent hashing algorithm implementation in golang

# overview 
Basically consistent hashing can be used both as load balancer(It majorly treats here as stateless) and in databases side 
distribute the data and route each request to proper node. In case of databases we need to be more carefully like on addition of node we need to re-allocate data also to this node from other node 
and in same way of node eviction . 
Current  implementation is only for load balancing. 

# Key Topics 
  1. We have used crc32 hashing library , but its extendable you could use any one just implement that logic. 
  2. AddNode is for adding new node 
  3. replicationfactor is number of copies of same node need to be added
  4. Get is for getting nearest node to serve a request 


# Future Extension 
   Implement this logic for databases , which would be more challenging and more fun. 