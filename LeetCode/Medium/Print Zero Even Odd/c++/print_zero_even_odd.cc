class ZeroEvenOdd {
private:
    int n;
    condition_variable cv;
    mutex mtx;
    bool is_zero;
    bool is_even;
    bool is_odd;
public:
    ZeroEvenOdd(int n) {
        this->n = n;
        is_zero = true;
        is_even = false;
        is_odd = false;
    }

    // printNumber(x) outputs "x", where x is an integer.
    void zero(function<void(int)> printNumber) {
        for (int i = 1; i <= n; i++) {
            unique_lock<mutex> lock(mtx);
            cv.wait(lock, [this]{ return is_zero; });
            printNumber(0);
            if (i % 2) {
                is_odd = true;
            } else {
                is_even = true;
            }
            is_zero = false;
            cv.notify_all();
        }
    }

    void even(function<void(int)> printNumber) {
        for (int i = 2; i <= n; i += 2) {
            unique_lock<mutex> lock(mtx);
            cv.wait(lock, [this]{ return is_even; });
            is_zero = true;
            is_even = false;
            printNumber(i);
            cv.notify_all();
        }
    }

    void odd(function<void(int)> printNumber) {
        for (int i = 1; i <= n; i += 2) {
            unique_lock<mutex> lock(mtx);
            cv.wait(lock, [this]{ return is_odd; });
            is_zero = true;
            is_odd = false;
            printNumber(i);
            cv.notify_all();
        }
    }
};


class ZeroEvenOdd {
private:
    int n;
    sem_t sem_zero;
    sem_t sem_even;
    sem_t sem_odd;
public:
    ZeroEvenOdd(int n) {
        this->n = n;
        sem_init(&sem_zero, 0, 1);
        sem_init(&sem_even, 0, 0);
        sem_init(&sem_odd, 0, 0);
    }

    // printNumber(x) outputs "x", where x is an integer.
    void zero(function<void(int)> printNumber) {
        for (int i = 1; i <= n; i++) {
            sem_wait(&sem_zero);
            printNumber(0);
            if (i % 2) {
                sem_post(&sem_odd);
            } else {
                sem_post(&sem_even);
            }
        }
    }

    void even(function<void(int)> printNumber) {
        for (int i = 2; i <= n; i += 2) {
            sem_wait(&sem_even);
            printNumber(i);
            sem_post(&sem_zero);
        }
    }

    void odd(function<void(int)> printNumber) {
        for (int i = 1; i <= n; i += 2) {
            sem_wait(&sem_odd);
            printNumber(i);
            sem_post(&sem_zero);
        }
    }
};
