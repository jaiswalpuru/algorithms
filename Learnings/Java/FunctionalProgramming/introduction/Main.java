package introduction;

//A language of functions.

// In Mathematical sense :
/**
 *  Same input -->  Same output
 *  This is stateless.
 *  purity, absence of side effects.
 *  There are no variables, no assignments, no loops.
 */

public class Main {
	public static void main(String[] args) {

	}

	public static int arraySumIterative(int[] array) {
		int sum = 0;
		for (int val : array) sum += val;
		return sum;
	}

	public static int arraySumRecursive(int[] array, int start) {
		if (start >= array.length) return 0;
		return array[start] + arraySumRecursive(array, start+1);
	}
}
