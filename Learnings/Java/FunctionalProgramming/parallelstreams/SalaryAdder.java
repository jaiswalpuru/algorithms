package parallelstreams;

import java.util.concurrent.atomic.LongAdder;

public class SalaryAdder {
	int total;
	LongAdder totalOne;
//	public void accept(Employee e) {
//		total += e.getSalary();
//	}

	public SalaryAdder() {
		total = 0;
		totalOne = new LongAdder();
	}
	//to overcome race condition
	public synchronized void accept(Employee e) {
		total += e.getSalary();
	}

	//the above method is slow to overcome that
	// we use LongAdder
	public void acceptTwo(Employee e) {
		totalOne.add(e.getSalary());
	}
}
