package lambdaexpressions;

import java.util.Comparator;
import java.util.function.Consumer;

public class Main {
	public static void main(String[] args) {
		// old way using anonymous class.
		Comparator<Employee> byName = new Comparator<Employee>() {
			@Override
			public int compare(Employee employee, Employee t1) {
				return employee.getName().compareTo(t1.getName());
			}
		};

		Comparator<Employee> byNameLambda1 =
			(Employee a, Employee b) -> { return a.getName().compareTo(b.getName()); };

		//removing parameter types.
		Comparator<Employee> byNameLambda2 =
			(a, b) -> { return a.getName().compareTo(b.getName());};

		//removing braces and return statement.
		Comparator<Employee> byNameLambda3 =
			(a, b) -> a.getName().compareTo(b.getName());

		/*
		 * Note : you cannot remove braces and have return keyword that
		 * will throw compile time error.
		 */

		//Expression with no parameters.
		Runnable r = () -> {
			System.out.println("Here");
		};

		Thread t1 = new Thread(r);
		t1.start();

		// now the above can be combined in a single line.

		Thread t2 = new Thread(() -> {
			System.out.println("Here I am");
		});

		t2.start();

		Consumer<String> lengthPrinter =
			s -> System.out.println(s + " -> " + s.length());

	}
}
