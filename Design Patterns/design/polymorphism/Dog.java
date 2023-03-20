package polymorphism;

public class Dog extends Animal{
	public Dog() {
		super();
		setSound("Bark");
	}

	public void dugHole() {
		System.out.println("Dug a hole");
	}
}
