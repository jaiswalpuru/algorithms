class Pair {
    String stationName;
    int t;
    public Pair(String name, int t) {
        this.t = t;
        this.stationName = name;
    }
}

class UndergroundSystem {
    Map<Integer, Pair> checkInTrack;
    Map<String, HashMap<String, Integer>> travel;
    Map<String, HashMap<String, Integer>> times;
    public UndergroundSystem() {
        checkInTrack = new HashMap<>();
        travel = new HashMap<>();
        times = new HashMap<>();
    }
    
    public void checkIn(int id, String stationName, int t) {
        checkInTrack.put(id, new Pair(stationName, t));
    }
    
    public void checkOut(int id, String stationName, int t) {
        Pair details = checkInTrack.get(id);
        int diff = t-details.t;
        travel.computeIfAbsent(details.stationName, k -> new HashMap<>());
        times.computeIfAbsent(details.stationName, k -> new HashMap<>());
        HashMap<String, Integer> toStation = travel.get(details.stationName);
        HashMap<String, Integer> toStationtimes = times.get(details.stationName);
        toStation.put(stationName, toStation.getOrDefault(stationName, 0) + diff);
        toStationtimes.put(stationName, toStationtimes.getOrDefault(stationName, 0)+1);
    }
    
    public double getAverageTime(String startStation, String endStation) {
        return (double)travel.get(startStation).get(endStation)/times.get(startStation).get(endStation);
    }
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * UndergroundSystem obj = new UndergroundSystem();
 * obj.checkIn(id,stationName,t);
 * obj.checkOut(id,stationName,t);
 * double param_3 = obj.getAverageTime(startStation,endStation);
 */
