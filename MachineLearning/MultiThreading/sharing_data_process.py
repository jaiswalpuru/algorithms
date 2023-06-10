import multiprocessing

res = []


def square(arr):
	global res
	for x in arr:
		res.append(x ** 2)
	print(f'within the process {res}')


def square_shared(arr, result):
	for idx, n in enumerate(arr):
		result[idx] = n ** 2


def value_example(numbers, result, val):
	print(numbers)
	print(result[:])
	val.value = 5.67


if __name__ == "__main__":
	numbers = [2, 3, 5]
	p = multiprocessing.Process(target=square, args=(numbers,))
	p.start()
	p.join()
	print(f'outside the process {res}')  # this will be null

	# to resolve the above issue we will create a shared array and pass it to the function
	result = multiprocessing.Array('i', 3)
	p2 = multiprocessing.Process(target=square_shared, args=(numbers, result,))
	p2.start()
	p2.join()
	print(f'outside {result[:]}')

	# value is a single value which is shared among processes
	value = multiprocessing.Value('d', 0.0)  # d-> double
	p3 = multiprocessing.Process(target=value_example, args=(numbers, result, value,))
	p3.start()
	p3.join()
	print(value.value)

# to share data between two processes we need to use shared memory
# there are two ways of sharing data
# 	1. Value
#	2. Array
