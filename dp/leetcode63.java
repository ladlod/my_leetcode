package dp;

public class leetcode63 {
    public int uniquePathsWithObstacles(int[][] obstacleGrid) {
        int m = obstacleGrid.length;
        int n = obstacleGrid[0].length;
        int[][] res = new int[m][n];
        int cango = 1;
        for (int i = 0; i < m; i++) {
            if(obstacleGrid[i][0] == 1)
                cango = 0;            
            res[i][0] = cango;
        }
        cango = 1;
        for (int i = 0; i < n; i++) {
            if(obstacleGrid[0][i] == 1)
                cango = 0;
            res[0][i] = cango;
        }
        for (int i = 1; i < m; i++) {
            for (int j = 1; j < n; j++) {
                if(obstacleGrid[i][j] == 1)
                    res[i][j] = 0;
                else
                    res[i][j] = res[i - 1][j] + res[i][j - 1];
            }
        }

        return res[m - 1][n - 1];
    }
}
