### Simple Algorithm
```sh
Algorithm R -> Alan Waterman (Slow)
The algorithm works by maintaing a reservoir of size k, which initially contains the first k items of the input.

It then iterates over the remaining items until the input is exhausted.
Using one-based array indexing  let i > k be the index of the current item under consideration.

The algo then generates a random number j, b/t [1, i]. If j is atmost k, then the item is selected and 
replaces whichever item currently occupies the j-th position in the reservoir.

Else the item is discarded. For all i, the ith element chosen to be included in the 
reservoir with probability of k/i. 

Similarly, at each iteration the jth element of the reservoir array is chosen to be replaced with 
probability 1/k X k/i = 1/i. 

It can be shown that when the algorithm has finished executing, each item in the input population has equal 
probability (k/n).
```

```sh
//s -> items to be sampled
//r -> will have the result
func reservoir_sampling_algo_r(s []int, r []int) {
    //fill the reserveoir array with the first k elements acc to the algorithm
    for i:=0; i<k; i++ {
        r[i] = s[i]
    }
    for i := k+1; i<n; i++ {
        j := rand.Intn(1,i+1) // generate random integer inclusive of 
        if j <= k { //if j is atmost k
            //then replace the element in the reservoir with the element under consideration
            r[j] =  s[i]
        }
        //else discard the element
    }
}
```


Reference : Wikipedia
