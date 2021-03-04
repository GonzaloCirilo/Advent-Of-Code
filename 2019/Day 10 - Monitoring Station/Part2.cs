using System;
using System.Collections.Generic;

public class Part2 {

    public static void Main(){
        String[] grid = new String[200];
        int k = 0;

        Util.getInput(ref grid, ref k);

        var result = Util.getAsteroids(grid ,k);
        int tx = result.Item2, ty = result.Item3, vaporized = 0;
        List<int> xs = new List<int>(), ys =  new List<int>();
        
        for(int i = 0; i < k; i++){//reorder to evaluate from epicenter (tx, ty)
            int indx = i;
            if(i <= ty)
                indx = ty - i;

            for(int j = 0; j < grid[indx].Length; j++){
                if(grid[indx][j] == '#'){
                    xs.Add(j);
                    ys.Add(indx);
                }
            }
        }

        while(true){
            if(vaporized > 200) break;
            HashSet<Tuple<int, int>> m = new HashSet<Tuple<int,int>>(), m2 = new HashSet<Tuple<int,int>>();
            List<int> toRemove = new List<int>();
            for(int j = 0; j < xs.Count; j++){
                if(xs[j] == tx && ys[j] == ty) continue;
                if(xs[j] == tx){
                    if(ys[j] < ty){
                        if(m2.Add(new Tuple<int,int>(999999, 999999)))
                            toRemove.Add(j);
                    }else{
                        if(m.Add(new Tuple<int,int>(999999, 999999)))
                            toRemove.Add(j);
                    }

                }else{
                    int num = ys[j] - ty, dem = xs[j] - tx; 
                    if(xs[j] < tx){
                        if(m2.Add(new Tuple<int,int>(num / Util.gcd(num, dem), dem / Util.gcd(num, dem))))
                            toRemove.Add(j);
                    }else{
                        if(m.Add(new Tuple<int,int>(num / Util.gcd(num, dem), dem / Util.gcd(num, dem))))
                            toRemove.Add(j);
                    }
                }
            }

            toRemove.Sort((a, b) => Util.getAngle(xs[a] - tx, ty - ys[a]).CompareTo(Util.getAngle(xs[b] - tx, ty - ys[b])));

            for(int j = 0; j < toRemove.Count; j++){
                vaporized++;
                if(vaporized == 200){
                    Console.WriteLine(xs[toRemove[j]] * 100 + ys[toRemove[j]]);
                }                
            }
            List<int> axs = new List<int>(), ays =  new List<int>();
            for(int j = 0; j < xs.Count; j++){
                if(!toRemove.Contains(j)){
                    axs.Add(xs[j]);
                    ays.Add(ys[j]);
                }
            }
            xs = axs;
            ys = ays;
        }
	}
}