import 'dart:convert';
import 'dart:math';

class Intcode {
  var codes;
  int rb;
  Map<Point, int> visited = {};

  Intcode(var codes, int rb){
    this.codes = codes;
    this.rb = rb;
  }

  int getParameter(ind, mode){
    switch(mode){
      case 0: {
        return this.codes[this.codes[ind]];
      }
      break;
      case 1: {
        return this.codes[ind];
      }
      break;
      default: {
        return this.codes[this.rb + this.codes[ind]];
      }
      break;
    }
  }

  void setParameter(ind, mode, val){
    switch(mode){
      case 0: {
        this.codes[this.codes[ind]] = val;
      }
      break;
      case 1: {
        this.codes[ind] = val;
      }
      break;
      default: {
        this.codes[this.rb + this.codes[ind]] = val;
      }
      break;
    }
  }

  int run(input){
    // Robot Logic
    int x = 0, y = 0, action = 0, currentColor = 0, ans = 0, dir = 0;
    var dx = [0, 1, 0, -1];
    var dy = [-1, 0, 1, 0];
    // intcode logic
    var opsize = [4, 4, 2, 2, 3, 3, 4, 4, 2];
    var i = 0;

    while(true){
      if(this.codes[i] % 100 == 99){
        break;
      }

      var op = this.codes[i] % 10;
      var mode = this.codes[i].toString().split("").map((x) => int.parse(x)).toList().reversed.toList();
      mode.removeAt(0); if(mode.length > 0) mode.removeAt(0);

      mode.addAll(new List<int>.filled(opsize[op - 1] - mode.length - 1, 0));

      switch(op){
        case 1: {
          var a = getParameter(i + 1, mode[0]);
          var b = getParameter(i + 2, mode[1]);
          setParameter(i + 3, mode[2], a + b);
          i += opsize[op-1];
        }
        break;
        case 2: {
          var a = getParameter(i + 1, mode[0]);
          var b = getParameter(i + 2, mode[1]);
          setParameter(i + 3, mode[2], a * b);
          i += opsize[op-1];
        }
        break;
        case 3: {
          var pos = Point(x, y);
          setParameter(i + 1, mode[0], visited.containsKey(pos)? visited[pos] : input);
          input = 0;
          i += opsize[op-1];
        }
        break;
        case 4: { // NOTE if needed return at this point
          var result = getParameter(i + 1, mode[0]);
          if(action % 2 == 0){// paint
            var pos = Point(x, y);
            if(!visited.containsKey(pos)){
              ans++;
            }
            visited[pos] = result;
          }else{// moves
            dir = ((result == 1 ? 1 : -1) + dir  + 4) % 4;
            x += dx[dir];
            y += dy[dir]; 
          }
          action++;
          i += opsize[op-1];
        }
        break;
        case 5: {
          var a = getParameter(i + 1, mode[0]);
          var b = getParameter(i + 2, mode[1]);
          i = (a != 0) ? b : i + opsize[op - 1];
        }
        break;
        case 6: {
          var a = getParameter(i + 1, mode[0]);
          var b = getParameter(i + 2, mode[1]);
          i = (a == 0) ? b : i + opsize[op - 1];
        }
        break;
        case 7: {
          var a = getParameter(i + 1, mode[0]);
          var b = getParameter(i + 2, mode[1]);
          setParameter(i + 3, mode[2], (a < b) ? 1 : 0);
          i += opsize[op-1];
        }
        break;
        case 8: {
          var a = getParameter(i + 1, mode[0]);
          var b = getParameter(i + 2, mode[1]);
          setParameter(i + 3, mode[2], (a == b) ? 1 : 0);
          i += opsize[op-1];
        }
        break;
        case 9: {
          this.rb += getParameter(i + 1, mode[0]);
          i += opsize[op-1];
        }
        break;
      }
    }
    return ans;
  }
}