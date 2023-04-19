package comparatorcomparable;

import java.util.Comparator;

public class Employee implements Comparable<Employee> {
	private int id;
	private String name;

	public Employee(int id, String name) {
		this.id = id;
		this.name = name;
	}

	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}

	public int getId() {
		return id;
	}

	public void setId(int id) {
		this.id = id;
	}

	@Override
	public int compareTo(Employee o1) {
		return getId() - o1.getId();
	}

	@Override
	public boolean equals(Object o) {
		if (o == null || getClass() != o.getClass()) return false;

		if (o == this)
			return true;

		Employee e = (Employee) o;
		return e.getId() == getId();
	}

	//just for example rather using prime numbers
	@Override
	public int hashCode() {
		return getId();
	}

	@Override
	public String toString() {
		return " ID : " + getId() + ", Name : " + getName();
	}

	public static Comparator<Employee> nameComparator = new Comparator<Employee>() {
		@Override
		public int compare(Employee o1, Employee o2) {
			return o1.getName().compareTo(o2.getName());
		}
	};
}
