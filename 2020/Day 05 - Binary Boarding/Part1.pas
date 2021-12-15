program part1;
Uses math;
var
    lr, rr, ls, rs, i, maxE: Integer;
    cmd: packed array [1..10] of char;
    c: char;

begin
    maxE := 0;
    while not eof do begin 
        lr := 0;
        ls := 0;
        rr := 127;
        rs := 7;
        readln(cmd);
        for i := 1 to 10 do begin
            case (cmd[i]) of
                'F': rr := Floor((rr + lr) / 2);// floor
                'B': lr := Ceil((rr + lr) / 2);// ceil
                'R': ls := Ceil((rs + ls) / 2);
                'L': rs := Floor((rs + ls) / 2);
            end;
        end;
        maxE := Max(maxE, lr * 8 + ls);
    end;
    writeln(maxE);
end.