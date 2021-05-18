def getNumTrees(dx, dy, g, mod){
    int col = 0, count = 0;
    for(i = 0; i < g.size(); i += dy){
        if(g[i][col] == '#'){
            count++;
        }
        col = (col + dx) % mod;
    }
    return count;
}

def solve(){
    Scanner sc = new Scanner(System.in);
    int lines = 0, mod = 0;
    def grid = [];
    while(sc.hasNextLine()){
        lines++;
        grid << sc.nextLine();
        mod = grid[lines-1].size();
    }
    def slopes = [new Tuple2(1,1), new Tuple2(3,1), new Tuple2(5, 1), new Tuple2(7, 1), new Tuple2(1, 2)];
    println(slopes.size());
    long ans = 1;
    for(j = 0; j < slopes.size(); j++){
        ans *= getNumTrees(slopes[j].first, slopes[j].second, grid, mod);
    }
    
    println(ans);
}

solve();