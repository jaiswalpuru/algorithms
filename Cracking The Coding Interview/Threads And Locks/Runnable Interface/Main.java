public class Main {
    public static void main(String[] args) {
        RunnableThreadExample rt = new RunnableThreadExample();
        Thread t = new Thread(rt);
        t.start();
        while(rt.count != 5) {
            try {
                Thread.sleep(250);
            }catch(InterruptedException ex) {
                ex.printStackTrace();
            }
        }
    }
}
