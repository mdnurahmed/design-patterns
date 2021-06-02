# factory-pattern 

This pattern is used when we need a way to hide the creation logic of the instances being created.The client only interacts with a factory struct and tells the kind of instances that needs to be created. The factory class interacts with the corresponding concrete structs and returns the correct instance back.



## Exaplaination 
In here , 

We have IVehicle interface which defines all methods a vehicle should have.
There is vehicle struct that implements the IVehicle interface.
Three concrete vehicles are car,bike and cycle. All embed vehicle struct and hence also indirectly implement all methods of IVehicle and hence are of IVehicle type
We have a vehicleFactory struct which creates the vehicle of type car,bike and cycle.

