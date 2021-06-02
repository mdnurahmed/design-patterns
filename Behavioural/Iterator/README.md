# iterator-pattern

This pattern is used when the collection or container provides an iterator which lets it go through each element in it in sequence without exposing its underlying implementation . For example an iterator that iterates nodes of graphs or binary serach tree.

Here we implimented a zigzag iterator which returns the leftmost unvisited element and then the rightmost unvisited element and keeps iterating the collection like this. 

For exmaple if the collection is [1,2,3,4] , it iterates 1,4,2,3 like this .
The iterator does not alter the collection , just iterates it maintaining some logic .  