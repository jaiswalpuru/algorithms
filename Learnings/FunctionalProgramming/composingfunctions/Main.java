package composingfunctions;

import java.io.FileNotFoundException;
import java.io.PrintWriter;
import java.util.Comparator;
import java.util.function.Consumer;
import java.util.function.Function;

/**
 * Objective : writing the same string to screen and to a file
 */

public class Main {
	public static void main(String[] args) throws FileNotFoundException {
		PrintWriter writer = new PrintWriter("filename.txt");
		Consumer<String> logger = writer::println;
		Consumer<String> screener = System.out::println;
		Consumer<String> both = screener.andThen(logger);
		both.accept("Program started");

		Employee m = new Employee("Puru", 90000);
		//combining two functions sequentially
		Function<Employee, String> getName = Employee::getName;
		Function<String, Character> getFirstCharacter = name -> name.charAt(0);
		Function<Employee, Character> initialLetter = getName.andThen(getFirstCharacter);
		System.out.println(initialLetter.apply(m));

		//method then comparing
		Comparator<Employee> byName = Comparator.comparing(Employee::getName);
		Comparator<Employee> bySalary = Comparator.comparing(Employee::getSalary);
		Comparator<Employee> byNameThenSalary = byName.thenComparing(bySalary);
	}
}
