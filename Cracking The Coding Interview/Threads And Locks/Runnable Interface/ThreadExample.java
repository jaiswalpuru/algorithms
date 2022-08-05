public class ThreadExample extends Thread{
    int count = 0;
    public void run() {
        System.out.println("Thread starting");
        try {
            while(count < 5) {
                Thread.sleep(500);
                count++;
                System.out.println("Inside thread count is " + count);
            }
        }catch(InterruptedException ex) {
            System.out.println("Thread interrupting");
        }
        System.out.println("Thread terminating");
    }
}
