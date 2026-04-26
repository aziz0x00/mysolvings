class MyQueue {
private:
    vector<int> stack1;
    vector<int> stack2;
public:
    MyQueue() {
    }

    void push(int x) {
        while (stack2.size()) {
            stack1.push_back(stack2.back());
            stack2.pop_back();
        }
        stack1.push_back(x);
    }

    int pop() {

        while(stack1.size()) {
            stack2.push_back(stack1.back());
            stack1.pop_back();
        }

        int ret = stack2.back();
        stack2.pop_back();
        return ret;
    }
    
    int peek() {

        while(stack1.size()) {
            stack2.push_back(stack1.back());
            stack1.pop_back();
        }

        return stack2.back();
    }
    
    bool empty() {
        return stack1.size() == 0 and stack2.size() == 0;
    }
};

/**
 * Your MyQueue object will be instantiated and called as such:
 * MyQueue* obj = new MyQueue();
 * obj->push(x);
 * int param_2 = obj->pop();
 * int param_3 = obj->peek();
 * bool param_4 = obj->empty();
 */
