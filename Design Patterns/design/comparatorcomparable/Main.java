package comparatorcomparable;

import java.util.Arrays;

public class Main {
	public static void main(String[] args) {
		Employee[] arr = new Employee[4];
		arr[0] = new Employee(5, "John");
		arr[1] = new Employee(4, "Pete");
		arr[2] = new Employee(3, "Neeraj");
		arr[3] = new Employee(2, "Ravi");

		Arrays.sort(arr);
		System.out.println(Arrays.toString(arr));
	}
}
