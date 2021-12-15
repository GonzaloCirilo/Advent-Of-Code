program part1;
Uses math;
var
    lr, rr, ls, rs, i, maxE, temp, j, count, it: Integer;
    cmd: packed array [1..10] of char;
    c: char;
    arr: array [0..800] of Integer;

begin
    maxE := 0;
    count := 0;
    while not eof do begin 
        lr := 0;
        ls := 0;
        rr := 127;
        rs := 7;
        readln(cmd);
        for it := 1 to 10 do begin
            case (cmd[it]) of
                'F': rr := Floor((rr + lr) / 2);// floor
                'B': lr := Ceil((rr + lr) / 2);// ceil
                'R': ls := Ceil((rs + ls) / 2);
                'L': rs := Floor((rs + ls) / 2);
            end;
        end;
        arr[count] := lr * 8 + ls;
        Inc(count)
    end;
    for i := 799 DownTo 0 do
		for j := 1 to i do
			if (arr[j-1] > arr[j]) then
			begin
				temp := arr[j-1];
				arr[j-1] := arr[j];
				arr[j] := temp;
			end;
    for i := 0 to 798 do
        if arr[i+1] - arr[i] = 2 then
            writeln(arr[i] + 1);
end.