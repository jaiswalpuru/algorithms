package com.pj.spring.myapp.common;

public interface Coach {
	String dailyWorkout();
	default void d() {
		System.out.println("Here I am from Coach interface");
		p();
	}

	private void p(){
		System.out.println("private function inside an interface");
	}
}
