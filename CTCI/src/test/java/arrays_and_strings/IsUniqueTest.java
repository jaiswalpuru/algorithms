package arrays_and_strings;

import com.pj.ctci.arrays_and_strings.IsUnique;
import org.junit.jupiter.api.Test;

public class IsUniqueTest {
	@org.junit.Test
	public void testIsUniqueOne() {
		assert IsUnique.isUnique("abc");
	}

	@Test
	public void testNonUnique() {
		assert !IsUnique.isUnique("aabbcc");
	}
}
