class ZeroEvenOdd {
    private int n;
    private Semaphore zeroSem, oddSem, eveSem;

    public ZeroEvenOdd(int n) {
        this.n = n;
        zeroSem = new Semaphore(1);
        eveSem = new Semaphore(0);
        oddSem = new Semaphore(0);
    }

    // printNumber.accept(x) outputs "x", where x is an integer.
    public void zero(IntConsumer printNumber) throws InterruptedException {
        for (int i=1; i<=n; i++) {
            zeroSem.acquire();
            printNumber.accept(0);
            if (i%2 == 0) eveSem.release();
            else oddSem.release();
        }
    }

    public void even(IntConsumer printNumber) throws InterruptedException {
        for (int i=2; i<=n; i+=2) {
            eveSem.acquire();
            printNumber.accept(i);
            zeroSem.release();
        }
    }

    public void odd(IntConsumer printNumber) throws InterruptedException {
        for(int i=1; i<=n; i+=2) {
            oddSem.acquire();
            printNumber.accept(i);
            zeroSem.release();
        }
    }
}
