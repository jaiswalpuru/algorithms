public class RunnableThreadExample implements Runnable {
    public int count = 0;

    public void run() {
        System.out.println("Runnable thread started");
        try {
            while(count < 5 ) {
                Thread.sleep(500);
                count++;
            }
        }catch (InterruptedException ex) {
            System.out.println("Runnable Thread intertupted");
        }
        System.out.println("Runnable Thread terminating");
    }
}
