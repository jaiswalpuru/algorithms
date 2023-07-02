// refer : https://github.com/CyC2018/CS-Notes/blob/master/notes/Java%20%E5%B9%B6%E5%8F%91.md#%E4%B8%80%E4%BD%BF%E7%94%A8%E7%BA%BF%E7%A8%8B

import java.util.concurrent.*;

public class Main {
	public static void main(String[] args) throws ExecutionException, InterruptedException {
		System.out.println("Hello world!");

		/*
		*					Java threads
		* 		There are three ways to use threads.
		* 		1. Implement runnable interface.
		*		2. Implement callable interface.
		* 		3. Extends the thread class.
		*
		* 	Class implementing the Runnable and Callable interfaces can only be regarded as a task that can run in thread.
		* 	In the end it needs to be called through thread.
		*  	eg : @1
		*
		 */

		// using runnable
		MyRunnable instance = new MyRunnable();
		Thread t = new Thread(instance);
		t.start();
		/*
		*									Using callable
		* Compared to runnable, Callable can have a return value, and the return value is encapsulated by
		* Future task.
		* eg : @2
		 */
		MyCallable mc = new MyCallable();
		FutureTask<Integer> ft = new FutureTask<>(mc);
		Thread t2 = new Thread(ft);
		t2.start();
		System.out.println(ft.get());

		/*
		*		Inherit the thread class
		* 	eg : @3
		 */
		MyThread mt = new MyThread();
		mt.start();

		/*
		 *	Why implementing interface is better ?
		 * 	because java does not support multiple inheritance.
		 * 	and, Thread class is too large.
		 */


		/*
		 *		Basic thread mechanism
		 *	A thread pool is a collection of threads which can be used to execute tasks concurrently, these are created in advance
		 * 	and store in a pool, (like queue). When needed the threads are taken out from pool and assigned to tasks.
		 *
		 * Executor : An executor manages the execution of multiple asynchronous tasks without requiring the programmer
		 * 			  to explicitly manage the lifecycle of thread.
		 * 			  The asynchrony here means that thread executes without interfering each other.
		 *
		 *	There are three main executors :
		 * 		1. CachedThreadPool : A task is created by thread.
		 * 		2. FixedThreadPool : All tasks can only use fixed size thread.
		 * 		3. SingleThreadExecutor : Equivalent to FixedThreadPool with size 1.
		 *
		 * 	1. CachedThreadPool : a task creates new thread as required, but will reuse previously constructed thread
		 * 						as required.
		 * 						When a task is submitted to a cachedThreadPool it will first check if there are any idle
		 * 						thread, if there is then the task will be assigned to those threads first, else pool will
		 * 						create a new thread to execute the task.
		 *	 					Once the thread has finished executing a task, it will be placed back in pool.
		 *						If the pool has more threads than it needs, then rest will be terminated.
		 *
		 *		Pros : 1) improved performance, 2) reduce resource usage
		 * 		Cons : 1) Increased memory usage 2) increased CPU usage.
		 * 		These threads are not suitable for running for long awaiting tasks, as they terminate idle threads.
		 *
		 */

		ExecutorService cached = Executors.newCachedThreadPool();
		ExecutorService fixed = Executors.newFixedThreadPool(2);
		ExecutorService single = Executors.newSingleThreadExecutor();
		for (int i=0; i<5; i++) {
			cached.execute(new MyRunnable());
		}
		cached.shutdown();
		fixed.shutdown();
		single.shutdown();

		/*
		 *		Daemon
		 * 	A daemon thread is a thread that provides services in background when program is running and is not an integral
		 * part of program.
		 *
		 * Main belongs to a non-daemon thread.
		 *
		 */

		Thread t3 = new Thread(new MyRunnable());
		t3.setDaemon(true);

		/*
		 * sleep() -> Will sleep the currently executing thread.
		 * @throws InterruptedException
		 * @4
		 */
		Thread t4 = new Thread(new Sleep());
		t4.start();
		System.out.println("After sleeping");

		/*
		 *		yield() -> static method present in Thread class
		 * 				Declares that this thread has completed most important part of lifecycle and can be switched
		 * 				to other threads for execution.
		 *		eg : @4
		 *
		 */

		/*
		*				Interruption
		*		A thread might end early if there is an interruption.
		*
		* 		Interrupt the thread by calling interrupt() method. If a thread is blocked or waiting for a long time
		* 		an InterruptedException will be thrown, to end the thread early.
		*
		* 		**** Note : I/O blocking and synchronized blocking cannot be interrupted.
		*		@5 -> example for interrupted thread
		*
		 */

		Interruptable.Thread1 t1 = new Interruptable.Thread1();
		t1.start();
		t1.interrupt();
		System.out.println("Main run");

		/*
		 *	If a thread's run method executes infinite loop and does not perform operations such as sleep() that
		 * 	would throw an interrupted exception.
		 *
		 * Calling the interrupt() method will set the thread's interrupt flag, and interrupted() will return true;
		 *
		 */

		Interruptable2.Thread2 t6 = new Interruptable2.Thread2();
		t6.start();
		t6.interrupt();
		System.out.println("Main run");

		/*
		*		Interrupt operation on Executor
		*
		* 	Calling the shutdown method on Executor will wait for the threads to finish executing before closing,
		* 	but if shutdownNow() is called then it is equivalent to interrupt() method of each thread.
		 */

		ExecutorService executorService = Executors.newCachedThreadPool();
		executorService.execute(() -> {
			try {
				Thread.sleep(2000);
				System.out.println("Thread run....");
			} catch (InterruptedException ex) {
				ex.printStackTrace();
			}
		});
		executorService.shutdownNow();
		System.out.println("Main exited");

		/*
		 *	If only want to interrupt the thread in the Executor, you can submit a thread by using the submit() method
		 * 	which will return Future<?> object and you can interrupt the thread by calling the object's cancel(true) method.
		 *
		 * 	eg : Future<?> future = executorService.submit(() -> { ... });
		 * 	future.cancel(true);
		 */




	}
}

class Interruptable2 {
	static class Thread2 extends Thread {
		@Override
		public void run() {
			while (!interrupted()) {
				System.out.println("Exited");
			}
		}
	}
}

class Interruptable {
	static class Thread1 extends Thread {
		@Override
		public void run() {
			try {
				Thread.sleep(2000);
				System.out.println("Thread run");
			} catch (InterruptedException ex) {
				ex.printStackTrace();
			}
		}
	}
}

//@4
class Sleep implements Runnable {
	@Override
	public void run() {
		try {
			Thread.sleep(5000);
			Thread.yield();
			System.out.println("Yield called");
		} catch (InterruptedException ie){
			ie.printStackTrace();
		}
	}
}

//@3
class MyThread extends Thread {
	@Override
	public void run() {
		System.out.println("extending the thread class");
	}

}

//@2
class MyCallable implements Callable<Integer> {
	public Integer call() {
		return 123;
	}
}

// @1
class MyRunnable implements Runnable {
	static int i = 0;

	@Override
	public void run() {
		System.out.println("I implemented the runnable interface " + i++);
	}
}
