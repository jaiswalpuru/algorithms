class Leaderboard {
    Map<Integer, Integer> playerExists;
    TreeMap<Integer, HashSet<Integer>> playerScoreBoard;
    public Leaderboard() {
        playerExists = new HashMap<>();
        playerScoreBoard = new TreeMap<>();
    }
    
    public void addScore(int playerId, int score) {
        int oldScore = 0;
        if (playerExists.containsKey(playerId)) {
            oldScore = playerExists.get(playerId);
            playerScoreBoard.get(oldScore).remove(playerId);
        }
        playerExists.put(playerId, score+oldScore);
        playerScoreBoard.computeIfAbsent(score+oldScore, k -> new HashSet<>()).add(playerId);
    }
    
    public int top(int K) {
        int sum = 0;
        for (Integer score : playerScoreBoard.descendingKeySet()) {
            if (K == 0) break;
            int size = playerScoreBoard.get(score).size();
            if (size > K) {
                sum += K * score;
                K = 0;
            } else {
                K -= size;
                sum += size*score;
            }
        }
        return sum;
    }
    
    public void reset(int playerId) {
        int score = playerExists.get(playerId);
        playerExists.remove(playerId);
        playerScoreBoard.get(score).remove(playerId);
        if (playerScoreBoard.get(score).size() == 0) playerScoreBoard.remove(score);
    }
}

/**
 * Your Leaderboard object will be instantiated and called as such:
 * Leaderboard obj = new Leaderboard();
 * obj.addScore(playerId,score);
 * int param_2 = obj.top(K);
 * obj.reset(playerId);
 */
