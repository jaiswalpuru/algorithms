import multiprocessing

res = []


def square(arr, q):
	for x in arr:
		q.put(x ** 2)


if __name__ == "__main__":
	number = [2, 3, 5]
	q = multiprocessing.Queue()
	p = multiprocessing.Process(target=square, args=(number, q,))

	p.start()
	p.join()

	while not q.empty():
		print(q.get())

'''
				Multiprocessing queue                   Queue Module
				1. import multiprocessing				import queue
				2. q = multiprocessing.Queue()			q = queue.Queue()
				3. Lives in shared memory				Lives in in-process memory
		4. Used to share data between processes.		Used to share data between threads
'''
