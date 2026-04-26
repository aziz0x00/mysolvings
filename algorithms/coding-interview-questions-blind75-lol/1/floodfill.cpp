class Solution {
public:
    vector<vector<int>> floodFill(vector<vector<int>>& image, int sr, int sc, int color) {

        if (image[sr][sc] != color)
            dfs(image, image[sr][sc], sr, sc, color, size(image), size(image[0]));

        return image;
    }

    void dfs(vector<vector<int>> & image, int startColor, int i, int j, int color, int n, int m) {

        image[i][j] = color;

        if (i+1 < n && image[i+1][j] == startColor)
            dfs(image, startColor, i+1, j, color, n, m);
        if (j+1 < m && image[i][j+1] == startColor)
            dfs(image, startColor, i, j+1, color, n, m);
        if (i-1 >= 0 && image[i-1][j] == startColor)
            dfs(image, startColor, i-1, j, color, n, m);
        if (j-1 >= 0 && image[i][j-1] == startColor)
            dfs(image, startColor, i, j-1, color, n, m);
    }
};
