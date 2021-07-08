### DP

```sh

Top Down := We write code similar to recursive solution, but check before each calculation whether the result has
already been stored in a cache. This is also called memoization.

Bottom Up := Methodically builds up value for F(1), F(2) .... one after the other, typically by adding values to an
array or dictionary.

1. Identify the recurrence relation : how can the problem be broken down into smaller parts.
2. Initialize a cache capable of storing the values for each subproblem.
3. Create a memoized function (if top down) or loop (if bottom up) which populates the cache values.
```