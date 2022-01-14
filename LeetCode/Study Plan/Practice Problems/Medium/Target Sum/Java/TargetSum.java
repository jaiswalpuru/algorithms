import java.util.Arrays;

public class TargetSum {
    public static int cnt = 0;
    public static void main(String[] args) {
        int[] nums ={
            1, 1, 1, 1, 1
        };
        System.out.println(findTargetSumWays(nums, 3));
        System.out.println(findTargetSumWaysMemo(nums, 3));
        System.out.println(findTargetSumWaysDP2(nums, 3));
    }

    private static int findTargetSumWaysDP2(int []nums, int target) {
        int total = Arrays.stream(nums).sum();
        int[][] dp = new int[nums.length][2*total+1];

        dp[0][nums[0]+total] = 1;
        dp[0][-nums[0]+total] +=1;

        for (int i=1;i<nums.length;i++){
            for (int sum=-total;sum<=total;sum++){
                if (dp[i-1][sum+total]>0){
                    dp[i][nums[i]+sum+total] +=dp[i-1][sum+total];
                    dp[i][-nums[i]+sum+total] +=dp[i-1][sum+total];
                }
            }
        }
        return Math.abs(target)>total?0:dp[nums.length-1][target+total];
    }

    public static int findTargetSumWays(int[] nums, int target) {
        int sum = 0;
        finalTargetSumWays(nums, target, sum,0);
        return cnt;
    }

    private static void finalTargetSumWays(int[] nums, int target,int sum, int start) {
        if (start == nums.length){
            if (sum == target) {
                cnt++;
            }
            return ;
        }
        finalTargetSumWays(nums, target, sum+nums[start], start+1);
        finalTargetSumWays(nums, target, sum-nums[start], start+1);
    }

    private static int findTargetSumWaysMemo(int []nums, int target){
        int total = Arrays.stream(nums).sum();
        int[][] memo = new int[nums.length][2*total+1];

        for (int[] is : memo) {
            Arrays.fill(is, Integer.MIN_VALUE);
        }
        return calculateSumWays(nums,0,0,target,total,memo);
    }

    private static int calculateSumWays(int[] nums, int start, int sum, int target,int total, int[][] memo){
        if (start == nums.length){
            if (sum == target){
                return 1;
            }
            return 0;
        }

        if (memo[start][sum+total]!=Integer.MIN_VALUE){
            return memo[start][sum+total];
        }
        int add = calculateSumWays(nums, start+1, sum+nums[start], target, total, memo);
        int subtract = calculateSumWays(nums, start+1, sum-nums[start], target, total, memo);
        memo[start][sum+total] = add+subtract;
        return memo[start][sum+total];
    }
}