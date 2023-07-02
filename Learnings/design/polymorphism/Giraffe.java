package polymorphism;

public class Giraffe extends Creature {

	private String name;
	private int weight;

	@Override
	public void setName(String name) {
		this.name = name;
	}

	@Override
	public String getName() {
		return name;
	}

	@Override
	public void setWeight(int weight) {
		this.weight = weight;
	}

	@Override
	public int getWeight() {
		return weight;
	}
}
