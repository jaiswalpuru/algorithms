class H2O {
    condition_variable cv;
    int h2;
    mutex mtx;
public:
    H2O() {
        h2 = 0;
    }

    void hydrogen(function<void()> releaseHydrogen) {
        unique_lock<mutex> lock(mtx);
        cv.wait(lock, [this]{ return h2 < 2; });
        h2++;
        // releaseHydrogen() outputs "H". Do not change or remove this line.
        releaseHydrogen();
        cv.notify_all();
    }

    void oxygen(function<void()> releaseOxygen) {
        unique_lock<mutex> lock(mtx);
        cv.wait(lock, [this]{ return h2 == 2; });
        // releaseOxygen() outputs "O". Do not change or remove this line.
        h2 = 0;
        releaseOxygen();
        cv.notify_all();
    }
};
