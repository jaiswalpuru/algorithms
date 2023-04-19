package solid.lsp;

public class Employee {
	int id;
	String name;

	public Employee(int id, String name) {
		this.id = id;
		this.name = name;
	}

	//not overriding as I was following a tutorial
	public void printMe() {
		System.out.println("ID : " + id + "\nName : " + name + "\n");
	}
}
