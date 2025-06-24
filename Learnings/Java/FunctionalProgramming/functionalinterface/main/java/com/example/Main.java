package com.example;

import java.util.Random;

public class Main {
	public static void main(String[] args) {
		// method 1
		Greeting greeting1 = new HelloGreeting();
		greeting1.sayHello();

		//method 2
		Greeting greeting2 = new Greeting() {
			@Override
			public void sayHello() {
				System.out.println("Hello World 2");
			}
		};
		greeting2.sayHello();

		//using lambda.
		Greeting greeting3 = () -> System.out.println("Hello World 3");
		greeting3.sayHello();

		Calculator c = (int a, int b) -> {
			Random r = new Random();
			return (a*b + r.nextInt(50));
		};
		System.out.println(c.calculate(1 ,2));

		IntBinaryOperator intBinaryOperator = (x, y) -> {
		  Random r = new Random();
		  int randomNumber = r.nextInt();
		  return (x*y) + randomNumber;
	  	};
	  	System.out.println(intBinaryOperator.applyAsInt(1, 2));
		
	}
}
