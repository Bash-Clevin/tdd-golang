# ARCHITECTURE USED HERE

**Hexagonal Architecture**, also known as Ports and Adapters, is a software design pattern that aims to create loosely coupled application components that can be easily connected to their software environment through ports and adapters. 

It was invented by Alistair Cockburn in 2005 to avoid structural pitfalls in object-oriented software design, such as undesired dependencies between layers and contamination of user interface code with business logic. 

The architecture puts inputs and outputs at the edges of the design, ensuring that business logic should not depend on whether we expose a REST or a GraphQL API, and it should not depend on where we get data from. 

The hexagonal architecture solves these problems by noting the symmetry in the situation: there is an application on the inside communicating over some number of ports with things on the outside, and the items outside the application can be dealt with symmetrically.

![Diagram](/resources/Hexagonal-architecture-2023-11-24%2006-21-24.png)