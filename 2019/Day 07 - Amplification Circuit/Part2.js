function findSignalLoop(phaseSetting){
    var codes = [3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5 ];
    var ampCodes = [];
    var indMem = [];
    var outs = [[],[],[],[],[]];

    for(var i = 0; i < 5; i++){
        ampCodes.push(Object.assign([], codes));
        indMem.push(0)
    }

    var nextInput = 0;
    var auxi=0;
    while(true){
        var i = auxi % 5;
        var state = intCode((auxi<5)?phaseSetting[i]:nextInput, nextInput, ampCodes[i], indMem[i]);
        if(state[1] != -1){
            nextInput = state[0];
            indMem[i] = state[1];
            ampCodes[i] = state[2];
            phaseSetting[i] = state[3];
            outs[i].push(nextInput);
        }else{
            break;
        }    
        auxi++;
    }
    return outs[4][outs[4].length - 1];
}

function solve2(){
    var phaseSetting = [5, 6, 7, 8, 9];
    var values = permutations(phaseSetting).map(x => findSignalLoop(x));
    console.log(Math.max(...values));
}