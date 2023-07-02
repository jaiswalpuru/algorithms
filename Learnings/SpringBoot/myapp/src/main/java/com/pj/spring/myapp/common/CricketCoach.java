package com.pj.spring.myapp.common;

import jakarta.annotation.PostConstruct;
import jakarta.annotation.PreDestroy;
import org.springframework.beans.factory.config.ConfigurableBeanFactory;
import org.springframework.context.annotation.Scope;
import org.springframework.stereotype.Component;

//when a class is marked with Component annotation, then it is marked as spring bean

@Component
@Scope(ConfigurableBeanFactory.SCOPE_PROTOTYPE)
public class CricketCoach implements Coach {

	@PostConstruct
	public void doMyStartUpStuff() {
		System.out.println("In do My startupstuff () -> " + getClass().getSimpleName());
	}

	@PreDestroy
	public void doCleanUp() {
		System.out.println("cleaning up resources : " + getClass().getSimpleName());
	}

	public CricketCoach() {
		System.out.println("In cricket coach constructor : " + getClass().getSimpleName());
	}

	@Override
	public String dailyWorkout() {
		return "Getting work out done !!!!";
	}
}
