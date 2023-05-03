package capturingvalues;

import java.util.function.Consumer;

/**
 * Lambdas are more succinct.
 * Lambdas do not create additional class files.
 * Not every occurrence of a lambda creates new object.
 * Lambdas only work for functional interface.
 * Anonymous classes can have state (That is fields)
 */

public class Main {
	public static void main(String[] args) {

		// if we see the output of the below snippet
		// 5 different objects will be created.
		System.out.println("Anonymous classes");
		for (int i=0; i<5; i++) {
			Consumer<String> myPrinter1 = new Consumer<String>() {
				@Override
				public void accept(String s) {
					System.out.println("Consuming : " + s);
				}
			};
			myPrinter1.accept(myPrinter1.toString());
		}

		// using the same as above but using lambdas
		// here we see the objects are the same.
		// it means lambdas do not create extra objects.
		for (int i=0; i<5; i++) {
			Consumer<String> myPrinter1 = msg -> System.out.println("Consuming : " + msg);
			myPrinter1.accept(myPrinter1.toString());
		}

		//capturing constant lambda, one instance.
		System.out.println("\nCapturing constant lambda : ");
		final int secret = 42;
		for (int i=0; i<5; i++) {
			Consumer<String> myPrinter1 = msg -> System.out.println("Consuming : " + msg + " secret : " + secret);
			myPrinter1.accept(myPrinter1.toString());
		}

		(new Main()).foo();
	}

	private int id = 1;
	public void foo() {
		System.out.println("\n Instance capturing lambda : ");
		for (int i=0; i<5; i++) {
			//this is instance capturing
			Consumer<String> myPrinter1 = msg -> System.out.println("Consuming : " + msg + " id : " + id);
			myPrinter1.accept(myPrinter1.toString());
		}
	}
}
