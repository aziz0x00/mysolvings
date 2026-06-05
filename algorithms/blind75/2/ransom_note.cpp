// *lesson learnt*: try simple array before hashing
class Solution {
public:
    bool canConstruct(string ransomNote, string magazine) {
        int i = 0;
        int counter[26] = {0};
        for(auto& c : magazine) ++counter[c - 'a'];
        for(auto& c : ransomNote) {
            if (!counter[c - 'a']) return false;
            --counter[c - 'a'];
        }
        return true;
    }
};
