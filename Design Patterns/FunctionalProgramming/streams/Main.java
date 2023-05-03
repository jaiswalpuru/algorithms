package streams;

import composingfunctions.Employee;

import java.util.*;
import java.util.function.Supplier;
import java.util.stream.Stream;

public class Main {
	public static void main(String[] args) {
		List<Integer> arr = new ArrayList<>();
		arr.add(1);
		arr.add(2);
		arr.add(3);

		//using enhanced for loop to iterate.
		for (Integer val: arr) System.out.println(val);

		//using iterators
		Iterator<Integer> it = arr.iterator();
		while(it.hasNext()) System.out.println(it.next());

		// using streams
		Stream<Integer> s = arr.stream();
		// idea is to pass a functional object to the for each method.
		s.forEach(System.out::println);

		Employee[] emps = new Employee[5];
		emps[0] = new Employee("A", 2300);
		emps[1] = new Employee("B", 2400);
		emps[2] = new Employee("C", 2500);
		emps[3] = new Employee("D", 2600);
		emps[4] = new Employee("E", 2700);

		Stream<Employee> empStream = Arrays.stream(emps);
		empStream.filter(e -> e.getSalary() >= 2500)
			.map(Employee::getName)
			.sorted()
			.forEach(System.out::println);

		Stream<Employee> eS = Stream.of(emps);
		eS.filter(e -> e.getSalary()>=200)
			.map(Employee::getName)
			.sorted()
			.forEach(System.out::println);

		//Below is an infinite stream.
		Random r = new Random();
		Stream<Integer> integerStream = Stream.generate(r::nextInt);

		final Random random = new Random();
		Supplier<Integer> supp = () -> {
			int result = random.nextInt();
			System.out.println("(supplying : "  +  result + ")");
			return result;
		};

		System.out.println("Test 1");
		Stream<Integer> randoms = Stream.generate(supp);
		System.out.println("First stream built : ");
		randoms.filter(n -> n >= 0)
			.limit(3)
			.forEach(System.out::println);

		System.out.println("\n\n\n Test 2");
		Stream<Integer> random2 = Stream.generate(supp);
		random2.limit(3).filter(n-> n>=0).forEach(System.out::println);

		//streams can only be traversed once
		Stream<Integer> s3 = Stream.of(1,2,3,4);
//		s3.forEach(System.out::println);
//		s3.forEach(System.out::println);


		//lets say we want to just print the first two elements
		// below method is wrong
//		s3.limit(2);
//		s3.forEach(System.out::println);
		// there are two ways of overcoming the above issue
//	1. 	s3.limit(2).forEach(System.out::println);

		// or 2.
		Stream<Integer> s5 = s3.limit(2);
		s5.forEach(System.out::println);

	}
}


// package com.example;

// import java.util.ArrayList;
// import java.util.Arrays;
// import java.util.List;
// import java.util.stream.Collectors;
// import java.util.stream.Stream;

// public class Main {

// 	public static void main(String[] args) {

// 		Integer[] scores = new Integer[]{80, 66, 73, 92, 43};
// 		Stream<Integer> scoreStream = Arrays.stream(scores);
// 		List<String> shoppingList = new ArrayList<>();
// 		shoppingList.add("coffee");
// 		shoppingList.add("bread");
// 		shoppingList.add("pineapple");
// 		shoppingList.add("milk");
// 		shoppingList.add("pasta");

// 		Stream<String> shoppingListStream = shoppingList.stream();

// 		Stream<String> lettersStream = Stream.of("a", "b", "c");
// 		Stream<String> sortedShoppingList = shoppingListStream.sorted();// will sort in alphabetical order;

// 		sortedShoppingList.forEach(System.out::println);


// 		shoppingList.stream()
// 			.sorted()
// 			.map(String::toUpperCase)
// 			.forEach(System.out::println);
// 		System.out.println("----------------------");
// 		List<String> sortedList = shoppingList.stream()
// 			.filter(item -> item.startsWith("p"))
// 			.map(String::toUpperCase)
// 			.collect(Collectors.toList());

// 		System.out.println(sortedList);
// 	}
// }

