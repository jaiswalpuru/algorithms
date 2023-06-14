package com.pj.spoj;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class PRIME1 {
	public static void main(String[] args) throws IOException {
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		int t = Integer.parseInt(br.readLine());
		while(t-- > 0) {
			String[] input = br.readLine().split(" ");
			int l = Integer.parseInt(input[0]);
			int h = Integer.parseInt(input[1]);
			segmentedSieve(l, h);
		}
	}

	public static void segmentedSieve(int l, int h) {
		List<Integer> prime = sieve((int)Math.sqrt(h));
		boolean[] dummy = new boolean[h-l+1];
		Arrays.fill(dummy, true);
		//dummy[0] indicates if l is prime or not, similarly 1 indicates if l+1 is prime or not.
		for (Integer val : prime) {
			int low = l/val;
			if (low <= 1) low = val+val;
			else if (l%val != 0) {
				low = (low*val)+val;
			}else {
				low = low*val;
			}
			for (int j = low; j<=h; j+=val) {
				dummy[j-l] = false;
			}
		}
		for (int i=l; i<=h; i++) {
			if (i == 1) continue;
			if (dummy[i-l]) System.out.println(i);
		}
		System.out.println();
	}

	// O(N log(logN))
	public static List<Integer> sieve(int n) {
		boolean[] s = new boolean[n+1];
		Arrays.fill(s, true);
		s[0] = false;
		s[1] = false;
		for (int i=2; (i*i)<=n; i++) {
			if (s[i]) {
				for (int j=i*i; j<=n; j+=i) {
					s[j] = false;
				}
			}
		}
		List<Integer> primeNumbers = new ArrayList<>();
		for (int i=2; i<=n; i++) {
			if (s[i]) primeNumbers.add(i);
		}
		return primeNumbers;
	}

	public static List<Integer> bruteForce(int n) {
		List<Integer> primeNumbers = new ArrayList<>();
		for (int i=2; i<=n; i++) {
			if (isPrime(i)) primeNumbers.add(i);
		}
		return primeNumbers;
	}

	public static boolean isPrime(int n) {
		if (n == 2 || n == 3 || n == 5 || n == 7) return true;
		if (n <= 10) return false;
		for (int i=2; i<=(int)Math.sqrt(n)+1; i++) if (n%i == 0) return false;
		return true;
	}

}


/*		Refer : Scaler : https://www.scaler.com/topics/segmented-sieve/
 * 		Approach for segmented sieve
 * 		1. Generate all primes in the range sqrt(R) and store it in array.
 * 		2. Initialize a dummy boolean array of size (R-L+1) and initialize all values to true.
 * 		3. The idea is to partition the interval [L, R] into smaller segments and calculate the prime, only upto sqrt(R)
 * 		4. The element at index 0 represents L
 */