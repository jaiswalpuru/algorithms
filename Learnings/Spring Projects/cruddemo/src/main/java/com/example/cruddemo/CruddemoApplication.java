package com.example.cruddemo;

import com.example.cruddemo.dao.StudentDAO;
import com.example.cruddemo.entity.Student;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;

@SpringBootApplication
public class CruddemoApplication {

	public static void main(String[] args) {
		SpringApplication.run(CruddemoApplication.class, args);
	}

	/*
	 *	Command line runner is executed after the spring beans have been loaded into the application context.
	 * 	The below bean is for creating a command line application.
	 */

	@Bean
	public CommandLineRunner commandLineRunner(StudentDAO studentDAO) {
		return runner -> {
			//createStudent(studentDAO);
			createMultipleStudent(studentDAO);
		};
	}

	private void createMultipleStudent(StudentDAO studentDAO) {
		System.out.println("Creating multiple student");
		String[] fName = {"Puru", "Neeraj", "Ravi", "Dijkstra", "Cormen", "Gilbert"};
		String[] lName = {"Jaiswal", "Mittal", "Prakash", "Edgar", "thomas", "Strang"};
		String[] email = {"pj@gmail.com", "nee@gmail.com", "ravi@gmail.com", "edgar@gmail.com", "thomas@gmail.com", "gil@gmail.com"};

		for (int i=0; i<fName.length; i++) {
			Student student = new Student(fName[i], lName[i], email[i]);
			studentDAO.save(student);
		}
	}

	private void createStudent(StudentDAO studentDAO) {
		// create the student object
		System.out.println("creating a new student object");
		Student student = new Student("Paul", "Doe", "pauldoe@gmail.com");
		//save the student
		System.out.println("Saving the student");
		studentDAO.save(student);
		//display id of the saved student
		System.out.println("Generated id : " + student.getId());
	}


}


/*
*	whenever we run the application we see the SpringBoot banner, to shut it down :
* 		application.properties
* 				Add spring.main.banner-mode=off
*
* 	To reduce the background information, these can also be hidden by adding
* 				logging.level.root=warn
*
* 	If there is an exception in spring boot application, spring will still throw exception.
*
*
*
* 							JPA Development process
* 			1. Annotate the Java class
* 			2. Develop Java code to perform database operations.
*
* 			As previously mentioned Hibernate is the default JPA implementation of spring boot.
*
* 			Entity Class : Java class that is mapped to a database table.
*			ORM (Object Relational mapping) :
*
* 				Java class ----------------> JPA -----------------------> DB
* 							<--------------		 <-----------------------
*
*
* 			Minimum requirements for Entity class :
* 				1. Must be annotated with @Entity
* 				2. Must have a public or protected no-arg constructor
* 				3. the class can have other constructors as well.
* 				4. In Java when we do not declare any constructor then a no - arg constructor is provided by Java
* 					but in case where we declare a constructor with arguments then a no-arg constructor is not
* 					provided by default.
*
*
*
* 			Java annotations
* 				1. Map class to database table
* 				2. Map fields to database columns
*
* 			@Column annotation is optional we don't need to explicitly provide it, column name is same as the Java
* 			field name. But in general it is preferable to not follow this approach, it might break the code in,
* 			future.
*
* 			@GeneratedValue(strategy=GenerationType.IDENTITY) -> this means, our code need not worry about this field
* 																 it will be handled by the database not code.
*
* 			 GenerationType.AUTO ---> Pick an appropriate strategy for the particular database.
* 			 GenerationType.IDENTITY -----> Assign primary keys using database identity column.
* 			 GenerationType.SEQUENCE ------> Assign primary keys using a database sequence.
* 			 GenerationType.TABLE ------> Assign primary keys using an underlying database table to ensure uniqueness.
*
*
* 			In general GenerationType.IDENTITY annotation is recommended.
*
*
* 			• You can define your own CUSTOM generation strategy.
* 			• Create implementation of org.hibernate.id.IdentifierGenerator
* 			• Override method public Serializable generate(...)
*
*
* 				Saving a Java Object
*
* 		1. Create a new Student
* 		2. Read a Student
* 		3. Update a Student
* 		4. Delete a Student
*
*		Student Data Access Object
* 	This is responsible for interfacing with the database.
* 	This is a common design pattern : Data Access Object (DAO) -> this is like a helper class for communicating with DB
* 	Crud Demo App <--------> DAO <-------> DB
*
*
* 	Student Data Access Object :
* 		Our DAO needs JPA entity manager
* 		JPA entity manager is the main component for saving/ retrieving entities.
*
* 		Student DAO <------> Entity Manger <--------> ..... <--------> DB
*
*
* 	JPA Entity Manger:
* 		1. Our JPA entity manager needs a data source.
* 		2. The datasource defines the database connection info
* 		3. JPA entity manager and data source are automatically created by springboot.
*
* 				Based on the info from maven pom file and application.properties
*
* 		4. We can autowire or inject JPA entity Manager info into our Student DAO.
*
* 		Student DAO <------> Entity Manger <--------> Data source <--------> DB
*
* 		Student DAO - Steps to create
* 				1. Define DAO interface
* 				2. Define DAO implementation
* 					• Inject the entity manager.
* 				3. Update the main app.
*
*
* 		@Transactional annotation
* 			1. Automatically begin and end transaction for your JPA code
* 				• No need to explicitly do this in code
*
*		Where to use the transactional annotation, when we save something into a database we want it to run
* 		inside a transaction.
*
* 		Specialized Annotation for DAO
* 			• Spring provides @Repository annotation
*
*	   				@Component
* 		@RestController   @Repository  ....
*
*
*	@Repository annotation :
* 		1. Supports component scanning
* 		2. Translates JDBC exceptions
*
*
* 	How to change the AUTO Increment values
* 		 ALTER TABLE student_tracker.student AUTO_INCREMENT=3000;
*
* 	How to reset AUTO_INCREMENT value to 1.
* 		TRUNCATE student_tracker.student; -> this query will reset all the data present in student table.
*
*
*
* 						READING OBJECTS FROM JPA
*
*
*
*
 */