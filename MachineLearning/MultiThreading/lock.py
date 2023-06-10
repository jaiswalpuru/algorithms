"""
To protect shared resources memory, files, databases, from multiple access we use locks
Whatever part of code which is being shared is called the critical section
"""

import time
import multiprocessing


def deposit(balance):
	for i in range(100):
		time.sleep(0.01)
		balance.value = balance.value + 1


def withdraw(balance):
	for i in range(100):
		time.sleep(0.01)
		balance.value = balance.value - 1


def deposit2(balance:int, lock:multiprocessing.Lock):
	for i in range(100):
		time.sleep(0.01)
		lock.acquire()
		balance.value = balance.value + 1
		lock.release()


def withdraw2(balance:int, lock:multiprocessing.Lock):
	for i in range(100):
		time.sleep(0.01)
		lock.acquire()
		balance.value = balance.value - 1
		lock.release()


if __name__ == "__main__":
	balance = multiprocessing.Value('i', 200)
	d = multiprocessing.Process(target=deposit, args=(balance,))
	w = multiprocessing.Process(target=withdraw, args=(balance,))
	d.start()
	w.start()
	d.join()
	w.join()
	print(balance.value)

	# now we see the problem above now below is the solution using locks
	balance2 = multiprocessing.Value('i', 200)
	lock = multiprocessing.Lock()
	d2 = multiprocessing.Process(target=deposit2, args=(balance2, lock,))
	w2 = multiprocessing.Process(target=withdraw2, args=(balance2, lock,))
	d2.start()
	w2.start()
	d2.join()
	w2.join()
	print(balance2.value)
