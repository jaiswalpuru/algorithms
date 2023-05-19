package com.pj.ctc.arrays_and_strings;

/**
 * Given a string, tell whether all the characters present in it are unique or not, what if you cannot use any
 * additional data structure.
 */

public class IsUnique {
	public static boolean isUnique(String str) {
		int[] chars = new int[128]; // here assuming 128 characters present in string.
		for (char c : str.toCharArray()) {
			chars[c]++;
			if (chars[c] > 1) return false;
		}
		return true;
	}
}
