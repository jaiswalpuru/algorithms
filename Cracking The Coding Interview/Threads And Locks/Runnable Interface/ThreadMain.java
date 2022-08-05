public class ThreadMain {
    public static void main(String[] args) {
        ThreadExample te = new ThreadExample();
        te.start();
        while(te.count != 5 ){
            try {
                Thread.sleep(250);
            }catch(InterruptedException ex) {
                ex.printStackTrace();    
            }
        }
    }
} 
