package solid.lsp;

// this is for liskov substitution principle

public class Main {
	public static void main(String[] args) {
		Main m = new Main();
		Employee e = new Employee(1, "P");

		TraineeEmployee te = new TraineeEmployee(2, "J");
		m.print(e);
		m.print(te); // now here we see the derived class can be substituted with the parent class without any issue.

		//Now one example where it will break the liskov substitution principle.
		// Class -> Square
		// Class Rectangle which extends Square, now this will work but in real world it is not true
		// so to avoid this we can use this principle
	}

	void print(Employee e) {
		e.printMe();
	}
}
