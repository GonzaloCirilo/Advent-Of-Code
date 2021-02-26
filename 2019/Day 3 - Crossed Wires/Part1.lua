
-- Split String
function splitString (inputstr, sep)
    local t={}
    for str in string.gmatch(inputstr, "([^" .. sep .. "]+)") do
            table.insert(t, str)
    end
    return t
end

grid = {}

minDistance = 100000


function createIfNotExist(x, y)
    if grid[y] == nil then
        grid[y] = {}
    end
    if grid[y][x] == nil then
         grid[y][x] = 0
    end
end

-- solution
threshold = 1
for line in io.lines("input.txt") do 
    sx = 17500 ; sy = 17500
    --print("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
    a = splitString(line, ',')
    for n, w in ipairs(a) do
        l = string.match(w, "%d+")
        c = string.sub(w,1,1)
        dif = 0
        if c == 'U' or c == 'L' then
            dif = -1
        else
            dif = 1
        end
        --print("Move " .. l .. " positions " .. c)
        for i = 1, l, 1 do
            if c == 'U' or c == 'D' then
                --print((sy + (i*dif)).. " " ..sx)
                createIfNotExist(sx, sy + (i*dif))
                if grid[sy + (i*dif)][sx] == 1 and threshold == 2 then
                    --print(math.abs(17500 - (sy + (i*dif))) + math.abs(17500 - sx))
                    minDistance = math.min(minDistance, (math.abs(17500 - (sy + (i*dif))) + math.abs(17500 - sx)))
                end
                grid[sy + (i*dif)][sx] = (grid[sy + (i*dif)][sx] + threshold)
            else
                --print(sy .. " " .. (sx + (i*dif)))
                createIfNotExist(sx + (i*dif), sy)
                if grid[sy][sx + (i*dif)] == 1 and threshold == 2 then
                    --print(math.abs(17500 - sy) + math.abs(17500 - (sx + (i*dif))))
                    minDistance = math.min(minDistance, (math.abs(17500 - sy) + math.abs(17500 - (sx + (i*dif)))))
                end
                grid[sy][sx + (i*dif)] = (grid[sy][sx + (i*dif)] + threshold)                
            end
        end
        if c == 'U' then sy = sy - l
        elseif c == 'D' then sy = sy + l
        elseif c == 'R' then sx = sx + l
        elseif c == 'L' then sx = sx - l
        end
    end
    threshold = threshold + 1
end

print(minDistance)
