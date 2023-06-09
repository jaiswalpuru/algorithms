import numpy as np

# filtering numpy arrays using boolean index lists
np1 = np.array([i for i in range(1, 11)])

# x = [True, True, False, False, False, False, False, False, False, False]

# print(np1)
# print(np1[x])

'''
filtered = []
for x in np1:
	if x % 2 == 0:
		filtered.append(True)
	else:
		filtered.append(False)

print(filtered)
print(np1[filtered])
'''

filtered = np1 % 2 == 0
print(filtered)
print(np1[filtered])
