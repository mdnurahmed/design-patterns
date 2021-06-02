# object-pool-pattern 

This pattern is used When the cost to create the object of the class is high and the object is immutable and the number of such objects that will be needed at a particular time is not much . So a pool of objects is initialized and created beforehand and kept in a pool. As and when needed, a client can request an object from the pool, use it, and return it to the pool. The object in the pool is never destroyed. It will boost the application performance significantly since the pool is already created

Some of the cases where the singleton object is applicable:

- DB connections

## Concurrency 
A mutex lock is used to make the object pool concurrency safe .