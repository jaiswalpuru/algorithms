import time
from multiprocessing import Pool

def f(n):
	sum = 0
	for x in range(1000):
		sum += x**2
	return sum

if __name__ == "__main__":
	p = Pool()
	t1 = time.time()
	res = p.map(f, range(10000))
	p.close()
	p.join()
	print("Pool took : ", time.time()-t1)

	t2 = time.time()
	res = []
	for x in range(10000):
		res.append(f(x))
	print("Pool took : ", time.time()-t2)

	p2 = Pool(processes=3)	# this will create only three processes
	result = p2.map(f, [1, 2, 3, 4, 5])
	for x in result:
		print(x)

