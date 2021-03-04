import 'dart:io';
import 'dart:convert';
import 'Intcode.dart';

void main(){
  List<int> codes = stdin.readLineSync().split(",").map((x) => int.parse(x)).toList();
  codes.addAll(new List<int>.filled(1000, 0));// Add some extra memory
  var rb = 0;
  Intcode intcode = new Intcode(codes, rb);  
  print(intcode.run(0));
}
