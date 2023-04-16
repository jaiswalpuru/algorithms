package singleton;

public class Main {
	public static void main(String[] args) {
		SingletonClass sg1 = SingletonClass.getSingletonClass();
		sg1.simpleMethod();

		SingletonClass sg2 = SingletonClass.getSingletonClass();
		sg2.simpleMethod();

		SingletonClass sg3 = SingletonClass.getSingletonClass();
		sg3.simpleMethod();

		System.out.println(sg1 == sg2);
	}
}
