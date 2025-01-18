/*
You are given an m x n integer array grid where grid[i][j] could be:

1 representing the starting square. There is exactly one starting square.
2 representing the ending square. There is exactly one ending square.
0 representing empty squares we can walk over.
-1 representing obstacles that we cannot walk over.
Return the number of 4-directional walks from the starting square to the ending square, that walk over every non-obstacle square exactly once.
*/

enum SquareTypes {
    Visited = 3,
    Start = 1,
    End = 2,
    Empty = 0,
    Obstacle = -1,
}

const dirs = [[1, 0], [0, 1], [-1, 0], [0, -1]];

function findStart(grid: number[][]): [number, number] {
    for (let x = 0; x < grid.length; x += 1) {
        const row = grid[x];
        for (let y = 0; y < row.length; y += 1) {
            if (row[y] === SquareTypes.Start) {
                return [x, y];
            }
        }
    }
    return [-1, -1];
};

function outOfBounds(grid: number[][], pos: number[]) {
    const x = pos[0];
    const y = pos[1];

    const maxX = grid.length;
    const maxY = grid.at(0) ? grid[0].length : 0;

    const inBounds = 0 <= x && x < maxX && 0 <= y && y < maxY
    return !inBounds;
}

function hasVisitedAll(visited: number[][]) {
    for (let x = 0; x < visited.length; x += 1) {
        const row = visited[x];
        for (let y = 0; y < row.length; y += 1) {
            if (row[y] === SquareTypes.Empty) {
                return false;
            }
        }
    }
    return true;
}

function findPathToEnd(grid: number[][], pos: number[], visited: number[][]) {
    const x = pos[0];
    const y = pos[1];
    const square = grid[x][y];

    if (square === SquareTypes.End) {
        return hasVisitedAll(visited) ? 1 : 0;
    }

    visited[x][y] = SquareTypes.Visited;

    if (square === SquareTypes.Obstacle) {
        return 0;
    }

    let count = 0;
    for (let i = 0; i < dirs.length; i += 1) {
        const newX = pos[0] + dirs[i][0]
        const newY = pos[1] + dirs[i][1]
        if (!outOfBounds(grid, [newX, newY]) && visited[newX][newY] !== SquareTypes.Visited) {
            count += findPathToEnd(grid, [newX, newY], structuredClone(visited));
        }
    }
    return count;
}

function uniquePathsIII(grid: number[][]): number {
    let visited = structuredClone(grid);
    const start = findStart(grid);
    visited[start[0]][start[1]] = SquareTypes.Visited;
    return findPathToEnd(grid, start, visited);
};

let grid = [[0,1],[2,0]] 
console.log('paths:', uniquePathsIII(grid)) // 0

grid = [[1,0,0,0],[0,0,0,0],[0,0,2,-1]] 
console.log('paths:', uniquePathsIII(grid)) // 2

grid = [[1, 0, 0, 0], [0, 0, 0, 0], [0, 0, 0, 2]]
console.log('paths:', uniquePathsIII(grid)) // 4

