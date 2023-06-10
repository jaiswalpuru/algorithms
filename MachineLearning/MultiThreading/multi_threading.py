import time
import threading


def square(l):
	print("square of numbers")
	for x in l:
		time.sleep(0.2)
		print("square : ", x ** 2);


def cube(l):
	print("cube of numbers")
	for x in l:
		time.sleep(0.2)
		print("cube : ", x ** 3)


arr = [2, 3, 4, 5, 6]

t = time.time()
# square(arr)
# cube(arr)

t1 = threading.Thread(target=square, args=(arr,))
t2 = threading.Thread(target=cube, args=(arr,))

t1.start()
t2.start()

t1.join()
t2.join()

print("total time : ", time.time()-t)

