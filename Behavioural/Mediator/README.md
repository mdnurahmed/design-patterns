# mediator-pattern

This pattern is used when we want to prevent dirrect communication between objects they become losely coupled. 

One very good example of a mediator pattern is the railway system platform.  Two trains never communicate between themselves for the availability of the platform. The stationManager acts as a mediator and makes the platform available to only one of the trains. The train connects with stationManager and acts accordingly. It maintains a queue of waiting trains. In case of any train leaving a platform, it notifies one of the train to arrive on the platform next.

## Concurrency 
A mutex lock is used to make the mediator concurrency safe .