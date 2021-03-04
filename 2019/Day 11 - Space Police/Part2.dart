import 'dart:io';
import 'dart:convert';
import 'dart:math';
import 'Intcode.dart';

void main(){
  List<int> codes = stdin.readLineSync().split(",").map((x) => int.parse(x)).toList();
  codes.addAll(new List<int>.filled(1000, 0));// Add some extra memory
  var rb = 0;
  Intcode intcode = new Intcode(codes, rb);  
  print(intcode.run(1));
  //print(intcode.visited);
  for(int i = -100; i < 100; i++) {
    for(int j = -40; j < 85; j++){
      if(intcode.visited.containsKey(Point(j,i))){
        stdout.write(intcode.visited[Point(j,i)] == 1 ? '#' : '.');
      }else{
        stdout.write('.');
      }
    }
    stdout.write('\n');
  }
}
