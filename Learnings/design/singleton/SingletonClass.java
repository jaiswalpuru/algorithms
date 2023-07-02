package singleton;

public class SingletonClass {
	int val = 0;
	private static SingletonClass sg = new SingletonClass();

	private SingletonClass(){

	}

	public static SingletonClass getSingletonClass() {
		return sg;
	}

	public void simpleMethod() {
		val++;
		System.out.println("hashcode of a singleton object " + sg + " called : " + val);
	}

	public int getVal() {
		return val;
	}
}
