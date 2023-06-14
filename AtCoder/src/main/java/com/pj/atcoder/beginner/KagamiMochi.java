package com.pj.atcoder.beginner;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.Arrays;
import java.util.Comparator;

public class KagamiMochi {
	public static void main(String[] args) throws IOException {
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		int n = Integer.parseInt(br.readLine());
		Integer[] in = new Integer[n];
		for (int i=0; i<n; i++)
			in[i] = Integer.parseInt(br.readLine());

		Arrays.sort(in, Comparator.reverseOrder());
		int cnt = 1;
		int currMax = 1;
		for (int i=1; i<n; i++) {
			if (in[i] < in[i-1]) {
				currMax++;
			} else if (in[i] > in[i-1]){
				cnt = Math.max(cnt, currMax);
				currMax = 1;
			}
		}
		cnt = Math.max(cnt, currMax);
		System.out.println(cnt);
	}
}
