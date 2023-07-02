package hashcodeequlas;

public class Employee {
	private int id;
	private String name;

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
}
