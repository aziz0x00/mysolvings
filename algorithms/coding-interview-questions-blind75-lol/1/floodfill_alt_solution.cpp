// faster than the DFS version
class Solution {
public:
    vector<vector<int>> floodFill(vector<vector<int>>& image, int sr, int sc, int color) {

        if (image[sr][sc] == color)
            return image;

        int n = size(image);
        int m = size(image[0]);
        int init_color = image[sr][sc];
        // BFS
        queue<pair<int, int>> q;
        q.push({sr, sc});
        
        while (!q.empty()) {
            auto [i, j] = q.front(); q.pop();
            image[i][j] = color;
            
            if (i+1 < n && image[i+1][j] == init_color) {
                q.push({i+1, j});
            }
            if (j+1 < m && image[i][j+1] == init_color) {
                q.push({i, j+1});
            }
            if (i > 0 && image[i-1][j] == init_color) {
                q.push({i-1, j});
            }
            if (j > 0 && image[i][j-1] == init_color) {
                q.push({i, j-1});
            }
        }
        return image;
    }
};
