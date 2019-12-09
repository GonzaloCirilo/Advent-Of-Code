<?php
    $codes = array(3,225,1,225,6,6,1100,1,238,225,104,0,2,171,209,224,1001,224,-1040,224,4,224,102,8,223,223,1001,224,4,224,1,223,224,223,102,65,102,224,101,-3575,224,224,4,224,102,8,223,223,101,2,224,224,1,223,224,223,1102,9,82,224,1001,224,-738,224,4,224,102,8,223,223,1001,224,2,224,1,223,224,223,1101,52,13,224,1001,224,-65,224,4,224,1002,223,8,223,1001,224,6,224,1,223,224,223,1102,82,55,225,1001,213,67,224,1001,224,-126,224,
        4,224,102,8,223,223,1001,224,7,224,1,223,224,223,1,217,202,224,1001,224,-68,224,4,224,1002,223,8,223,1001,224,1,224,1,224,223,223,1002,176,17,224,101,-595,224,224,4,224,102,8,223,223,101,2,224,224,1,224,223,223,1102,20,92,225,1102,80,35,225,101,21,205,224,1001,224,-84,224,4,224,1002,223,8,223,1001,224,1,224,1,224,223,223,1101,91,45,225,1102,63,5,225,1101,52,58,225,1102,59,63,225,1101,23,14,225,4,223,99,0,0,0,677,0,0,0,0,0,0,0,0,0,0,0,1105,0,
        99999,1105,227,247,1105,1,99999,1005,227,99999,1005,0,256,1105,1,99999,1106,227,99999,1106,0,265,1105,1,99999,1006,0,99999,1006,227,274,1105,1,99999,1105,1,280,1105,1,99999,1,225,225,225,1101,294,0,0,105,1,0,1105,1,99999,1106,0,300,1105,1,99999,1,225,225,225,1101,314,0,0,106,0,0,1105,1,99999,1008,677,677,224,1002,223,2,223,1006,224,329,101,1,223,223,1108,226,677,224,1002,223,2,223,1006,224,344,101,1,223,223,7,677,226,224,102,2,223,223,1006,
        224,359,1001,223,1,223,8,677,226,224,102,2,223,223,1005,224,374,1001,223,1,223,1107,677,226,224,102,2,223,223,1006,224,389,1001,223,1,223,1008,226,226,224,1002,223,2,223,1005,224,404,1001,223,1,223,7,226,677,224,102,2,223,223,1005,224,419,1001,223,1,223,1007,677,677,224,102,2,223,223,1006,224,434,1001,223,1,223,107,226,226,224,1002,223,2,223,1005,224,449,1001,223,1,223,1008,677,226,224,102,2,223,223,1006,224,464,1001,223,1,223,1007,677,226,224,1002,223,2,223,1005,
        224,479,1001,223,1,223,108,677,677,224,1002,223,2,223,1006,224,494,1001,223,1,223,108,226,226,224,1002,223,2,223,1006,224,509,101,1,223,223,8,226,677,224,102,2,223,223,1006,224,524,101,1,223,223,107,677,226,224,1002,223,2,223,1005,224,539,1001,223,1,223,8,226,226,224,102,2,223,223,1005,224,554,101,1,223,223,1108,677,226,224,102,2,223,223,1006,224,569,101,1,223,223,108,677,226,224,102,2,223,223,1006,224,584,1001,223,1,223,7,677,677,224,1002,223,2,223,1005,224,599,101,1,
        223,223,1007,226,226,224,102,2,223,223,1005,224,614,1001,223,1,223,1107,226,677,224,102,2,223,223,1006,224,629,101,1,223,223,1107,226,226,224,102,2,223,223,1005,224,644,1001,223,1,223,1108,677,677,224,1002,223,2,223,1005,224,659,101,1,223,223,107,677,677,224,1002,223,2,223,1006,224,674,1001,223,1,223,4,223,99,226);

    for($i = 0; $i < sizeof($codes); ){
        // opcode & modes
        $digits = str_split(strrev((string)$codes[$i]));
        $opcode = $digits[0];
        
        if(array_key_exists(1, $digits) && $digits[0] == "9" && $digits[1] == "9")
            break;

        if($opcode == "2" || $opcode == "1" || $opcode == "7" || $opcode == "8" || $opcode == "5" || $opcode == "6"){
            $p = [];
            for($j = 2; $j <= 3; $j++){
                if(array_key_exists($j, $digits) && $digits[$j] == "1")
                    $p[] = $codes[$i + $j - 1];
                else
                    $p[] = $codes[$codes[$i + $j - 1]];
            }
                       
            $target = $codes[$i + 3];

            if($opcode == "1")
                $codes[$target] = $p[0] + $p[1];
            elseif ($opcode == "2")
                $codes[$target] = $p[0] * $p[1];
            elseif ($opcode == "7")
                $codes[$target] = (int)($p[0] < $p[1]);
            elseif ($opcode == "8")
                $codes[$target] = (int)($p[0] == $p[1]);
            elseif ($opcode == "5"){
                if($p[0] != 0){
                    $i = $p[1];
                    continue;
                }
                $i = $i + 3; continue;
            } elseif ($opcode == "6"){
                if($p[0] == 0){
                    $i = $p[1];
                    continue;
                }
                $i = $i + 3; continue;
            }
            
            $i = $i + 4;
        } elseif ($opcode == "3" || $opcode == "4"){
            $a = 0;
            if(array_key_exists(2, $digits) && $digits[2] == "1")
                $a = $codes[$i + 1];
            else
                $a = $codes[$codes[$i + 1]];
            
            if($opcode == "3")
                $codes[$codes[$i + 1]] = 5;
            else 
                echo $a, '<br>';
            $i = $i + 2;
        }
    }
    echo 'finished';
?>