import std.stdio;
import std.array: array;
import std.conv;
import std.algorithm;
import std.string: split, strip;
import std.bigint;
import std.format: format;
import intcode;

void main(string[ ] args){
    auto file = File("in.txt", "r"); // put input in .txt file
    auto codes = file.readln().split(',').map!(x => BigInt(strip(x)));
    Intcode intcode = new Intcode(codes.array);
    int ans = 0;

    while(true){
        auto x = intcode.getNextOutput(0);
        if(-99999999 == x.toInt()) break; 
        auto y = intcode.getNextOutput(0);
        auto v = intcode.getNextOutput(0);

        if(v == 2) ans++;
    }

    writeln(ans);
    file.close();
}