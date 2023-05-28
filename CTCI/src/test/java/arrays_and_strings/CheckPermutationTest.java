package arrays_and_strings;

import com.pj.ctci.arrays_and_strings.CheckPermutation;
import org.junit.jupiter.api.Test;

public class CheckPermutationTest {
	@Test
	public void checkPermutationOne() {
		assert CheckPermutation.approachOne("bvcxz", "zxcvb");
	}

	@Test
	public void checkPermutationTwo() {
		assert !CheckPermutation.approachOne("abc", "qwe");
	}
}
