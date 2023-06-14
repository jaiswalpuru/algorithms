package com.pj.atcoder.beginner;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class SomeSums {
	public static void main(String[] args) throws IOException {
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		String[] in = br.readLine().split(" ");
		int n = Integer.parseInt(in[0]);
		int a = Integer.parseInt(in[1]);
		int b = Integer.parseInt(in[2]);
		int sum = 0;
		for (int i=1; i<=n; i++) {
			int s = sumOfDigits(i);
			if (s >= a && s <= b) sum += i;
		}
		System.out.println(sum);
	}

	private static int sumOfDigits(int n) {
		int sum = 0;
		while (n > 0) {
			sum += n%10;
			n /= 10;
		}
		return sum;
	}
}
