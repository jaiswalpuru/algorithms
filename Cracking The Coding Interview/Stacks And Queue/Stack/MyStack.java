//not executbale code it just a template
public class MyStack {
    private static class StackNode<T> {
        private T data;
        private StackNode<T> next;

        public StackNode(T data) { this.data = data; }
    }

    private StackNode<T> top;
    
    public T Pop() {
        if (top == null) throw new EmptyStackException();
        T item = top.data;
        top = top.next;
        return item;
    }

    public void Push(T item) {
        StackNode<T> t = new StackNode<T>(item);
        t.next = top;
        top = t;
    }

    public T peek() {
        if (top == null) throw new EmptyStackException();
        return top.data;
    }

    public boolean isEmpty() { return top == null; }
}

