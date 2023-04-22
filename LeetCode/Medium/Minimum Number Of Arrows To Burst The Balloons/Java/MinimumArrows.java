import java.util.Arrays;
import java.util.Comparator;

public class MinimumArrows{
    public static void main(String[] args) {
        int[][] balloons = {
            {10,16},
            {2,8},
            {1,6},
            {7,12}
        };
        System.out.println(MinimumArrowsShot(balloons, balloons.length));
    }

    public static int MinimumArrowsShot(int[][] balloons, int n){
        Arrays.sort(balloons,new Comparator<int[]>() {
            public int compare(int[] a, int[] b){
                return Integer.compare(a[0], b[0]);
            }
        });

        int maxEnd = balloons[0][1];
        int minArrows = 1;
        for (int i=1; i<n; i++){
            if (balloons[i][0] <= maxEnd){
                maxEnd = Math.min(maxEnd, balloons[i][1]);
            }else{
                minArrows++;
                maxEnd = balloons[i][1];
            }
        }


        return minArrows;
    }
}