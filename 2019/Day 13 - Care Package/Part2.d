import std.stdio;
import std.array: array;
import std.conv;
import std.algorithm;
import std.string: split, strip;
import std.bigint;
import std.format: format;
import std.math;

import intcode;

void main(string[ ] args){
    auto file = File("in.txt", "r"); // put input in .txt file
    auto codes = file.readln().split(',').map!(x => BigInt(strip(x))).array;
    codes[0] = 2;
    Intcode intcode = new Intcode(codes);
    BigInt score = 0;
    BigInt xpaddle = 0;
    BigInt xball = 0;
    int input = 0;
    while(true){
        auto x = intcode.getNextOutput(input);
        if(-99999999 == x.toInt()) break; 
        auto y = intcode.getNextOutput(input);
        auto v = intcode.getNextOutput(input);

        if(x == -1 && y == 0) {
            score = v;
            continue;
        }else{
            if(v == 3){
                xpaddle = x;
            }
            if(v == 4){
                xball = x;
            }
        }
        //writeln(sgn((x - xpaddle).toInt()));
        input = sgn((xball - xpaddle).toInt());
    }

    writeln(score);

    file.close();
    
}