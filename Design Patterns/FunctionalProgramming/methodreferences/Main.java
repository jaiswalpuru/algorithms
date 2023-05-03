package methodreferences;

import java.util.Comparator;
import java.util.function.Consumer;
import java.util.function.Function;
import java.util.function.Supplier;

public class Main {

	interface ThreadSupplier {
		Thread giveMeaThread();
	}
	public static void main(String[] args) {
		//functional interface
		Supplier<Thread> s1 = Thread::currentThread;

		ThreadSupplier ts = Thread::currentThread;

		Employee frank = new Employee("frank", 3000);
		Integer i = frank.getSalary();
		Supplier<Integer> s2 = frank::getSalary;

		System.out.println(s2.get());

		Consumer<String> c1 = System.out::println;

		Function<Employee, Integer> f1 = Employee::getSalary;

		Integer frankSalary = f1.apply(frank);

		Comparator<Employee> byName = Comparator.comparing(Employee::getName);

		Employee[] dept = new Employee[5];
		dept[0] = new Employee("A", 1000);
		dept[1] = new Employee("B", 2000);
		dept[3] = new Employee("C", 3000);
		dept[4] = new Employee("D", 4000);
		printAll(dept, Employee::getName);
	}

	public static <T> void printAll(T[] arr, Function<T, String> toStringFun) {
		int i  = 0;
		for (T t : arr)
			System.out.println(i++ + "\t" + toStringFun.apply(t));
	}
}
