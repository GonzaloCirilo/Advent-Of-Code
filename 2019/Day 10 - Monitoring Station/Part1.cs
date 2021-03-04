using System;

public class Part1{

    public static void Main() {
        String[] grid = new String[200];
        int k = 0;

        Util.getInput(ref grid, ref k);

        var result = Util.getAsteroids(grid, k);
        Console.WriteLine(result.Item1);
	}
}