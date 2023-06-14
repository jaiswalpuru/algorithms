package com.pj.atcoder.beginner;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.Collections;
import java.util.PriorityQueue;

public class CardGameForTwo {
	public static void main(String[] args) throws IOException {
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		int n = Integer.parseInt(br.readLine());
		String[] in = br.readLine().split(" ");
		PriorityQueue<Integer> pq = new PriorityQueue<>(Collections.reverseOrder());
		for (String s : in) pq.add(Integer.parseInt(s));
		int alice = 0;
		int bob = 0;
		boolean turn = true;
		while (!pq.isEmpty()) {
			if (turn) alice += pq.poll();
			else bob += pq.poll();
			turn = !turn;
		}
		System.out.println(alice-bob);
	}
}
