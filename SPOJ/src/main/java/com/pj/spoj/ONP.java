package com.pj.spoj;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.time.Duration;
import java.time.Instant;
import java.util.*;

public class ONP {
	public static void main(String[] args) throws IOException {
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		int t = Integer.parseInt(br.readLine());
		String text;
		Instant start = Instant.now();
		while (t-- > 0) {
			text = br.readLine();
			ArrayDeque<Character> q = new ArrayDeque<>();
			ArrayDeque<Character> stackOp = new ArrayDeque<>();
			for (Character c : text.toCharArray()) {
				if (c == '+' || c == '-' || c == '*' || c == '/' || c == '^') {
					stackOp.push(c);
				} else if (c == ')') {
					if (!stackOp.isEmpty()) {
						q.add(stackOp.pop());
					}
				} else if (c >= 'a' && c <= 'z') {
					q.add(c);
				}
			}
			for (Character c : q) System.out.print(c);
			System.out.println();
		}
		Instant end = Instant.now();
		System.out.println("Time took : " + Duration.between(start, end).toMillis());
	}
}
