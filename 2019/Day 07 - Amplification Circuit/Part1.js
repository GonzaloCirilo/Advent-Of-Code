function findSignal(phaseSetting){
    var codes = [3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5];
    var ampCodes = [];
    for(var i = 0; i < 5; i++){
        ampCodes.push(Object.assign([], codes));
    }
    var nextInput = 0;
    for(var i = 0; i < 5; i++){
        var state = intCode(phaseSetting[i], nextInput, ampCodes[i], 0);
        nextInput = state[0];
    }
    return nextInput;
}

function solve1(){
    var phaseSetting = [0, 1, 2, 3, 4];
    var values = permutations(phaseSetting).map(x => findSignal(x));
    console.log(Math.max(...values));
}