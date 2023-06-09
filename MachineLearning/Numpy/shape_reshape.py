import numpy as np

# create 1-D numpy array and get shape
np1 = np.array([1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12])
print(np1)
print(f'shape : {np1.shape}')

# 2d array using numpy
np2 = np.array([[1, 2], [3, 4], [4, 5]])
print(np2.shape)

# reshape
np3 = np1.reshape(3, 4)
print(np3)
print(f'np3 shape : {np3.shape}')

np4 = np1.reshape(2, 3, 2)
print(np4)

# flatten the array to 1-D
np5 = np4.reshape(-1)
print(np5)
