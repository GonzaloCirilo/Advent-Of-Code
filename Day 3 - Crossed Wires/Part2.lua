-- Split String
function splitString (inputstr, sep)
    local t={}
    for str in string.gmatch(inputstr, "([^" .. sep .. "]+)") do
            table.insert(t, str)
    end
    return t
end
-- matrix for storing the circuit
grid = {}

steps = {}

minSteps = 100000


function createIfNotExist(x, y)
    if grid[y] == nil then
        grid[y] = {}
    end
    if grid[y][x] == nil then
         grid[y][x] = 0
    end
end

function createStepIfNotExists(x, y)
    if grid[17500 * x + y] == nil then
        grid[17500 * x + y] = 0
    end
end

-- solution
threshold = 1
for line in io.lines("input.txt") do 
    sx = 17500 ; sy = 17500
    st = 0
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
        for i = 1, l, 1 do
            if c == 'U' or c == 'D' then
                createIfNotExist(sx, sy + (i*dif))
                -- we found a crossing point
                if grid[sy + (i*dif)][sx] == 1 and threshold == 2 then
                    -- new logic
                    minSteps = math.min(minSteps, steps[17500 * sx + sy + (i*dif)] + st + i)
                end
                -- calculate the steps for
                if threshold == 1 then
                    createStepIfNotExists(sx, sy + (i*dif))
                    steps[17500 * sx + sy + (i*dif)] = st + i
                end
                grid[sy + (i*dif)][sx] = (grid[sy + (i*dif)][sx] + threshold)
            else
                createIfNotExist(sx + (i*dif), sy)
                -- we found a crossing point
                if grid[sy][sx + (i*dif)] == 1 and threshold == 2 then
                    -- new logic
                    minSteps = math.min(minSteps, steps[17500 * (sx + (i*dif)) + sy] + st + i)
                end
                -- calculate the steps for
                if threshold == 1 then
                    createStepIfNotExists(sx + (i*dif), sy)
                    steps[17500 * (sx + (i*dif)) + sy] = st + i
                end
                grid[sy][sx + (i*dif)] = (grid[sy][sx + (i*dif)] + threshold)                
            end
        end
        if c == 'U' then sy = sy - l
        elseif c == 'D' then sy = sy + l
        elseif c == 'R' then sx = sx + l
        elseif c == 'L' then sx = sx - l
        end
        st = st + l
    end
    threshold = threshold + 1
end

print(minSteps)