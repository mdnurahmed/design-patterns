# singleton-pattern 

This pattern is used when only a single instance of the struct should exist. This single instance is called a singleton object. Some of the cases where the singleton object is applicable:

- DB instance 
- Logger instance 

## Concurrency 
Lock should be used to make the singleton object concurrency safe. A sync.RWMutex is preferable for data that is mostly read, and the resource that is saved compared to a sync.Mutex is time since readers don't have to wait for each other. They only have to wait for writers holding the lock.