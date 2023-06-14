package com.pj.atcoder.beginner;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.Arrays;
import java.util.stream.IntStream;

public class ShiftOnly {
	public static void main(String[] args) throws IOException {
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		int n = Integer.parseInt(br.readLine());
		String[] s = br.readLine().split(" ");
		IntStream intStream = Arrays.stream(s).mapToInt(Integer::parseInt);
		int[] arr = intStream.toArray();
		int op = 0;
		HERE : while (true) {
			for (int i = 0; i < n; i++) {
				if (arr[i]%2 ==1) {
					break HERE;
				}
				arr[i] /= 2;
			}
			op++;
		}

		System.out.println(op);
	}
}
