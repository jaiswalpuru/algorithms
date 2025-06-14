class Foo {
    bool is_first;
    bool is_second;
    mutex mtx;
    condition_variable cv;
    
public:
    Foo() {
        is_first = false;
        is_second = false;
    }

    void first(function<void()> printFirst) {
        unique_lock<mutex> lock(mtx);
        // printFirst() outputs "first". Do not change or remove this line.
        printFirst();
        is_first = true;
        cv.notify_all();
    }

    void second(function<void()> printSecond) {
        unique_lock<mutex> lock(mtx);
        cv.wait(lock, [this]{ return is_first; });
        // printSecond() outputs "second". Do not change or remove this line.
        printSecond();
        is_second = true;
        cv.notify_all();
    }

    void third(function<void()> printThird) {
        unique_lock<mutex> lock(mtx);
        cv.wait(lock, [this]{ return is_first && is_second; });
        // printThird() outputs "third". Do not change or remove this line.
        printThird();
        cv.notify_all();
    }
};
