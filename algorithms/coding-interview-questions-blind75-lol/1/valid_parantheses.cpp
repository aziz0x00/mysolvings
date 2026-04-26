class Solution {
public:
    bool isValid(string s) {
        // map<char, int> ctr     = {{'{', 0}, {'(', 0}, {'[', 0}};
        map<char, int> closing = {{'}', '{'}, {']', '['}, {')', '('}};
        char last_opened;
        vector<char> stack;
        for (char c : s) {
            if (closing.count(c)) {
                c = closing[c];
                if (!size(stack)) return false;
                last_opened = stack.back();
                stack.pop_back();
                if (c != last_opened)
                    return false;
            } else {
                stack.push_back(c);
            }
        }
        return size(stack) == 0;
    }
};
