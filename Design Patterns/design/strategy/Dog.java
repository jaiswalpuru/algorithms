package strategy;

public class Dog extends Animal {
	public void dugHole() {
		System.out.println("Hole dug");
	}

	public Dog() {
		super();
		setSound("Bark");
		flyingType = new CantFly();
	}

	@Override
	public void setFlyingType(Fly flyingType) {
		super.setFlyingType(flyingType);
	}
}
