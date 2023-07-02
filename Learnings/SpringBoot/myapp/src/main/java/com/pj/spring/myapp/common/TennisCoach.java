package com.pj.spring.myapp.common;

import jakarta.annotation.PostConstruct;
import jakarta.annotation.PreDestroy;
import org.springframework.context.annotation.Primary;
import org.springframework.stereotype.Component;

@Component
//@Primary
public class TennisCoach implements Coach{

	@PostConstruct
	public void doMyStartUpStuff() {
		System.out.println("In do My startupstuff () -> " + getClass().getSimpleName());
	}

	@PreDestroy
	public void doCleanUp() {
		System.out.println("cleaning up resources : " + getClass().getSimpleName());
	}

	public TennisCoach() {
		System.out.println("In cricket coach constructor : " + getClass().getSimpleName());
	}
	@Override
	public String dailyWorkout() {
		return "practice swinging";
	}
}
