@attr = ();
$ans = 0;
while (<>){
    $ln  = length($_);
    if ($ln == 2){
        $ans = (scalar @attr) == 7 ? $ans + 1: $ans;
        @attr = ();
    }else {
        push(@attr, grep($_ ne "cid",map { (split ':', $_)[0] } (split ' ', $_)));
    }
}
$ans = (scalar @attr) == 7 ? $ans + 1: $ans;
print "$ans\n"