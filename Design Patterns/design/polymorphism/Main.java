package polymorphism;

public class Main {
	public static void main(String[] args) {
		Animal[] animals = new Animal[2];
		animals[0] = new Dog();
		animals[1] = new Cat();
		System.out.println(animals[0].getSound());
		System.out.println(animals[1].getSound());
		speak(animals[0]);
		speak(animals[1]);

		Animal d = new Dog();
		((Dog)d).dugHole();

		Giraffe g = new Giraffe();
		g.setName("Giraffe");
		System.out.println(g.getName());

	}

	public static void speak(Animal a) {
		System.out.println(a.getSound());
	}
}
