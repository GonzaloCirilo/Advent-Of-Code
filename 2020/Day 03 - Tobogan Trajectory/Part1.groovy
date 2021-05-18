def solve(){
    Scanner sc = new Scanner(System.in);
    int lines = 0, mod = 0;
    def grid = [];
    while(sc.hasNextLine()){
        lines++;
        grid << sc.nextLine();
        mod = grid[lines-1].size();
    }

    int col = 0, ans = 0;
    for(i = 0; i < lines; i++){
        if(grid[i][col] == '#'){
            ans++;
        }
        col = (col + 3) % mod;
    }
    println(ans);
}

solve();