package com.pj.atcoder.beginner;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class WelcomeToAtCoder {
	public static void main(String[] args) throws IOException {
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		int a = Integer.parseInt(br.readLine());
		String[] s = br.readLine().split(" ");
		String msg = br.readLine();
		int b = Integer.parseInt(s[0]);
		int c = Integer.parseInt(s[1]);
		System.out.println((a+b+c) + " " + msg);
	}
}
