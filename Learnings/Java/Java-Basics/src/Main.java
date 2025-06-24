public class Main {
	public static void main(String[] args) {
		System.out.println("Hello world!");
		/*
		*		-----------------------------------Java Basics------------------------------------
		*
		* 		Data Type :
		* 				1. byte -  8
		* 				2. char - 16
		* 				3. short - 16
		* 				4. int - 32
		* 				5. float - 32
		* 				6. long - 64
		* 				7. double - 64
		* 				8. boolean / ~
		*
		* 	boolean has only two values which is true or false. JVM converts them to 1 or 0 (int) at compile time, but the
		* 	specific size is not clearly specified.
		 */

		System.out.println("Byte : " + Byte.SIZE);
		System.out.println("Char : " + Character.SIZE);
		System.out.println("Short : " + Short.SIZE);
		System.out.println("Int : " + Integer.SIZE);
		System.out.println("Float : " + Float.SIZE);
		System.out.println("Long : " + Long.SIZE);
		System.out.println("Double : " + Double.SIZE);

		/*
		*	------------------------------------------------------------------------------------
		*
		* 	Types of Packaging :
		* 		Basic types have corresponding packaging types, and the assignment between basic types and their
		* 		corresponding packaging types is done using autoboxing and unboxing.
		 */

		Integer x = 2; // autoboxing, Integer.valueOf (2);
		int y = x;	// unboxing, x.intValue();

		/*
		*							Cache Pool
		*
		* 		The difference between new Integer(123) and Integer.valueOf(123) :
		* 				1. new Integer creates a new object every time.
		* 				2. Integer.valueOf(123) will use the object present in the buffer pool, multiple call will get
		* 					the same reference.
		*
		*
		 */

		Integer x1 = new Integer(123);
		Integer y1 = new Integer(123);
		System.out.println((x1==y1)); // false
		Integer p = Integer.valueOf(123);
		Integer q = Integer.valueOf(123);
		System.out.println(p == q); // true

		/*
		*			Implementation of valueOf operator.
		* 		public static Integer valueOf(int i) {
		* 			if (i >= IntegerCache.low && i <= IntegerCache.high) return IntegerCache.cache[i+(-Integer.cache.low)]
		* 			return new Integer(i);
		* 		}
		*
		*
		* 		From Java 8 the Integer buffer pool defaults from -128 to 127
		* 		The compiler will call the valueOf() during autoboxing process, so multiple integers with the same,
		* 		value within the scope of the buffer pool are created using autoboxing, and they refer to the same
		* 		object.
		*
		* 		Integer m = 123;
		* 		Integer n = 123;
		* 		sout(m == n); true
		*
		* 		The buffer pool corresponding to the basic types are as follows:
		* 			1. boolean value true and false.
		* 			2. all byte values
		* 			3. short values between -128 and 127
		* 			4. int values between -128 and 127
		* 			5. char in range \u0000 to \u007f
		*
		* 		When using the packaging types corresponding to these basic types, if the value range is within the scope
		* 		of the buffer pool, you can directly use the objects in the buffer pool.
		*
		* 		Integer buffer pool is very special among these different pools.
		* 		Lower bound is -128 and upper bound is 127 by default, but this upper bound is adjustable.
		* 		When starting the jvm pass -XX:AutoBoxCacheMax=<size> to specify the size of the buffer pool,
		* 		this option will set a system property named java.lang.IntegerCache.high when the JVM is initialized,
		* 		and then IntegerCache will read the system property when it is initialized attribute to determine the
		* 		upper bound.
		*
		 */


		//--------------------------------------------------------------------------------------------------------------
		/*
		*											Strings
		*
		*		String is declared as final, so it cannot be inherited. (Wrapper class such as Integer cannot be inherited.)
		*
		* 		In JAVA 8 String internally uses, char array to store data.
		* 		public final class String implements java.io.Serializable, Comparable<String>, CharSequence{
		*			private final char value[];
		* 		}
		*
		* 		In JAVA 9, the implementation of String class uses a coder to identify the encoding.
		* 		public final String implements java.io.Serializable, Comparable<String>, CharSequence {
		* 			private final byte[] value;
		* 			private final byte coder;
		* 		}
		*
		* 		The value array is declared final, that is once it is initialized, it cannot refer other array elements.
		*
		*
		* 		Immutable benefits :
		* 			1. The hash value can be cached.
		* 					Because the hash value of string is often used as key in hash map. The immutable feature make
		* 					the hash value also immutable, so only calculation is required.
		*
		* 			2. The need for String pool.
		* 					If a string object has been created, then it will get a reference from the String pool.
		* 					String pool is only possible if String is immutable.
		*
		*
		* 			3. Security.
		* 					String is often used as a parameter, and the immutability of String can ensure that the param
		* 					is immutable. eg: If the String is mutable as a network connection param, then the String
		* 					during the network connection, and the party who changes the string thinks that it is
		* 					connecting to other hosts, but the actual situation is not necessarily so.
		*
		* 			4. Thread safety.
		* 					String immutability is inherently thread safe and can be used safely from multiple threads.
		*
		*
		*
		* 	--------------------------------------------------------------------------------------------------------
		* 							String, StringBuffer, StringBuilder
		*
		* 			1. Variability :
		* 				String is immutable whereas StringBuffer and StringBuilder are mutable.
		*
		* 			2. Thread safety :
		* 				String is immutable and thus is thread safe.
		* 				StringBuilder is not thread safe.
		* 				StringBuffer is thread safe and uses synchronized internally for synchronization.
		*
		*
		* 			String Pool :
		*				String constant pool (String pool) saves all string literals, these literals are determined
		* 				at compile time.
		* 				String's "intern()" method can also be used to add string to the string pool.
		*
		* 				When a String calls the intern() method, if there is already a string, in the String Pool
		* 				that is equal to the string(using the equals() method to determine), then a reference to the
		* 				string in the String Pool will be returned, else a new String will be added to the String Pool,
		* 				and a reference to the new String will be returned.
		*
		 */

		String s1 = new String("aaa");
		String s2 = new String("aaa");
		System.out.println(s1 == s2); // false
		String s3 = s1.intern();
		String s4 = s2.intern();
		System.out.println(s3 == s4); // true

		/*
		*		If the string is created in literal form then the String will automatically put into the String Pool.
		 */

		String s5 = "bbb";
		String s6 = "bbb";
		System.out.println(s5 == s6); // true

		/*
		 *		Before Java7 String pool was placed in the runtime constant pool, which belonged to the permanent
		 * 		generation. And in Java7 String pool was moved to heap. This is because the space of permanent generation
		 * 		is limited, and OutOfMemoryError will occur in scenarios where large number of string are used.
		 */

		/*
		*						new String("abc")
		*		In this way a total of two object strings will be created (provided that there is not "abc" in String pool)
		* 		• "abc" is a string literal, so a string object will be created in the String pool during compilation, pointing
		* 		to this "abc" string literal
		* 		• The way to use new will create a string object in heap
		*
		*		NOTE : when a string object is used as the constructor parameter of another String object the contents
		* 			   of the value array will not be completely copied, but will point to the same value array.
		*
		* 		Operation :
		* 			Parameter passing :
		* 				Java parameters are passed to methods in the form of value passing, not reference passing.
		*
		* 			let say there is a class Dog
		* 			public class Dog {
		* 				String name;
		* 				Dog(String name) { this.name = name; }
		*			}
		*
		* 			Changing the field value of object in the method will change the field value of original object,
		* 			because the same object is referenced.
		*
		* 			class PassByValueReference {
		* 				psvm(String[] args) {
		* 					Dog d = new Dog("A");
		* 					func(d);
		* 					sout(d.getName());
		* 				}
		*
		* 				private void func(Dog d) {
		*					d.setName("xyz");
		* 					sout(d.name); //B
		* 				}
		* 			}
		*
		* 		But if the pointer refers to other objects in the method, then two pointers in the method and outside
		* 		the method point to different objects, and changing the content of the object pointed to by one pointer
		* 		has no effect on the object pointed to by the other pointer.
		*
		*		public class PassByValue {
		* 			psvm(String[] args) {
		* 				Dog d = new Dog("A");
		* 				sout(d.getObjectAddress()); // 0x123;
		* 				func(d);
		* 				sout(d.name);// A
		* 			}
		*
		* 			public static void func(Dog d) {
		* 				sout(d.getObjectAddress()); //0x123
		* 				d = new Dog("C");
		* 				sout(d.getObjectAddress()); //0x456
		* 				sout(d.name); // C
		* 			}
		*
		* 		}
		*
		*
		* 	------------------------------------float and double------------------------------------------
		* 	Java cannot perform downcasting implicitly, as this would result in loss of precision.
		* 	1.1 Literals are of double type, and 1.1 cannot be directly assigned to a float variables, because this
		* 	will be downcasting.
		*
		* 	float f = 1.1; // not allowed
		*
		* 	1.1f -> this is allowed
		* 	float f = 1.1f;
		*
		*
		*
		* 	Implicit type conversion :
		* 		Since literal 1 is of type int which has a higher precision than short, tf int cannot be downcasted to
		* 		short.
		*
		* 		short s = 1;
		*
		* 		But using the += and ++ operator will perform an implicit type conversion.
		* 		s += 1;
		* 		s++;
		* 		The above statement is equivalent to downcasting like below.
		* 		s = (short) s + 1;
		*
		*
		*
		* 								Switch statements
		*
		* 			Beginning with JDK 7, objects can be used in switch conditional statements.
		* 			String s = "a";
		* 			switch(s){
		* 				case "a":
		* 					sout("aaa");
		* 					break;
		* 				case "b":
		* 					sout("bbb");
		* 					break;
		* 			}
		*
		* 			Switch does not support long, float double , because the original intention of switch is to judge
		* 			the equivalence of those types with only a few values. If the value is too complicated, it is more
		* 			appropriate to use if.
		*
		*
		*
		*
		* 			final :
		*
		* 					1. Data :
		* 						Once a variable is declared final it cannot be changed.
		* 						- For primitive types, final makes the value unchanged.
		* 						- For reference types final makes the reference unchanged, so other objects cannot be
		* 						  referenced, but the referenced objects itself can be modified.
		*
		*
		* 						final int x = 2;
		* 						x = 2 not allowed to assign to a final variable.
		* 						final A y = new A();
		* 						y.a = 2;
		*
		* 				 	2. Method :
		* 						A method declared final cannot be overridden by the subclasses.
		* 						But if a method with the same signature is defined in subclass then it will be
		* 						new implementation rather than overridding it.
		*
		* 					3. A class declared as final cannot be inherited.
		*
		*
		*
		* 			static :
		* 				1. Static variables :
		* 						• Static variables : These variable belong to the class, and all instance of the class
		*	 						will share the same variable.
		* 						Static variables can be accessed through class names.
		* 						• Instance variables: Every time an instance is created, an instance variable is generated
		* 							which lives and dies with the instance.
		*				2. Static methods :
		* 						A static method exists when class is loaded, and it does not depend on any instance.
		* 						So a static method must have an implementation i.e it cannot be an abstract method.
		*						It can only access static fields and static methods of class to which it belongs, the methods
		* 						cannot have this or super keywords, because these two keywords are associated with
		* 						specific objects.
		* 				3. Static blocks:
		* 						Static statement blocks run once when the class is initialized.
		*
		*
		* 				4. Static inner class :
		* 						The non-static inner class depends on the instance of outer class. You need to create an
		* 						instance of outer class in order to create an instance of non-static inner class.
		* 						Static inner classes do not require above.
		*
		* 						StaticInnerClass cannot access non-static variables and methods of outer class.
		*
		* 				5. Static package import
		* 						When using static variables and methods, there is no need to specify the ClassName,
		* 						which simplifies the code, but the readability is greatly reduced.
		*
		* 						import static com.xxx.ClassName.*;
		*
		* 				6. Initialization Sequence
		* 						Static variables and static statement blocks take precedence over instance variables
		* 						and ordinary statement blocks, and the initialization order of static variables and
		* 						static statement blocks depend on their order in the code.
		*
		*						
		*
		*
		*
		*
		*
		*
		*
		*
		*
		*
		*
		 */

	}

}