package com.pj.spring.myapp.common;

import org.springframework.beans.factory.config.ConfigurableBeanFactory;
import org.springframework.context.annotation.Scope;
import org.springframework.stereotype.Component;

public class SwimCoach implements Coach {

	public SwimCoach() {
		System.out.println("Inside the swim coach constructor : " + getClass().getSimpleName());
	}

	@Override
	public String dailyWorkout() {
		return "Swim daily";
	}

}
