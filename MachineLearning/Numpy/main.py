import numpy as np

'''
Anything in side a numpy array has to be same data type.
ndArray -> data type for numpy is ndArray -> n dimensional array
'''

nd1 = np.array([1, 0, 2, 3, 4, 5, 6, 76, 7, 9])
# print the array
print(nd1)

# shape or basically the length of nd array
print(nd1.shape)

nd2 = np.arange(10)
print(nd2)

# using step function
nd3 = np.arange(1, 10, 2)
print(nd3)

# arrays of zeroes
nd4 = np.zeros(10)
print(nd4)

# multi dimensional arrays (zeroes)
nd5 = np.zeros((2, 10))
print(nd5)

nd6 = np.full((10), 6)
print(nd6)

nd7 = np.full((2, 10), 6)
print(nd7)

# Convert python list to numpy
my_list = [1, 2, 3, 4, 5]
nd8 = np.array(my_list)
print(nd8)

print(nd8[4])

print(nd7[1])

############################################################################################################

# --------------------------------Slicing numpy array--------------------------------

np1 = np.array([1, 2, 3, 4, 5, 6, 7, 8, 9])

# return [2, 3, 4, 5] slice
print("Here : ", np1[1:3])

print("till the end : ", np1[3:])

# print from back
# not including last element
print(np1[-3: -1])

# steps
print(np1[1:5:2])

# steps on entire array
print(np1[::2])

np2 = np.array([[1, 2, 3], [4, 5, 6], [7, 8, 9]])
print(np2[2, 1])

# slice a 2d array
print(np2[0:1, 1:2])


############################################################################################################


# 									Universal functions
np3 = np.array([1, 1, 2, 3, 4, 5, 6, 7, 8, 9])
print(np3)

# square root of each element
print(np.sqrt(np3))

print(np.absolute(np3))

# exponential functions
print(np.exp(np3))

# print min max
print(np.max(np3))
print(np.min(np3))

# sign positive or negative
print(np.sign(np3))

print(np.sin(np3))

print(np.log(np3))


