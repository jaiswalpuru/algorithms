package streams;

import java.util.*;
import java.util.stream.Collectors;
import java.util.stream.IntStream;
import java.util.stream.Stream;

public class Filtering {
	public static void main(String[] args) {
		//Select 10 random positive integers
		final Random random = new Random();
		Stream<Integer> randoms = Stream.generate(random::nextInt);
		randoms.filter(n -> n>=0)
			.distinct()
			.limit(10)
			.forEach(System.out::println);

		// take while here will print until the elements generated are positive after
		// that all the integers will be ignored.
		Stream<Integer> randoms2 = Stream.generate(random::nextInt);
		randoms2.takeWhile(n -> n>=10)
			.forEach(System.out::println);

		// take while is more efficient as it will stop early rather than filter.

		Employee[] emps = new Employee[10];
		emps[0] = new Employee("A", 20000);
		emps[1] = new Employee("B", 30000);
		emps[2] = new Employee("C", 40000);
		emps[3] = new Employee("D", 50000);
		emps[4] = new Employee("E", 60000);
		emps[5] = new Employee("F", 70000);
		emps[6] = new Employee("G", 80000);
		emps[7] = new Employee("H", 90000);
		emps[8] = new Employee("I", 100000);
		emps[9] = new Employee("J", 110000);
//
//		Stream<Employee> employeeStream = Arrays.stream(emps);
//		employeeStream.map(Employee::getAddress)
//			.map(Address::getCity)
//			.map(City::getCity)
//			.distinct()
//			.forEach(System.out::println);

		// print the names of the 10 employees with the highest salary
		System.out.println("\n\n");
		Stream<Employee> employeeStream = Arrays.stream(emps);
		employeeStream.sorted(Comparator.comparingInt(Employee::getSalary).reversed())
			.limit(10)
			.map(Employee::getName)
			.forEach(System.out::println);

		List<Integer> arr= new ArrayList<>();
		arr.add(1);
		arr.add(1);
		arr.add(2);
		arr.add(2);
		arr.add(2);
		//efficiently count the number of distinct integers in arr
		long n = arr.parallelStream()
			.unordered()
			.distinct()
			.count();
		System.out.println(n);

		Collection<Employee> employeeCollection = Arrays.asList(emps);
		boolean nameValid = employeeCollection.stream()
			.allMatch(e -> e.getName() != null && e.getName().length()!= 0);
		System.out.println(nameValid);

		List<String> arr2 = new ArrayList<>();
		arr2.add("Abs");
		arr2.add("Abss");
		arr2.add("Absaa");
		arr2.add("Abqqqqs");
		Optional<String> longest = arr2.stream().max(Comparator.comparingInt(String::length));
		System.out.println(longest);

		List<String> strings = new ArrayList<>();
		strings.add("puru");
		strings.add("neeraj");
		strings.add("ravi");
		strings.add("sridhar");

		StringBuilder sb = strings.stream()
			.collect(
				() -> new StringBuilder(), // factory
				(StringBuilder builder, String s) -> builder.append(" ").append(s), //accumulator
				(StringBuilder one, StringBuilder two) -> one.append(two) // combiner
			);
		System.out.println(sb.toString());
		StringBuilder sb2 = strings.stream().collect(
			StringBuilder::new,
			StringBuilder::append,
			StringBuilder::append
		);

		System.out.println(sb2.toString());

		String summary = strings.stream().collect(Collectors.joining()); // alternatively String.join can also be used
		System.out.println(summary);

		Stream<Employee> employeeStream1 = Arrays.stream(emps);
		TreeSet<Employee> tree = employeeStream1.collect(
			Collectors.toCollection(
				() -> new TreeSet<>(
					Comparator.comparing(Employee::getSalary)
				)
			)
		);
		System.out.println(tree);
		tree.forEach(e -> System.out.println(e.getName()));

		Stream<Employee> employeeStream2 = Arrays.stream(emps);
		Map<String, Integer> map = employeeStream2.collect(
			Collectors.toMap(Employee::getName, Employee::getSalary)
		);
		System.out.println(map);

		Stream<Employee> employeeStream3 = Arrays.stream(emps);
		Map<Integer, List<Employee>> brackets = employeeStream3.collect(
			Collectors.groupingBy(e -> e.getSalary()/1000)
		);

//		for (int i=0; i<10; i++) System.out.println(i);
		IntStream.range(0, 10).forEachOrdered(System.out::println);

		OptionalDouble avgSalary = Arrays.stream(emps).mapToInt(Employee::getSalary).average();
		System.out.println(avgSalary);

		int[] s = {1, 2,3, 4};
		OptionalInt max = Arrays.stream(s).max();
	}
}
