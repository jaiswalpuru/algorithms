package com.pj.spoj;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class TEST {
	public static void main(String[] args) throws IOException {
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		int num;
		while(true) {
			num = Integer.parseInt(br.readLine());
			if (num == 42) break;
			System.out.printf("%d\n", num);
		}
	}
}
