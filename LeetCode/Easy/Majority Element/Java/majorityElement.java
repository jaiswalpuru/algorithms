// method 1 using hashmap
class Solution {
    public int majorityElement(int[] nums) {
        HashMap<Integer, Integer> mp = new HashMap<>();
        for (int num : nums) mp.put(num, mp.getOrDefault(num, 0)+1);
        int major = 0;
        int cnt = 0;
        for (Integer k : mp.keySet()) {
            int val = mp.get(k);
            if (cnt < val) {
                cnt = val;
                major = k;
            }
        }
        return major;
    }
}

//method 2 using sorting
class Solution {
    public int majorityElement(int[] nums) {
        Arrays.sort(nums);
        return nums[nums.length/2];
    }
}

// using bits 
//refer LC solutions
class Solution {
    public int majorityElement(int[] nums) {
        int n = nums.length;
        int major = 0;
        for (int i=0; i<32; i++) {
            int bit = 1<<i;
            //to count how many numbers have this bit set
            int bitCount = 0;
            for (int num : nums)
                if ((num & bit) != 0) bitCount++;

            if (bitCount > n/2) major |= bit; // prepare the number
        }
        return major;
    }
}
