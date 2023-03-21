package strategy;

public class Main {
	// now for the strategy pattern if I want to add a feature such as flying to each class
	// so rather than duplicating code in each child class creating a interface of that type
	public static void main(String[] args) {
		Animal d = new Dog();
		Animal b= new Bird();
		System.out.println("Dog : " + d.canFly());
		System.out.println("Bird : " + b.canFly());
		d.setFlyingType(new CanFly());
		System.out.println("Dog : " + d.canFly());
	}
}

//When to use strategy pattern :
//When you define a class that will have one behavior that is similar to its other behavior.
// When you need to use one of several behaviors dynamically.
