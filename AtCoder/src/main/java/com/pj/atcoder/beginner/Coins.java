package com.pj.atcoder.beginner;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.Arrays;

public class Coins {
	static int[] changes = new int[3];
	public static void main(String[] args) throws IOException {
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		int[] coins = new int[3];
		coins[0] = Integer.parseInt(br.readLine()); // 500
		coins[1] = Integer.parseInt(br.readLine()); // 100
		coins[2] = Integer.parseInt(br.readLine()); // 50
		changes[0] = 500;
		changes[1] = 100;
		changes[2] = 50;
		int x = Integer.parseInt(br.readLine());
		System.out.println(recur(x, coins, 0, 0));
	}

	private static int recur(int total, int[] coins, int sum, int ind) {
		if (ind >= 3) {
			if (sum == total) return 1;
			return 0;
		}
		int take = 0;
		if (coins[ind] > 0) {
			coins[ind] -= 1;
			take = recur(total, coins, sum + changes[ind], ind);
			coins[ind] += 1;
		}
		int dntTake = recur(total, coins, sum, ind+1);
		return take+dntTake;
	}
}
