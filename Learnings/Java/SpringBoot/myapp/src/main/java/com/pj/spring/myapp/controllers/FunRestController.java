package com.pj.spring.myapp.controllers;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class FunRestController {

	@Value("${coach.name}")
	private String name;

	//expose "/" to the world
	@GetMapping("/")
	public String sayHello() {
		return "Hello World!";
	}

	@GetMapping("/coachname")
	public String getCoachName() {
		return name;
	}
}

/*
* 	Spring Container works as object factory
* 	It has two primary functions :
* 	1. Create and manage objects (Inversion of control)
* 	2. Inject object dependencies (Dependency control)
 */


