module intcode;
import std.array: array;
import std.bigint;
import std.format: format;
import std.string: split;
import std.range: drop;
import std.algorithm;
import std.conv;
import std.stdio;

class Intcode {
    BigInt[] codes;
    int rb; 
    int i;

    this(BigInt[] _codes){
        codes = _codes;
        for(int j = 0; j < 1000; j++){
            codes ~= BigInt(0);
        }
    }

    BigInt getParam(int ind, int mode){
        switch(mode){
            case 0:
                return codes[codes[ind].toLong()]; break;
            case 1:
                return codes[ind]; break;
            case 2:
                return codes[rb + codes[ind].toLong()]; break;
            default: break;
        }
        return BigInt(0);
    }

    void setParam(int ind, int mode, BigInt val){
        switch(mode){
            case 0: 
                codes[codes[ind].toLong()] = val; break;
            case 1: 
                codes[ind] = val; break;
            case 2: 
                codes[rb + codes[ind].toLong()] = val; break;
            default: break;
        }
    }

    BigInt getNextOutput(int input){
        // intcode logic
        auto opsize = [4, 4, 2, 2, 3, 3, 4, 4, 2];
        
        while(true){
            if(codes[i] % 100 == 99){
                return BigInt(-99999999); break;
            }
            auto op = codes[i] % 10;
            auto mode = format("%d",codes[i]).split("").map!(x => to!int(x)).array().reverse.drop(2);
            auto size = mode.length;
            for(int j = 0; j< opsize[op - 1] - 1 - size; j++){
                mode ~= 0;
            }

            switch(op){
                case 1:
                    auto a = getParam(i + 1, mode[0]);
                    auto b = getParam(i + 2, mode[1]);
                    setParam(i + 3, mode[2], a + b);
                    i += opsize[op - 1];
                    break;
                case 2:
                    auto a = getParam(i + 1, mode[0]);
                    auto b = getParam(i + 2, mode[1]);
                    setParam(i + 3, mode[2], a * b);
                    i += opsize[op - 1];
                    break;
                case 3:
                    setParam(i + 1, mode[0], BigInt(input));
                    i += opsize[op - 1];
                    break;
                case 4:
                    auto a = getParam(i + 1, mode[0]);
                    i += opsize[op - 1];
                    return a;
                    break;
                case 5:
                    auto a = getParam(i + 1, mode[0]);
                    auto b = getParam(i + 2, mode[1]);
                    i = (a != BigInt(0)) ? b.toInt() : i + opsize[op - 1];
                    break;
                case 6:
                    auto a = getParam(i + 1, mode[0]);
                    auto b = getParam(i + 2, mode[1]);
                    i = (a == 0) ? b.toInt() : i + opsize[op - 1];
                    break;
                case 7:
                    auto a = getParam(i + 1, mode[0]);
                    auto b = getParam(i + 2, mode[1]);
                    setParam(i + 3, mode[2], BigInt((a < b) ? 1 : 0));
                    i += opsize[op - 1];
                    break;
                case 8:
                    auto a = getParam(i + 1, mode[0]);
                    auto b = getParam(i + 2, mode[1]);
                    setParam(i + 3, mode[2], BigInt((a == b) ? 1 : 0));
                    i += opsize[op - 1];
                    break;
                case 9:
                    rb += getParam(i + 1, mode[0]).toLong();
                    i += opsize[op - 1];
                    break;
                default: break;
            }            
        }
        return BigInt(0);
    }
}