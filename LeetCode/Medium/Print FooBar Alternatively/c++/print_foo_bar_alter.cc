class FooBar {
private:
    int n;
    bool is_foo_printed;
    condition_variable cv;
    mutex mtx;

public:
    FooBar(int n) {
        this->n = n;
        is_foo_printed = false;
    }

    void foo(function<void()> printFoo) {
        
        for (int i = 0; i < n; i++) {
            unique_lock<mutex> lock(mtx);
            cv.wait(lock, [this]{ return !is_foo_printed; });
        	printFoo();
            is_foo_printed = true;
            cv.notify_all();
        }
    }

    void bar(function<void()> printBar) {
        
        for (int i = 0; i < n; i++) {
            unique_lock<mutex> lock(mtx);
            cv.wait(lock, [this]{ return is_foo_printed; });
        	printBar();
            is_foo_printed = false;
            cv.notify_all();
        }
    }
};
