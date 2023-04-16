package hashcodeequlas;

public class Main {
	public static void main(String[] args) {
		Employee e1 = new Employee();
		e1.setId(1);
		e1.setName("John");

		Employee e2 = new Employee();
		e2.setId(1);
		e2.setName("Steven");

		// now these employees have been created at different memory location.
		// now how to compare these two objects and tell whether they are equal or not.
		System.out.println("Shallow comparison : " + (e1==e2)); //this will return true if memory locations are same.

		System.out.println("Deep comparison : " + e1.equals(e2));

	}
}
