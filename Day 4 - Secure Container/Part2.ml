(*Fucntion for generating range*)
let range a b =
  List.init (b - a) ((+) a);;

let counter = ref 0;;

let digits n = 
  let rec ite n arr =
    if n = 0 then arr (* return the array*)
    else ite (n/10) (n mod 10:: arr) in
  match n with
  | 0 -> [0]
  | _ -> ite n [];;

let rec _validate arr nZeros mini = 
  if (List.length arr) == 1 then !mini == 1 || !nZeros == 1
  else begin 
    if ((List.nth arr 0) - (List.nth arr 1)) == 0 then incr nZeros;
    if ((List.nth arr 0) - (List.nth arr 1)) < 0 && !nZeros != 0 then begin 
      if mini > nZeros then mini := !nZeros;
      nZeros := 0;      
    end;
    
    (_validate (List.tl arr) nZeros mini) && ((List.nth arr 0) - (List.nth arr 1)) <= 0;
  end;;

let validate n = 
  _validate (digits n) (ref 0) (ref 7);;

let f elem =
    if (validate elem) then  incr counter;  in
  List.iter f (range 123257 647015);;

Printf.printf "Answer is %d\n" !counter
