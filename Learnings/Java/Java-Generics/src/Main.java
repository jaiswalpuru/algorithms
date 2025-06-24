import java.util.Comparator;

public class Main {
	public static void main(String[] args) {
		System.out.println("Hello world!");

		/*
		*			Java Generics
		* It allows us to create a single class, interface, etc...
		* that can be used by different data types.
		 */

		//eg : for creating a generic class
		GenericClass<Integer> g = new GenericClass<>(3);
		System.out.println(g.getData());

		// generic method
		GenericMethodClass g2 = new GenericMethodClass();
		System.out.println(g2.genericMethod("Here I am"));

		/*
		*		GenericClassTwo<String> obj = new GenericClassTwo<String>();
		* 	This is wrong as String is not a subclass of Number
		 */
		GenericClassTwo<Integer> n = new GenericClassTwo<>();

		GenericClassThere<P> gp = new GenericClassThere<>();


		GenericClassMultiple<Integer, String> gm = new GenericClassMultiple<>(1, "Here I am");
		System.out.println(gm.d);
		System.out.println(gm.q);

	}
}

class GenericClassMultiple<T, P> {
	T d;
	P q;
	public GenericClassMultiple(T a, P b){
		this.d = a;
		this.q = b;
	}
}

class P extends Number implements Comparator<P>{

	@Override
	public int intValue() {
		return 0;
	}

	@Override
	public long longValue() {
		return 0;
	}

	@Override
	public float floatValue() {
		return 0;
	}

	@Override
	public double doubleValue() {
		return 0;
	}

	@Override
	public int compare(P o1, P o2) {
		return 0;
	}
}

//Bounded type example two.
class GenericClassThere<T extends Number & Comparator<T>> {

}

/**
 * Bounded types.
 * <T extends A>
 *     This means T can only accept data that are subtypes of A
 */

class GenericClassTwo <T extends Number> {
	public void display() {
		System.out.println("This is a generic class which is bounded by Number");
	}
}


// generic class
class GenericClass<T> {
	public T data;
	public GenericClass(T data) {
		this.data = data;
	}
	public T getData() {
		return this.data;
	}
}

/*
* Similar to generic class, one can also create generic method.
 */

class GenericMethodClass {
	public <T> T genericMethod(T data) {
		System.out.println("Generic data : " + data);
		return data;
	}
}