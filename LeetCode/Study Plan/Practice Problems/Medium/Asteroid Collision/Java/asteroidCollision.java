class Solution {
    public int[] asteroidCollision(int[] asteroids) {
        int size = asteroids.length;
        Stack<Integer> s = new Stack<>();

        for (Integer aster : asteroids) {
            if (s.isEmpty()) {
                s.add(aster);
                continue;
            }
            
            if (aster<0 && s.peek().equals(Math.abs(aster))) {
                s.pop();
            }else {
                boolean push = true;
                while (!s.isEmpty() && (s.peek() > 0 && aster < 0)) {
                    if (Math.abs(aster) > s.peek()) {
                        s.pop();
                    }else if (Math.abs(aster) == s.peek()) {
                        push = false;
                        s.pop();
                        break;
                    }else {
                        push = false;
                        break;
                    }
                }
                if (push) s.add(aster);
            }
        }
        int stackSize = s.size();
        int[] res = new int[s.size()];

        for (int i=stackSize-1; i>=0; i--) {
            res[i] = s.pop();
        }

        return res;
    }
}
