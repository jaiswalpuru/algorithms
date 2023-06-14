package com.pj.atcoder.beginner;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.HashMap;

public class Traveling {
	public static void main(String[] args) throws IOException {
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		int n = Integer.parseInt(br.readLine());
		int x = 0;
		int y = 0;
		int initT = 0;
		boolean isPossible = true;
		for (int k=0; k<n; k++) {
			String[] str = br.readLine().split(" ");
			int t = Integer.parseInt(str[0]);
			int i = Integer.parseInt(str[1]);
			int j = Integer.parseInt(str[2]);
			if (!canReach(t-initT, Math.abs(x-i), Math.abs(y-j))) {
				isPossible = false;
			}
			x = i;
			y = j;
			initT = t;
		}
		System.out.println(isPossible ? "Yes" : "No");
	}

	private static boolean canReach(int t, int x, int y) {
		return (x+y == t) || (t > x + y && (t - (x + y)) % 2 == 0);
	}
}