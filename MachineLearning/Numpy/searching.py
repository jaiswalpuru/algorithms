import numpy as np

np1 = np.array([1, 2, 3, 4, 3, 6, 8, 9])

# find all the 3's in array, this returns a tuple of index
x = np.where(np1 == 3)
print(x)

# return even items
y = np.where(np1 % 2 == 0)
print(y)


