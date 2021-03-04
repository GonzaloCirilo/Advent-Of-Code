using System;
using System.Collections.Generic;
using System.Linq;

public class Util {
    public static int gcd(int n, int m){
        while(n != 0){
            int aux = m % n;
            m = n;
            n = aux;
        }
        return m;
    }

    public static Double getAngle(int x, int y){
        return Math.Atan2(x, y) * 180.0 / Math.PI + ((x < 0) ? 360.0 : 0.0);
    }

    public static void getInput(ref String[] grid, ref int k){
        String line;
        while((line = Console.ReadLine()) != null && line != ""){
            grid[k++] = line;
        }
    }

    public static Tuple<List<int>,int,int> getAsteroids(String[] grid, int k){
        List<int> xs = new List<int>(), ys =  new List<int>(), asteroids = new List<int>();

        for(int i = 0; i < k; i++)
            for(int j = 0; j < grid[i].Length; j++){
                if(grid[i][j] == '#'){
                    xs.Add(j);
                    ys.Add(i);
                    asteroids.Add(0);
                }
            }

        for(int i = 0; i < asteroids.Count; i++){
            HashSet<Tuple<int, int>> m = new HashSet<Tuple<int,int>>();
            for(int j = i + 1; j < asteroids.Count; j++){
                if(xs[j] == xs[i]){
                    if(m.Add(new Tuple<int,int>(999999, 999999))){
                        asteroids[i]++;
                        asteroids[j]++;
                    }
                }else{
                    int num = ys[j] - ys[i], dem = xs[j] - xs[i]; 
                    if(m.Add(new Tuple<int,int>(num / Util.gcd(num, dem), dem / Util.gcd(num, dem)))){
                        asteroids[i]++;
                        asteroids[j]++;
                    }
                }
            }
        }
        int ind = asteroids.IndexOf(asteroids.Max()), tx = xs[ind], ty = ys[ind];
        return Tuple.Create(asteroids.Max(), tx, ty);
    }
}