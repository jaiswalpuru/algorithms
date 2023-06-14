package com.pj.atcoder.beginner;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class PlacingMarbles {
	public static void main(String[] args) throws IOException {
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		String s = br.readLine();
		int marbles = 0;
		for (char c : s.toCharArray()) if (c=='1') marbles++;
		System.out.println(marbles);
	}
}
