package com.pj.spring.myapp.common;

import jakarta.annotation.PostConstruct;
import jakarta.annotation.PreDestroy;
import org.springframework.stereotype.Component;

@Component
public class BaseBall implements Coach{

	@PostConstruct
	public void doMyStartUpStuff() {
		System.out.println("In do My startup stuff () -> " + getClass().getSimpleName());
	}

	@PreDestroy
	public void doCleanUp() {
		System.out.println("cleaning up resources : " + getClass().getSimpleName());
	}

	public BaseBall() {
		System.out.println("In cricket coach constructor : " + getClass().getSimpleName());
	}


	@Override
	public String dailyWorkout() {
		return "Spend 30 minutes in batting practice.";
	}
}
