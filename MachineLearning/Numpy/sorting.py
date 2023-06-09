import numpy as np

np1 = np.array([10, 9, 8, 71000, 123, 89898])
np1.sort()
print(np1)

np2 = np.array(["John", "Tina", "Aaron"])
print(np.sort(np2))
print(f'Original : ', np2)

# 2-d
np3 = np.array([[1, 2, 3, 4, 5, 6, 7, 8, 9], [9, 8, 7, 6, 5, 4, 3, 2, 1]])
# sorting on 2-d array will sort each row not combined
print(np.sort(np3))



