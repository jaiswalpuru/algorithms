class LRUCache {

    public class Node {
        int key;
        int val;
        Node(int k, int v) {
            key = k;
            val = v;
        }
    }

    private final int capacity;
    private Map<Integer, Node> map;
    Deque<Node> dll;

    public LRUCache(int capacity) {
        map = new HashMap<>();
        dll = new LinkedList<>();
        this.capacity = capacity;
    }
    
    public int get(int key) {
        if (map.containsKey(key)) {
            Node ent = map.get(key);
            dll.remove(ent);
            dll.addFirst(ent);
            return ent.val;
        }
        return -1;
    }
    
    public void put(int key, int value) {
        if (map.containsKey(key)) {
            Node ent = map.get(key);
            dll.remove(ent);
            ent.val = value;
            dll.addFirst(ent);
        } else {
            if (map.size() == capacity) {
                Node ent = dll.removeLast();
                map.remove(ent.key);
            }
            Node ent = new Node(key, value);
            map.put(key, ent);
            dll.addFirst(ent);
        }
    }
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache obj = new LRUCache(capacity);
 * int param_1 = obj.get(key);
 * obj.put(key,value);
 */

//a better approach
class LRUCache extends LinkedHashMap<Integer, Integer> {
    private final int capacity;

    public LRUCache(int capacity) {
        super(capacity, 0.75f, true);
        this.capacity = capacity;    
    }
    
    public int get(int key) {
        return super.getOrDefault(key, -1);
    }
    
    public void put(int key, int value) {
        super.put(key, value);
    }
    
    @Override
    protected boolean removeEldestEntry(Map.Entry<Integer, Integer> eldest) {
        return size() > capacity;
    }
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache obj = new LRUCache(capacity);
 * int param_1 = obj.get(key);
 * obj.put(key,value);
 */

