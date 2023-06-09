import numpy as np

# Copy vs View
np1 = np.array([0, 1, 2, 3, 4, 5])

# create a view
np2 = np1.view()

print(f'Original np1 {np1}')
print(f'Original np2 {np2}')

np1[0] = -1
print (f'Original array {np1}')
print(f'Original array {np2}')

# view is basically a shallow copy in reference to Java
# where as copy is deep copy

np3 = np1.copy()
print(np3)
np3[0] = 10
print(np3)
print(np1)
