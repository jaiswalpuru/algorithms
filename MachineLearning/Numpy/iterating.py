import numpy as np

np1 = np.array([1, 2, 3, 4, 5])
# for x in np1:
# 	print(x)
#
np2 = np.array([[1, 2, 3, 4, 5], [6, 7, 8, 9, 10]])
# for row in np2:
# 	for val in row:
# 		print(val)


np3 = np.array(
	[[[1, 2], [2, 3]], [[4, 5], [5, 6]]]
)
# print(np3.shape)
# for x in np3:
# 	print(x)
# for x in np3:
# 	for y in x:
# 		for z in y:
# 			print(z)

# using iterators
for x in np.nditer(np3):
	print(x)
