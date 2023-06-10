import time
import multiprocessing

square_res = []


def square(arr):
	global square_res
	print("square of numbers")
	for x in arr:
		# time.sleep(5)
		print("square : ", x ** 2)
		square_res.append(x ** 2)
	print("within process : ", square_res)


def cube(arr):
	print("cube of numbers")
	for x in arr:
		time.sleep(5)
		print("cube : ", x ** 3)


if __name__ == "__main__":
	arr = [2, 3, 4, 5]
	p1 = multiprocessing.Process(target=square, args=(arr,))
	# p2 = multiprocessing.Process(target=cube, args=(arr,))

	p1.start()
	# p2.start()

	p1.join()
	# p2.join()

	print(square_res) # the output here will be null, because every process has its own address space so the
	# square_res to which we are appending will get updated in the address space of p1
	# if we had done this using multi threading, then this would have been updated as thread shares spaces in a process.

'''
difference between threading and multiprocessing :
	Both are the ways to achieve multi tasking : doing multiple task at the same time. 
	Multiple threads lives with in the same process. 
	Threads will share the address space, where as the different process will have their own address space.
	
	The benefit of multiprocessing is that error or memory leak in one process won't affect the execution of other
	processes.
'''
