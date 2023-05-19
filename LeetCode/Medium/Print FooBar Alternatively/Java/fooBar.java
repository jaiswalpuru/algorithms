class FooBar {
    private int n;
    private static AtomicBoolean atb = new AtomicBoolean(false);
    public FooBar(int n) {
        this.n = n;
    }

    public void foo(Runnable printFoo) throws InterruptedException {
        for (int i = 0; i < n; i++) {
            while(atb.get()){}
        	// printFoo.run() outputs "foo". Do not change or remove this line.
        	printFoo.run();
            atb.set(true);
        }
    }

    public void bar(Runnable printBar) throws InterruptedException {
        for (int i = 0; i < n; i++) {
            while(!atb.get()){}
            // printBar.run() outputs "bar". Do not change or remove this line.
        	printBar.run();
            atb.set(false);
        }
    }
}
