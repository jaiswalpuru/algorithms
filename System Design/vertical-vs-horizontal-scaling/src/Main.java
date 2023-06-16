public class Main {
	public static void main(String[] args) {
		System.out.println("vertical - vs - horizontal - scaling");

		/*
		 *							Vertical vs horizontal scaling
		 *
		 * 						Request ------------------------>
		 *		Client /												Server
		 * 			   \
		 * 			    		Response <------------------------
		 *
		 * 		Client communicates with the server using an API and API returns a response to the client.
		 *		Client requests the server using an API.
		 *
		 * 		Server running on a computer, might be connected to a database, and we might require to configure the
		 * 		endpoints which client will make a request to.
		 *
		 *
		 *		Cloud is a set of desktop which third party provides with some cost.
		 * 		eg : AWS.
		 *
		 * 		so what AWS does is provides a set of desktop where out code can run on, we can deploy our code using
		 * 		remote login.
		 *
		 * 		Why we need cloud is because it provides us with scalability, extensibility et-cetra.
		 *
		 * 		So there will be case where multiple clients are trying to connect but the server is not able to handle it.
		 *
		 * 		Scalability :
		 * 			1. One of the solution is to buy a bigger machine. -> Vertical scaling
		 * 			2. Buy more machines -> Horizontal scaling
		 *
		 *
		 * 				Horizontal Scaling    		| 			Vertical Scaling
		 *
		 * 		1. Load balancing required				1. Not required
		 * 		2. it is resilient						2. Single point failure
		 * 		3. Network calls (RPC) so this can		3. Interprocess communication, fast.
		 * 		   be slow
		 *
		 * 		4. Data is complicated to maintain		4. Data consistent.
		 *			(Data inconsistency)
		 *			eg : if there is a transaction
		 * 			then operations need to be atomic
		 * 			in that we need to lock all the
		 * 			servers in order for the transaction
		 * 			to proceed.
		 * 		5. scales well							5. Hardware limit. (as we are increasing the number of
		 *													resources in a single system.
		 *
		 * 		In real word both are used, some good characteristics of vertical scaling is data consistency and
		 * 		interprocess communication which is faster.
		 *
		 * 		and from horizontal scaling we see it scales well (when some other server crashes some other server might
		 * 		come up)
		 *
		 * 		Initially we can use vertical scaling when building a system, but when user start getting increased then we can go
		 * 		for horizontal scaling.
		 */

	}

}