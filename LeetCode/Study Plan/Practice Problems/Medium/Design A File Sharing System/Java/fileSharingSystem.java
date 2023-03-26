class FileSharing {
    int uid;
    int chunks;
    PriorityQueue<Integer> userIDLeft;
    HashMap<Integer, HashSet<Integer>> userChunkID;
    HashMap<Integer, HashSet<Integer>> chunkUserID;

    public FileSharing(int m) {
        chunks = m;
        uid = -1;
        userIDLeft = new PriorityQueue<>();
        userChunkID = new HashMap<>();
        chunkUserID = new HashMap<>();
    }
    
    public int join(List<Integer> ownedChunks) {
        int id = -1;
        if (uid == -1) {
            id = 1;
            uid = 1;
        }else {
            if (userIDLeft.size() == 0) {
                id = uid+1;
                uid++;
            }else {
                id = userIDLeft.remove();
            }
        }
        userChunkID.put(id, new HashSet<>());
        userChunkID.get(id).addAll(ownedChunks);
        return id;
    }
    
    public void leave(int userID) {
        userIDLeft.add(userID);
        userChunkID.remove(userID);
        for (Integer id : chunkUserID.keySet()) {
            if (chunkUserID.get(id).contains(userID)) {
                chunkUserID.get(id).remove(userID);
            }
        }
    }
    
    public List<Integer> request(int userID, int chunkID) {
        List<Integer> users = new ArrayList<>();
        for (Integer id : userChunkID.keySet()) {
            if (userChunkID.get(id).contains(chunkID)) 
                users.add(id);
        }

        if (users.size() > 0) {
            userChunkID.get(userID).add(chunkID);
            chunkUserID.getOrDefault(chunkID, new HashSet<>()).add(userID);
        }
        Collections.sort(users);
        return users;
    }
}

/**
 * Your FileSharing object will be instantiated and called as such:
 * FileSharing obj = new FileSharing(m);
 * int param_1 = obj.join(ownedChunks);
 * obj.leave(userID);
 * List<Integer> param_3 = obj.request(userID,chunkID);
 */
