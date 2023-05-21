package arrays_and_strings;

import com.pj.ctc.arrays_and_strings.URLify;
import org.junit.jupiter.api.Test;

import java.util.Arrays;

public class URLifyTest {
	@Test
	public void urlifyTestOne() {
		assert Arrays.equals(URLify.urLify("Mr John Smith      ", 13), "Mr%20John%20Smith".toCharArray());
	}
}
