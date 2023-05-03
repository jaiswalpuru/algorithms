package parallelstreams;

import java.util.Arrays;
import java.util.stream.Stream;

public class Main {
	public static void main(String[] args) {
		Employee[] emps = new Employee[7];
		emps[0] = new Employee("A", 5);
		emps[1] = new Employee("B", 4);
		emps[2] = new Employee("C", 3);
		emps[3] = new Employee("D", 2);
		emps[4] = new Employee("E", 1);
		emps[5] = new Employee("F", 6);
		emps[6] = new Employee("P", 7);

		Stream<Employee> employeeStream = Arrays.stream(emps);
		SalaryAdder sa = new SalaryAdder();
		employeeStream.parallel()
			.forEach(sa::acceptTwo);
		System.out.println(sa.totalOne);
	}
}
