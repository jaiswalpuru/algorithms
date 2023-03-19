class BrowserHistory {
    ArrayList<String> urls;
    int currPointer;
    int lastPointer;

    public BrowserHistory(String homepage) {
        urls = new ArrayList<>();
        urls.add(homepage);
        currPointer = 0;
        lastPointer = 0;    
    }
    
    public void visit(String url) {
        currPointer += 1;
        if (urls.size() > currPointer) urls.set(currPointer, url);
        else urls.add(url);
        lastPointer = currPointer;
    }
    
    public String back(int steps) {
        currPointer = Math.max(currPointer-steps, 0);
        return urls.get(currPointer);
    }
    
    public String forward(int steps) {
        currPointer = Math.min(lastPointer, currPointer+steps);
        return urls.get(currPointer);
    }
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * BrowserHistory obj = new BrowserHistory(homepage);
 * obj.visit(url);
 * String param_2 = obj.back(steps);
 * String param_3 = obj.forward(steps);
 */
