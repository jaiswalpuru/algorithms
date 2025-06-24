package typeinference;

import java.util.function.Consumer;

/**
 * Code does not specify the type parameter for "getFirst"
 * Still no cast is required.
 * Type parameter is inferred from arr
 */
public class Main {
	public static void main(String[] args) {
		String[] arr = {"p", "q", "r"};
		String one = getFirst(arr);
		System.out.println(one);

		Consumer<String> c1 = msg -> System.out.println("Here I am");
//		Object o1 = msg -> System.out.println(msg.length()); compile time error.

//		Object o2 = (String msg) -> System.out.println(msg.length()); still compile time error not enough info

		Object o3 = (Consumer<String>) ((String msg) -> System.out.println(msg.length()));
		System.out.println(o3);


		// for the msg parameter the inferred type will be Object.
		Consumer<?> c2 = msg -> System.out.println(msg);

		// but this is ok
		Consumer<?> c3 = (String msg) -> System.out.println(msg.length());

	}

	public static<T> T getFirst(T[] array) {
		return array[0];
	}
}

/*
 * Context must contain enough info to identify the receiving functional interface.
 * RHS of assignment
 * Consumer<String> = lambda
 * Actual parameter of a method or constructor
 * new Thread(lambda)
 *
 * Argument of return instruction
 * return lambda
 *
 * Argument of cast
 * (Consumer<String>) lambda
 */

/**
 * Lambda expression can access static fields.
 * instance fields of enclosing objects.
 */