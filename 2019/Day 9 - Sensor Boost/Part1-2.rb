$codes = STDIN.read.split(',').map(&:to_i)
$rb = 0

def getParameter(ind, mode)
    case mode
    when 0
        return $codes[$codes[ind]]
    when 1
        return $codes[ind]
    else
        return $codes[$rb + $codes[ind]]
    end
end

def setParameter(ind, mode, val)
    case mode
    when 0
        $codes[$codes[ind]] = val
    when 1
        $codes[ind] = val
    else
        $codes[$rb + $codes[ind]] = val
    end
end

def intcode(input)
    opsize = [4, 4, 2, 2, 3, 3, 4, 4, 2];
    i = 0
    while true do
        if $codes[i] % 100 == 99 then
            break
        end

        op = $codes[i] % 10

        mode = $codes[i].to_s.split('').map(&:to_i).reverse().drop(2)

        mode = mode.fill(0, mode.size, opsize[op - 1] - 1 - mode.size)

        case op
        when 1 
            a = getParameter(i + 1, mode[0])
            b = getParameter(i + 2, mode[1])
            setParameter(i + 3, mode[2], a + b)
            i += opsize[op-1]
        when 2
            a = getParameter(i + 1, mode[0])
            b = getParameter(i + 2, mode[1])
            setParameter(i + 3, mode[2], a * b)
            i += opsize[op-1]
        when 3
            setParameter(i + 1, mode[0], input)
            i += opsize[op-1]
        when 4
            puts getParameter(i + 1, mode[0])
            i += opsize[op-1]
        when 5
            a = getParameter(i + 1, mode[0])
            b = getParameter(i + 2, mode[1])
            i = a != 0 ? b : i + opsize[op - 1]
        when 6
            a = getParameter(i + 1, mode[0])
            b = getParameter(i + 2, mode[1])
            i = a == 0 ? b : i + opsize[op - 1]
        when 7
            a = getParameter(i + 1, mode[0])
            b = getParameter(i + 2, mode[1])
            setParameter(i + 3, mode[2], a < b ? 1 : 0)
            i += opsize[op-1]
        when 8
            a = getParameter(i + 1, mode[0])
            b = getParameter(i + 2, mode[1])
            setParameter(i + 3, mode[2], a == b ? 1 : 0)
            i += opsize[op-1]
        when 9
            $rb += getParameter(i + 1, mode[0])
            i += opsize[op-1]
        else
        end
    end
end

$aux = $codes
# Part 2
intcode(2)
$codes = $aux
puts "======"
$rb = 0
# Part 1
intcode(1)