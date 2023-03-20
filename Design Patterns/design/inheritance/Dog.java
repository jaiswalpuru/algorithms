package inheritance;

public class Dog extends Animal {
	public void digHole() {
		System.out.println("Dug a hole");
	}

	public Dog() {
		super();
		setSound("....barks.....");
	}

	//this function is static, and it is not tied to this class.
	public static void changeName(Dog d) {
		d.setName("Diana");
	}

	public static void main(String[] args) {
		Dog d = new Dog();
		d.setName("Charles");
		System.out.println(d.getName());
		//reference is passed to the changeName function.
		changeName(d);
		System.out.println(d.getName());
	}
}
