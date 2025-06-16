class FizzBuzz {
private:
    int n;
    int curr;
    condition_variable cv;
    mutex mtx;
public:
    FizzBuzz(int n) {
        this->n = n;
        this->curr = 1;
    }

    // printFizz() outputs "fizz".
    void fizz(function<void()> printFizz) {
        while (true) {
            unique_lock<mutex> lock(mtx);
            cv.wait(lock, [this]{ return curr > n || curr % 3 == 0 && curr % 5 != 0; });
            if (curr > n) break;
            printFizz();
            ++curr;
            cv.notify_all();
        }
    }

    // printBuzz() outputs "buzz".
    void buzz(function<void()> printBuzz) {
        while (true) {
            unique_lock<mutex> lock(mtx);
            cv.wait(lock, [this]{ return curr > n || (curr % 3 != 0 && curr % 5 == 0); });
            if (curr > n) break;
            printBuzz();
            ++curr;
            cv.notify_all();
        }
    }

    // printFizzBuzz() outputs "fizzbuzz".
	void fizzbuzz(function<void()> printFizzBuzz) {
        while (true) {
            unique_lock<mutex> lock(mtx);
            cv.wait(lock, [this]{ return curr > n || (curr % 3 == 0 && curr % 5 == 0); });
            if (curr > n) break;
            printFizzBuzz();
            ++curr;
            cv.notify_all();
        }
    }

    // printNumber(x) outputs "x", where x is an integer.
    void number(function<void(int)> printNumber) {
        while (true) {
            unique_lock<mutex> lock(mtx);
            cv.wait(lock, [this]{ return curr > n || (curr % 3 != 0 && curr % 5 != 0); });
            if (curr > n) break;
            printNumber(curr);
            ++curr;
            cv.notify_all();
        }
    }
};
