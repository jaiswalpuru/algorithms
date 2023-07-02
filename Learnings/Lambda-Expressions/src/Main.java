import java.util.*;
import java.util.function.*;

public class Main {

	record Person(String firstName, String lastName) {
		@Override
		public String toString() {
			return firstName + " : " + lastName;
		}
	}

	public static void main(String[] args) {
		System.out.println("Hello world!");

		Person p = new Person("P", "J");
		System.out.println(p);

		List<Person> arr = new ArrayList<>(Arrays.asList(
			new Person("a", "b"),
			new Person("z", "c"),
			new Person("p", "a")
		));

		var comparatorClass = new Comparator<Person>() {
			@Override
			public int compare(Person p1, Person p2) {
				return p1.firstName.compareTo(p2.firstName);
			}
		};

		arr.sort(comparatorClass);
		System.out.println(arr);

		/*
		 *
		 *	(a, b) -> a.firstName.compareTo(b.firstName) this is equivalent ot
		 * 	int compare(T a, T b);
		 *
		 */

		/*
		 **    Functional Interface
		 *	It is an interface that has one, and only abstract method.
		 * 	This is how Java can infer the method, to derive the parameters and return type, for the lambda expression.
		 *
		 * 	Also know os SAM which is short form for single abstract method.
		 * 	A functional interface is the target type for lambda expression.
		 *
		 */

		arr.sort((a, b) -> a.firstName.compareTo(b.firstName));

		interface EnhancedInterface<T> extends Comparator<T> {
			int secondLevel(T o1, T o2);
		}

		var customComparator = new EnhancedInterface<Person>() {
			@Override
			public int compare(Person o1, Person o2) {
				int result = o1.lastName.compareTo(o2.lastName);
				return result == 0 ? secondLevel(o1, o2) : result;
			}

			@Override
			public int secondLevel(Person o1, Person o2) {
				return o1.firstName.compareTo(o2.firstName);
			}
		};

		// for this we cannot use lambda expression as the interface had two functions to be implemented.
		arr.sort(customComparator);
		System.out.println(arr);

		List<String> list = new ArrayList<>(List.of("alpha", "bravo", "charlie", "alpha"));
		for (String str : list) System.out.println(str);

		list.forEach((s) -> System.out.println(s));

		list.forEach((s) -> {
			char first = s.charAt(0);
			System.out.println("First character : " + first);
		});

		/*
		 *	We can also have variables inside the lambda expression only until
		 * 	it is final.
		 */

		String nato = "nato";
		list.forEach((s) -> {
			System.out.println(nato + " -> " + s);
		});

		/*
		 * 				Functional Programming
		 * 		This is a programming paradigm which focuses on computing and returning results.
		 */

		int result = calculator((a, b) -> a + b, 5, 2);
		var result2 = calculator((a, b) -> a.toUpperCase() + " " + b.toUpperCase(), "ralph", "p");
		System.out.println(result2);


		/*
		 *	Functional Interface, Consumer, Predicate
		 * 	Java provides a library of functional interface in java.util.function package
		 *
		 *
		 * 	Some basic functional interfaces in java.util.function package
		 *
		 * 	1. Consumer -> void accept(T t) -> execute code without returning data.
		 * 	2. Function -> R apply(T t) -> return a result of an operation or function.
		 * 	3. Predicate -> boolean test(T t) -> test if condition is true or false.
		 * 	4. Supplier -> T get() -> return an instance of something.
		 *
		 * NOTE : A lambda expression can be assigned to a local variable.
		 */

		/*
		 *			Two most common Consumer interfaces :
		 * 		1. Consumer -> void accept(T t)
		 * 			eg : s-> System.out.println(s);
		 *
		 * 		2. BiConsumer -> void accept(T t, U u)
		 *
		 */

		var coords = Arrays.asList(
			new double[]{47.2, 50},
			new double[]{67.5, 90.0}
		);

		coords.forEach(s -> System.out.println(Arrays.toString(s)));

		BiConsumer<Double, Double> p1 = (lat, lon) -> System.out.println("Lat : " + lat + ", long : " + lon);

		processPoint(coords.get(0)[0], coords.get(0)[1], p1);

		coords.forEach(s -> processPoint(s[0], s[1], p1));


		/*
		 *	PREDICATE Interface
		 * They are used to test a condition, and if the condition is true, some operation will be performed.
		 *
		 * There are two types of predicate interfaces
		 * 	1. Predicate -> boolean test(T t)
		 * 	2. BiPredicate -> boolean test(T t, U u)
		 *
		 * eg : s -> s.equalsIgnoreCase("Hello").
		 */

		list.removeIf(s -> s.equalsIgnoreCase("bravo"));
		System.out.println(list);

		list.forEach(System.out::println);

		/*
		 *	Function Interface
		 *	1. Function<T, R> -> R apply(T t) -> UnaryOperator<T> -> T apply(T t)
		 * 	2. BiFunction<T, U, R> -> R apply(T t, U u) -> BinaryOperator -> T apply(T t1, T t2)
		 *
		 *
		 *
		 * eg : (s) -> s.split(",")
		 *		R apply(T t)
		 * 		R -> is an array of strings
		 * 		T -> of type string
		 *
		 * 	IntFunction -> is an functional interface in Java which accepts one argument of type int
		 * 				and returns result R. The interface contains one method i.e. apply
		 */

		list.replaceAll(s -> s.charAt(0) + " - " + s.toUpperCase());
		System.out.println(list);

		String[] emp = new String[10];
		System.out.println(Arrays.toString(emp));

		Arrays.fill(emp, "");
		System.out.println(Arrays.toString(emp));

		Arrays.setAll(emp, i -> (i + 1) + ". " +
			switch (i) {
				case 1 -> "one";
				case 2 -> "two";
				default -> "default";
			}
		);
		System.out.println(Arrays.toString(emp));

		//********************************************************************************************

		/*
		 *	Supplier -> interface which takes no arguments but returns an instance of some type T.
		 * 	Supplier -> T get()
		 * 	Can think of something like factory method code.
		 *
		 * It will produce an instance of some object.
		 * However, this doesn't have to be new or distinct result returned.
		 *
		 *
		 *  eg : () -> random.nextInt(1, 100);
		 */

		String[] names = {"Ann", "Bob", "Carol", "Jake"};
		String[] randomList = randomlySelectedValues(15, names,
			() -> new Random().nextInt(names.length));

		System.out.println(Arrays.toString(randomList));

		/*
		 *
		 * 			Method References
		 * 	1. Java gives us an alternative syntax to use for this second kind of lambda.
		 * 	2. These are called method references.
		 * 	3. These are more compact, easier to read.
		 *
		 */

		List<String> l = new ArrayList<>(List.of("a", "b", "c", "d"));
		l.forEach(System.out::println);

		/*
		 *	Why are these equal ?
		 * 		Lambda expression				Method Reference
		 * 		s -> System.out.println()		System::println
		 *
		 * A method reference abstracts the lambda expression even further, eliminating the need to
		 * declare formal params.
		 *
		 * A method reference has "::" between qualifying type and the method name inferred.
		 *
		 * In this example of Consumer Interface, not only the method is inferred, but the params as well.
		 */

		calculator2(Integer::sum, 10, 24);

		/*
		 *	Whenever you create variables that are lambda expressions or method references, it's
		 * 	important to remember that the code isn't  invoked at that time.
		 *	The statement or code block gets invoked at the point in the code that the targeted
		 * 	functional method is called.
		 */
		Supplier<Q> sp = Q::new;
		Q q = sp.get();
		System.out.println(sp.get());

		/*
		 *	Method reference eg:
		 */

		BinaryOperator<String> b1= String::concat;
		BiFunction<String, String, String> bf = String::concat;
		System.out.println(b1.apply("a", " b"));

		/*
		 * Chaining functional interfaces.s
		 */

		String name = "Tim";
		Function<String, String> uCase = String::toUpperCase;
		System.out.println(uCase.apply(name));

		Function<String, String> lastName = s -> s.concat(" Buchaka");

		Function<String, String> uCaseLastName = uCase.andThen(lastName);
		System.out.println(uCaseLastName.apply(name));

		/*
		 * andThen() -> this will execute the uCase first and then will apply the last name.
		 * compose() -> opposite of andThen() in this case first lastName will be executed and then uCase will be applied
		 */

		Function<String, String> uCaseLastNameFirst = uCase.compose(lastName);
		System.out.println(uCaseLastNameFirst.apply(name));

		Function<String, String[]> f0 = uCase
			.andThen(s -> s.concat(" Buchaka"))
			.andThen(s -> s.split(" "));
		System.out.println(Arrays.toString(f0.apply(name)));

		String[] name2 = {"Anna", "John", "Dijkstra"};
		Consumer<String> s0 = s -> System.out.print(s.charAt(0));
		Consumer<String> s1 = System.out::println;
		Arrays.asList(name2).forEach(s0
			.andThen(s -> System.out.print(" - "))
			.andThen(s1));

		// predicate returns a boolean value
		Predicate<String> p2 = s -> s.equals("TIM");
		Predicate<String> p3 = s -> s.equalsIgnoreCase("Tim");
		Predicate<String> p4 = s -> s.startsWith("T");
		Predicate<String> p5 = s -> s.endsWith("e");

		Predicate<String> combined1= p2.or(p3);
		System.out.println("combined1 : " + combined1.test(name));

		Predicate<String> combined2 = p2.and(p3);
		System.out.println("combined 2 : " + combined2.test(name));

		Predicate<String> combined3 = p4.and(p5).negate();
		System.out.println("combined 3 : " + combined3.test(name));

		record Person(String firstName, String lastName) {}
		List<Person> parr = new ArrayList<>(Arrays.asList(
			new Person("ana", "mica"),
			new Person("pap", "qwe"),
			new Person("ada", "ghj")
		));

		// now if we want to sort the parr array with the last name then
		// approach one :
		parr.sort((s, b) -> s.lastName.compareTo(b.lastName));
		System.out.println(parr);

		System.out.println("-------------------");
		//approach two
		parr.sort(Comparator.comparing(Person::lastName));
		System.out.println(parr);

		System.out.println("-------------------");
		System.out.println("-------------------");

		parr.sort(Comparator.comparing(Person::lastName)
			.thenComparing(Person::firstName)); // here the person will be sorted first by last name and then by first name
		parr.forEach(System.out::println);

		//for reverse
		parr.sort(Comparator.comparing(Person::lastName)
			.thenComparing(Person::firstName).reversed());
		parr.forEach(System.out::println);
	}

	public static <T> void calculator2(BinaryOperator<T> function, T val1, T val2) {
		T result = function.apply(val1, val2);
		System.out.println(result);
	}

	public static String[] randomlySelectedValues(int count, String[] values, Supplier<Integer> s) {
		String[] selectedValue = new String[count];
		for (int i=0; i<count; i++) {
			selectedValue[i] = values[s.get()];
		}
		return selectedValue;
	}

	public static <T> void processPoint(T t1, T t2, BiConsumer<T, T> consumer) {
		consumer.accept(t1, t2);
	}

	public static <T> T calculator(Operation<T> function, T val1, T val2) {
		T result = function.operate(val1 ,val2);
		System.out.println("Result of operation : " + result);
		return result;
	}
}

class Q {
	static int d;

	public Q() {
		System.out.println("New " + d++);
	}
}


@FunctionalInterface
interface Operation<T> {
	T operate(T val1, T val2);
}