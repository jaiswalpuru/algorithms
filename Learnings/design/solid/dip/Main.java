package solid.dip;

// Dependency inversion principle.
// for example lets say there is a class

class Address {

}

class Student {
	Address add;
	public Student() {
		add = new Address(); // now doing this it makes address and Student class tightly coupled.
	}

	// but to avoid this situation rather than instantiating

}

public class Main {
}
