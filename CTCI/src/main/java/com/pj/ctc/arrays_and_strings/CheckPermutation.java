package com.pj.ctc.arrays_and_strings;

import java.util.Arrays;

/**
 * Check permutation : Given two strings, write a method to decide if one is a permutation of other.
 */

public class CheckPermutation {

	public static boolean approachOne(String s1, String s2) {
		int[] charCnt = new int[128];
		for (char c : s1.toCharArray()) charCnt[c]++;
		for (char c : s2.toCharArray()) {
			charCnt[c]--;
			if (charCnt[c] < 0) return false;
		}
		return true;
	}

	public static boolean approachTwo(String s1, String s2) {
		char[] charS1 = s1.toCharArray();
		char[] charS2 = s2.toCharArray();
		Arrays.sort(charS1);
		Arrays.sort(charS2);
		return Arrays.toString(charS1).equals(Arrays.toString(charS2));
	}
}
