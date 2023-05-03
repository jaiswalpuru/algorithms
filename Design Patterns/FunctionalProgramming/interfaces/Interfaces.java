package interfaces;

import java.util.Comparator;

public class Interfaces {
	public static void main(String[] args) {
		Employee mike = new Employee("Mike", 2000),
			louise = new Employee("Louise", 2500);

		Comparator<Employee> byName = new Comparator<Employee>() {
			@Override
			public int compare(Employee a, Employee b) {
				return a.getName().compareTo(b.getName());
			}
		};

		System.out.println("By Name : ");
		System.out.println(byName.compare(mike, louise));
		try {
			System.out.println(byName.compare(mike, null));
		}catch (NullPointerException e) {
			System.out.println(e);
		}

		Comparator<Employee> byNameLastsNull = Comparator.nullsLast(byName);
		System.out.println("The Null");
		System.out.println(byNameLastsNull.compare(mike, louise));
		System.out.println(byNameLastsNull.compare(mike, null));

		Comparator<Employee> byNameAscending = byNameLastsNull.reversed();
		System.out.println("reversed");
		System.out.println(byNameAscending.compare(mike, louise));
		System.out.println(byNameAscending.compare(mike, null));
	}
}
