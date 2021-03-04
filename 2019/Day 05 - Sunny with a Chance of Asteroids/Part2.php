<?php
    $codes = array(3,9,8,9,10,9,4,9,99,-1,8);

    for($i = 0; $i < sizeof($codes); ){
        // opcode & modes
        $digits = str_split(strrev((string)$codes[$i]));
        $opcode = $digits[0];
        
        if(array_key_exists(1, $digits) && $digits[0] == "9" && $digits[1] == "9")
            break;

        if($opcode == "2" || $opcode == "1" || $opcode == "7" || $opcode == "8" || $opcode == "5" || $opcode == "6"){
            $p = [];
            for($j = 2; $j <= 3; $j++){
                if(array_key_exists($j, $digits) && $digits[$j] == "1") $p[] = $codes[$i + $j - 1];
                else $p[] = $codes[$codes[$i + $j - 1]];
            }
                       
            $target = $codes[$i + 3];

            if($opcode == "1") $codes[$target] = $p[0] + $p[1];
            elseif ($opcode == "2") $codes[$target] = $p[0] * $p[1];
            elseif ($opcode == "7") $codes[$target] = (int)($p[0] < $p[1]);
            elseif ($opcode == "8") $codes[$target] = (int)($p[0] == $p[1]);
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
            if(array_key_exists(2, $digits) && $digits[2] == "1") $a = $codes[$i + 1];
            else $a = $codes[$codes[$i + 1]];
            
            if($opcode == "3") $codes[$codes[$i + 1]] = 5;
            else  echo $a, '<br>';
            $i = $i + 2;
        }
    }
    echo 'finished';
?>