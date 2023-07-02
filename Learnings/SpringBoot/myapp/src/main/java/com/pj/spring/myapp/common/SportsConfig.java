package com.pj.spring.myapp.common;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class SportsConfig {
	@Bean("aquatic")
	public Coach swimCoach() {
		return new SwimCoach();
	}

	@Bean("her")
	public Coach swim() {
		return new SwimCoach();
	}
}
