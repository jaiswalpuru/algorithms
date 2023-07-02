package streams;

public class Employee {
	private String name;
	private int salary;
	private String address;

	public Employee(String name, int salary) {
		this.name = name;
		this.salary = salary;
		this.address = "Junk";
	}

	public String getAddress() {
		return address;
	}

	public void setName(String name) {
		this.name = name;
	}

	public void setSalary(int salary) {
		this.salary = salary;
	}

	public String getName() {
		return name;
	}

	public int getSalary() {
		return salary;
	}
}
