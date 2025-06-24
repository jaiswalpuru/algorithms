package com.pj.spring.myapp.controllers;

import com.pj.spring.myapp.common.Coach;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class DemoController {
	private Coach myCoach;
	private Coach myCoach2;
	private Coach myCoach3;
	private Coach myCoach4;
	// if there is only one constructor, then the @Autowired annotation is optional

	//constructor injection
	//	@Autowired
	//	public DemoController(Coach coach) {
	//		this.myCoach = coach;
	//	}

	/*
		*Below is the example of field injection, but it is not recommended by spring.io team.
		* @Autowired
		* private Coach getMyCoach; //now for this we don't setters or getters.
	*/

//	@Autowired
	public DemoController(@Qualifier("baseBall") Coach c, @Qualifier("baseBall") Coach c2,
						  @Qualifier("cricketCoach") Coach c3, @Qualifier("cricketCoach") Coach c4) { // here c and c2 both point to the same instance
		//	as only one instance is created by spring for each bean
		this.myCoach = c;
		this.myCoach2 = c2;
		this.myCoach3 = c3;
		this.myCoach4 = c4;
	}

	@Autowired
	public DemoController(@Qualifier("her") Coach c) {
		System.out.println("In constructor : " + getClass().getSimpleName());
		myCoach = c;
	}


	@GetMapping("/check")
	public String check() {
		return "comparing if two beans are equal : " + (myCoach == myCoach2) +
			" another two comparison for which I used the prototype " + (myCoach3 == myCoach4);

	}


	// example or setter injection
	/*	@Autowired
	//	public void setMyCoach(Coach c) {
	//		this.myCoach = c;
	}*/

	@GetMapping("/dailyworkout")
	public String getDailyWorkout(){
		return myCoach.dailyWorkout();
	}
}


/*
 * So lets say if the Coach is implemented by many other classes so how will spring know which one to call
 * Here, the use of qualifier will come in.
 * Check demo controller for @Qualifier annotation
*/

/*
*	There is another alternate option for Qualifier -> Primary
* 	If we don't care which coach is returning the result.
*	The User has to figure whom to make the primary
* 	When using primary there is no need to use Qualifier
*
* 	One caveat with Primary
* 	We can only mark one bean as primary
* 	If we mark multiple bean as primary as it will throw an error
*/

/*
*		If we mix @Primary and @Qualifier, then Qualifier will have
* 		higher priority.
*
* 	In general, it is recommended to use Qualifier
* 	1. More specific
* 	2. Higher Priority
*/

/*
*		Lazy Initialization
* 		By default when the application start all beans are initialized
* 		@Component etc...
*
*		Instead of creating all beans up front we can specify lazy initialization.
* 		A bean will only be initialized in the following cases :
*
* 		1. It is needed for dependency injection.
* 		2. Or it is explicitly requested.
* 		3. We just need to add the @Lazy annotation to the given class.
*
* 		rather than going to each bean and setting @Lazy, this can be done in application.properties files
*
* 		# this specifies all beans are lazy and no bean is created until it is required.
*		spring.main.lazy-initialization=true
*
* 		Advantages :
*
* 			1. Only create objects when needed.
* 			2. May help with faster startup time if you have large number of components.
*
* 		Disadvantages :
*
* 			1. If you web related components like @RestController, not created until requested.
* 			2. May not discover configuration issues until too late.
* 			3. Need to make sure you have enough memory for all beans once created.
*
*/

/*
*			Beans scope
* 		- Scope refers to the lifecycle of the bean.
* 		- How long does the bean live
* 		- How is the bean shared.
* 		- How many instances are created.
*
* 		Default scope is Singleton
* 			- Only one instance of bean is created.
* 			- It is cached in memory
* 			- All dependency injection for the bean
* 					- will reference the same bean.
*
* 		Can explicitly specify the bean scope.
* 			- @Scope(ConfigurableBeanFactory.SCOPE_SINGLETON)
*
*		singleton -> create a single shared instance of the bean. Default scope.
* 		prototype -> Creates a new bean for each container request
* 		request -> Scoped to an HTTP request
* 		session -> Scoped to an HTTP web session. (only used for web apps)
*		global-session -> scoped to an HTTP web session. Only used for web apps.
*
 */

/*
*							Bean LifeCycle
* 	Container started -> Bean Instantiated -> Dependencies injected -> Internal Spring processing -> your spring method
* 						Bean is ready for use, once the container is shutdown, your custom gets destroyed (STOP)
*
* 	Bean Lifecycle methods or hooks
* 			- For adding custom code during bean initialization.
* 			- Calling custom business logic method
* 			- Setting up handles to resources (db, sockets, file ...) etc
*
* 	Can also add custom code during bean destruction.
* 		- Calling custom business logic method.
* 		- Clean up handles to resources (db,sockets, ...)
*
* 	Annotation for doing stuff Post startup
* 	 		- @PostConstruct
*
* 	Pre Destruct
* 			- @PreDestroy
*
* 		Few Notes :
* 				1. For prototype scoped beans spring does not call destroy method.
* 				2. Prototype beans are lazy by default.
*
* 	Java Config Bean
* 		1. Create @Configuration class.
* 		2. Define @Bean method to configure the bean.
* 		3. Inject the bean into our controller.
*
*		Use case for @Bean
* 				1. Make an existing third-party class available to spring framework.
* 				2. You may not have access to the source code of third party class.
* 				3. if want to use third party class as SpringBean
*
* 		A custom bean ID can also be provided
* 		@Bean("id name") and then that id can be used as Qualifier
*
*/


/*
*								Hibernate and JPA
*
* 		Hibernate -> A framework for persisting/ storing Java objects into a database.
* 									High Level
* 				Java App -------> Hibernate ----------> DB
*				hibernate handles low level SQL code.
*
* 	It minimizes the amount of JDBC code you have to develop
*
* 	Hibernate provides the Object to Relational Mapping (ORM)
*
* 	Developer defines the mapping of Java class and database table
*
* 	JPA -> Jakarta persistence API. -- previously known as Java persistence API
* 	- Standard API for ORM.
*
* 			JPA Vendor Implementations
* 				/				\
* 			  /					 \
* 			Hibernate			EclipseLink
*
* 		Hibernate is the most popular implementation for JPA, and is also the default implementation for springboot.
*
* 	- By having a standard API, you are not locked to vendors implementation.
* 	- Maintain portable code.
* 	- Can theoretically switch vendor implementations.
*
* 	Saving Java object with JPA
*
* 	// create the java object
* 	Student s = new Student("asb", "23", "name@email.com");
*
* 	// save it to the database
* 	entityManager.persist(s);
*
* 	//Retrieve from database
* 	int id = 1;
* 	Student data = entityManager.find(Student.class, id);
*
* 	Querying java objects
*
* 	1. Let's say we need to retrieve list of students
* 		TypedQuery<Student> theQuery = entityManger.createQuery("from Student", Student.class);
* 		List<Student> students = theQuery.getResultList();
*
*
* 					Your Java 		----->  JPA hibernate ----> DB
* 					application		<-----				  <-----
*
* 		Create objects
* 		Read Objects
* 		Update Objects
* 		Delete Objects
*
*
* 		How does hibernate JPA relates to JDBC.
* 			Hibernate JPA uses JDBC for all database communications.
* 			Hibernate JPA is another layer of abstraction on top of JDBC.
*
*
*
* 		In SpringBoot, Hibernate is the default implementation of JPA.
* 		EntityManager is main component for creating queries etc.....
* 		EntityManager is from Jakarta Persistence API.
*
* 		Based on configs, Spring Boot will automatically create the beans :
* 			- Data Source, EntityManager
*
* 		You can then inject these into your app, for example your DAO
*
* 		Add dependencies :
* 			1. MySQL Driver : mysql-connector-j
* 			2. Spring Data JPA : spring-boot-started-jpa
*
* 		And we store the DB connection information in application.properties file.
*
* 		application.properties
* 		a) spring.datasource.url = jdbc:mysql://localhost:3306/student_tracker
* 		b) spring.datasource.username = springstudent
* 		c) spring.datasource.password = springstudent
*
* 		Note : no need to give JDBC driver class name Spring Boot, will automatically detect it based on the URL.
*
* 		For example refer to different projects in SpringBoot.
* 				1) command line application
*					Rest of the notes will be continued in the projects.
*
*
*
*
*/

