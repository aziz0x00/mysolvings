class Solution {
public:
    vector<vector<int>> updateMatrix(vector<vector<int>>& mat) {
        int n = mat.size(), m = mat[0].size();

        vector<vector<int>> ans(n, vector<int>(m, -1));
        vector<vector<int>> visited(n, vector<int>(m, 0));

        queue<pair<int, int>> q;
        for(int i = 0; i < n; ++i)
            for(int j = 0; j < m; ++j)
                if (mat[i][j] == 0) {
                    q.push({i, j});
                    ans[i][j] = 0;
                }

        int DIR[] = {0, -1, 0, 1, 0};
        while(!q.empty()) {
            auto [i, j] = q.front(); q.pop();
            for(int k = 0; k < 4; ++k) {
                int ni = i + DIR[k];
                int nj = j + DIR[k+1];
                if(ni >= 0 && ni < n && nj >= 0 && nj < m) {

                    if (ans[ni][nj] != -1) continue;

                    ans[ni][nj] = ans[i][j] + 1;
                    q.push({ni, nj});
                }
            }
        }
        return ans;
    }
};
