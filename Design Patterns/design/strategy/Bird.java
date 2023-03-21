package strategy;

public class Bird extends Animal {
	public void doSomething() {
		System.out.println("Doing something bird");
	}

	public Bird() {
		super();
		super.setSound("chirp");
		flyingType = new CanFly();
	}
}
