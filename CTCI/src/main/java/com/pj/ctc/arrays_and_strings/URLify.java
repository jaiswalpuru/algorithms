package com.pj.ctc.arrays_and_strings;

import java.util.Arrays;

/**
 * Replace all the space with %20, you may assume that the string has enough space to store at the end.
 */

public class URLify {
	public static char[] urLify(String str, int n) {
		char[] chars = new char[str.length()];
		int k = 0;
		for (int i=0; i<n; i++) {
			if (str.charAt(i) == ' ') {
				chars[k] = '%';
				chars[k+1] = '2';
				chars[k+2] = '0';
				k += 3;
			} else {
				chars[k++] = str.charAt(i);
			}
		}
		return Arrays.copyOf(chars, k);
	}

}
