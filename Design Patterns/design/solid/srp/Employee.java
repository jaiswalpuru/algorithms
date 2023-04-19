package solid.srp;

public class Employee {
	private int id;
	private String name;

	Employee(int id, String name) {
		this.id = id;
		this.name = name;
	}

	public void printMe() {
		System.out.println("ID : " + id + "\n NAME : " + name+"\n");
	}

	// now lets say we also want address to be stored, so rather than writing it here
	// create another class which will handle the address.

}
