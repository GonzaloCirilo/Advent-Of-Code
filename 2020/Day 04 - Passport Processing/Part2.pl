@attr = (); # note: add two new line to input
$ans = 0;
while (<>){
    $ln  = length($_);
    if ($ln == 2){
        @aux = grep($_ !~ m[cid], @attr);
        if ((scalar @aux) == 7){
            $valid = 1;
            foreach my $token (@attr) {
                $localValid = 0;
                if((split ':', $token)[0] eq "byr"){
                    if((split ':', $token)[1] =~ /^(19[2-9]\d|200[0-2])$/){
                        $localValid = 1;
                    }
                }elsif((split ':', $token)[0] eq "iyr"){
                    if((split ':', $token)[1] =~ /^(201[0-9]|2020)$/){
                        $localValid = 1;
                    }
                }elsif((split ':', $token)[0] eq "eyr"){
                    if((split ':', $token)[1] =~ /^(202[0-9]|2030)$/){
                        $localValid = 1;
                    }
                }elsif((split ':', $token)[0] eq "hgt"){
                    if((split ':', $token)[1] =~ /^((59|6\d|7[0-6])in|(1[5-8]\d|19[0-3])cm)/){
                        $localValid = 1;
                    }
                }
                elsif((split ':', $token)[0] eq "hcl"){
                    if(length((split ':', $token)[1]) == 7 && (split ':', $token)[1] =~ /#([0-9]|[a-f]){6}/){
                        $localValid = 1;
                    }
                }
                elsif((split ':', $token)[0] eq "ecl"){
                    if((split ':', $token)[1] =~ /(amb|blu|brn|gry|grn|hzl|oth)/){
                        $localValid = 1;
                    }
                }elsif((split ':', $token)[0] eq "pid"){
                    if(length((split ':', $token)[1]) == 9 && (split ':', $token)[1] =~ /\d{9}/){
                        $localValid = 1;
                    }
                }else{
                    next;
                }
                $valid = $valid & $localValid;
            }
            if($valid eq 1){
                $ans = $ans + 1;
            }
        }
        @attr = ();
    }else {
        push(@attr, (split ' ', $_));
    }
}
print "$ans\n"