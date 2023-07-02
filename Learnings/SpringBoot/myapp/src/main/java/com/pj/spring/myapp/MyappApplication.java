package com.pj.spring.myapp;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;


/*
*		Spring by default scans everything in com.pj.spring.myapp
* 		By default spring will not scan anything inside util package
* 		For this we will need to tell spring explicitly to scan which packages
 */

@SpringBootApplication( // explicitly telling spring boot which packages to scan
	scanBasePackages = {"com.pj.spring.myapp",
							"com.pj.spring.util"}
)
public class MyappApplication {
	public static void main(String[] args) {
		SpringApplication.run(MyappApplication.class, args);
	}
}

/*
Field injection is not recommended by spring.io team
 */