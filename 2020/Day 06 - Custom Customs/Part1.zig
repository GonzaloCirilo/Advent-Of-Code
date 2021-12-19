const std = @import("std");

pub fn main() !void {
    const stdin = std.io.getStdOut().reader();
    const stdout = std.io.getStdOut().writer();

    var file = try std.fs.cwd().openFile("in6.txt", .{});
    defer file.close();

    var buf_reader = std.io.bufferedReader(file.reader());
    var in_stream = buf_reader.reader();
    var buf: [1024]u8 = undefined;

    var dict: [26]i32 = [_]i32{0} ** 26;
    var ans: i32 = 0; 
    var prev: bool = false;
    while(try in_stream.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        try stdout.print("{s}\n", .{line});


        if(line.len == 0){
            if(prev) break;
            var aux: i32 = 0;
            for(dict) |count|{
                if(count > 0){
                    aux += 1;
                }
            }
            ans += aux;
            prev = true;
            dict = [_]i32{0} ** 26;
        }else{
            prev = false;
            for(line) |c| {
                dict[c-'a'] += 1;
            }
        }
    }

    try stdout.print("{}\n", .{ans});
}