package com.pj.spring.myapp.common;

import jakarta.annotation.PostConstruct;
import jakarta.annotation.PreDestroy;
import org.springframework.context.annotation.Lazy;
import org.springframework.stereotype.Component;

@Component
@Lazy // now this bean will only be initialized when it is needed for dependency injection.
public class TrackCoach implements Coach{

	@PostConstruct
	public void doMyStartUpStuff() {
		System.out.println("In do My startupstuff () -> " + getClass().getSimpleName());
	}

	@PreDestroy
	public void doCleanUp() {
		System.out.println("cleaning up resources : " + getClass().getSimpleName());
	}

	public TrackCoach() {
		System.out.println("Inside track coach : " + getClass().getSimpleName());
	}

	@Override
	public String dailyWorkout() {
		return "practive runnig 5K";
	}
}
