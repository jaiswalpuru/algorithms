public class Main {
	public static void main(String[] args) {
		System.out.println("Hello world!");
		/*
		 *						Monolithic VS Microservices
		 *
		 * 		*******************************Below definitions are wrong**********************************
		 * 		In monolith all the functionalities are performed on a single machine, while in microservices architecture
		 * 		there will be different machines performing different actions, and might have their own database.
		 * 		*******************************Above definitions are wrong**********************************
		 *
		 *
		 *
		 * 		Monolith doesn't need to be single machine in terms on which we are running it on.
		 * 		There can be multiple machines running the same monolith, we can also scale out using monolith.
		 *
		 *
		 * 		Microservices :
		 *		Idea is to split the application into set of smaller, interconnected services, instead of building a single
		 * 		monolithic application. Each microservice is a small application, that has its own hexagonal architecture,
		 * 		consisting of business logic along with various adapters. Each service will have its own database schema to
		 * 		ensure loose coupling.
		 *
		 * 		Advantages over microservices :
		 * 			1. The microservices architecture is easier to reason about/design for a complicated system.
		 * 			2. They allow new members to train for a shorter period of time, and have less context before
		 * 				touching the system.
		 * 			3. Deployments are fluid and continuous for each service.
		 * 			4. They allow decoupling service logic on the basis for service responsibility.
		 * 			5. Individual services can be written in different programming languages.
		 * 			6. The developer team can talk to each other through API sheets, instead of working on the same repo.
		 * 			7. New services can be tested easily and individually.
		 *
		 *
		 * 		Where microservices are not favorable over monoliths :
		 * 			1. The technical/ development team is very small.
		 * 			2. The service is simple to think as a whole.
		 * 			3. The service requires high efficiency, where network calls are avoided as much as possible.
		 * 			4. All developers must have context of all services.
		 *
		 *
		 */
	}
}