//refer geeks for geeks for Future and FutureTask

import java.util.concurrent.*;
import java.util.logging.Level;
import java.util.logging.Logger;

public class Main {
	public static void main(String[] args) {
		System.out.println("Hello world!");

		/*
		 *		Future and Future Task
		 *
		 *	A Future interface provides method to check if computation is complete, to wait for its computation
		 * 	to retrieve the results.
		 *
		 * 	using the get() method the results are retrieved.
		 *
		 * 	Future and FutureTask are both available in java.util.concurrent package
		 *
		 * 	FutureTask :
		 *
		 * 		1. FutureTask is an implementation of Future, Runnable, RunnableFuture interfaces and therefore
		 * 					can be submitted to Executor services.
		 *
		 * 		2. When calling ExecutorService.submit() -> on a Callable or Runnable instance, the ExecutorService
		 * 					returns a Future representing task, it can be created manually also
		 *
		 * 		3. FutureTask acts as similar to the CountDownLatch (this is used to make sure that a task wait for other
		 * 															task before it starts)
		 * 					when calling get() in that it waits for the task to complete or error out.
		 *
		 *
		 *		eg : When submitting a FutureTask instance to a thread pool(Executor service instance), it returns a
		 * 			 Future object. This object can be used for task completion and getting result asynchronously
		 * 		Submit task -------------> (task1) - (task2) ----------------------> Thread 1
		 * 									(Blocking queue) 			-----------> Thread 2
		 *
		 */

		MyRunnable mt1 = new MyRunnable(1000);
		MyRunnable mt2 = new MyRunnable(2000);
		FutureTask<String> ft1 = new FutureTask<>(mt1, "future task 1 is completed");
		FutureTask<String> ft2 = new FutureTask<>(mt2, "future task 2 is completed");

		// creating a thread pool of size 2
		ExecutorService es = Executors.newFixedThreadPool(2);

		// submit task 1
		es.submit(ft1);

		//submit task 2
		es.submit(ft2);

		/*
		* 	In the below example, After one task is completely executed, then after waiting 2000 ms second
		* 	task will be executed.
		 */

		while (true) {
			try {
				if (ft1.isDone() && ft2.isDone()) {
					System.out.println("Both returned from future.");
					es.shutdown();
					return;
				}
				if (ft1.isDone()) {
					System.out.println("Future task 1 is completed" + ft1.get());
				}
				System.out.println("Waiting for future task 2 to complete");
				String s = ft2.get(250, TimeUnit.MILLISECONDS);
				if (s != null) {
					System.out.println("Future taks 2 output : " + s);
				}
			} catch (Exception e) {
				e.printStackTrace();
			}
		}

	}
}

class MyRunnable implements Runnable {

	private final long waitTime;

	public MyRunnable(int milli) {
		this.waitTime = milli;
	}

	@Override
	public void run() {
		try {
			Thread.sleep(waitTime);
		} catch (InterruptedException ie) {
			Logger.getLogger(getClass().getSimpleName()).log(Level.SEVERE, null, ie);
		}
	}
}