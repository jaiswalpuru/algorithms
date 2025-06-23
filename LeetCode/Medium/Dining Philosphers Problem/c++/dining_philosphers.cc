class DiningPhilosophers {
    int n;
    mutex mtx;
    vector<bool> forks;
    condition_variable cv;
public:
    DiningPhilosophers() {
        n = 5;
        forks = vector<bool>(n, false);
    }

    void wantsToEat(int p,
                    function<void()> pickLeftFork,
                    function<void()> pickRightFork,
                    function<void()> eat,
                    function<void()> putLeftFork,
                    function<void()> putRightFork) {
		unique_lock<mutex> l(mtx);
        cv.wait(l, [&]{ return !forks[p] && !forks[p + 1]; });
        forks[p] = forks[(p + 1) % n] = true;
        pickLeftFork();
        pickRightFork();
        eat();
        putLeftFork();
        putRightFork();
        forks[p] = forks[(p + 1) % n] = false;
        cv.notify_all();
    }
};
