const permutations = arr => { /*taken from https://www.30secondsofcode.org/js/s/permutations*/
    if (arr.length <= 2) return arr.length === 2 ? [arr, [arr[1], arr[0]]] : arr;
    return arr.reduce(
      (acc, item, i) =>
        acc.concat(
          permutations([...arr.slice(0, i), ...arr.slice(i + 1)]).map(val => [
            item,
            ...val,
          ])
        ),
      []
    );
  };

function intCode(phase, inputCode, codes, ind){
    var opsize = [4, 4, 2, 2, 3, 3, 4, 4];

    var input = phase;
    while(true){
        if(codes[ind] % 100 == 99) break;

        var op = codes[ind] % 10;
        var modes = [];
        for(var i = 0; i < opsize[op - 1]; i++){
            if(i == 0){
                modes.push(0); modes.push(0);
            }else{
                modes.push((parseInt(codes[ind]/Math.pow(10, i + 1))) % 10);
            }
        }
        switch(op){
            case 1:
                var a = modes[2] == 0? codes[codes[ind + 1]] : codes[ind + 1];
                var b = modes[3] == 0? codes[codes[ind + 2]] : codes[ind + 2];
                codes[codes[ind + 3]] = a + b;
                ind += opsize[op - 1];
                break;
            case 2:
                var a = modes[2] == 0? codes[codes[ind + 1]] : codes[ind + 1];
                var b = modes[3] == 0? codes[codes[ind + 2]] : codes[ind + 2];
                codes[codes[ind + 3]] = a * b;
                ind += opsize[op - 1];
                break;
            case 3:
                codes[codes[ind + 1]] = input;
                ind += opsize[op - 1];
                input = inputCode;
                break;
            case 4:
                var out = modes[2] == 0? codes[codes[ind+1]] : codes[ind+1];
                ind += opsize[op - 1];
                return [out, ind, codes, input];
                break;
            case 5:
                var a = modes[2] == 0? codes[codes[ind + 1]] : codes[ind + 1];
                var b = modes[3] == 0? codes[codes[ind + 2]] : codes[ind + 2];
                if(a != 0){
                    ind = b;
                }else{
                    ind += opsize[op - 1];
                }
                break;
            case 6:
                var a = modes[2] == 0? codes[codes[ind + 1]] : codes[ind + 1];
                var b = modes[3] == 0? codes[codes[ind + 2]] : codes[ind + 2];
                if(a == 0){
                    ind = b;
                }else{
                    ind += opsize[op - 1];
                }
                break;
            case 7:
                var a = modes[2] == 0? codes[codes[ind + 1]] : codes[ind + 1];
                var b = modes[3] == 0? codes[codes[ind + 2]] : codes[ind + 2];
                codes[codes[ind + 3]] = a < b ? 1 : 0;
                ind += opsize[op - 1];
                break;
            case 8:
                var a = modes[2] === 0? codes[codes[ind + 1]] : codes[ind + 1];
                var b = modes[3] === 0? codes[codes[ind + 2]] : codes[ind + 2];
                codes[codes[ind + 3]] = a == b ? 1 : 0;
                ind += opsize[op - 1];
                break;
        }
    }
    return [0, -1, [], 0];
}