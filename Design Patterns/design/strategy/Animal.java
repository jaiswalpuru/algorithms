package strategy;

public class Animal {
	private String name;
	private double height;
	private int weight;
	private String sound;

	//behavior can change without side effects.
	public Fly flyingType;

	public String getSound() {
		return sound;
	}

	public void setSound(String sound) {
		this.sound = sound;
	}

	public void setName(String name) {
		this.name = name;
	}

	public void setHeight(double height) {
		this.height = height;
	}

	public void setWeight(int weight) {
		this.weight = weight;
	}

	public int getWeight() {
		return weight;
	}

	public String getName() {
		return name;
	}

	public double getHeight() {
		return height;
	}

	public String canFly() {
		return flyingType.fly();
	}

	public void setFlyingType(Fly flyingType) {
		this.flyingType = flyingType;
	}

	@Override
	public String toString() {
		return "Animal{" +
			"name='" + name + '\'' +
			", height=" + height +
			", weight=" + weight +
			'}';
	}
}
