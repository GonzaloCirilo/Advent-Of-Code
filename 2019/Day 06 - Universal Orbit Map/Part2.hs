--8.0.2
import Data.List
import Data.Char
import Data.HashMap.Strict
import Data.HashSet
import Data.Maybe

split :: String -> Char -> [String]
split "" _ = []
split xs c = let (ys, zs) = break (== c) xs
    in  if Data.List.null zs then [ys] else ys : split (tail zs) c

makeArray :: String -> String -> HashMap String  [String] -> [String]
makeArray a b graph = b:Data.HashMap.Strict.lookupDefault [] a graph         

inputToGraph :: IO(HashMap String [String])
inputToGraph = do
    line <- getLine
    if Data.List.null line
        then return (Data.HashMap.Strict.empty)
    else 
        do 
            let s = split line ')'
            aux <- inputToGraph
            let aux2 = (Data.HashMap.Strict.insert (s!!0) (makeArray (s!!0) (s!!1) aux) aux) -- u to v
            return (Data.HashMap.Strict.insert (s!!1) (makeArray (s!!1) (s!!0) aux2) aux2) -- v to u

-- Params: current node, depth, graph, depth hash table          
bfs :: String -> Int -> (HashMap String [String]) -> (HashMap String Int) -> (HashMap String Int)
bfs n d graph depths = do
    let dests = fromMaybe []  (Data.HashMap.Strict.lookup n graph)
    let visited = Data.HashMap.Strict.keys depths
    let av = (dests \\ visited)
    if Data.List.null av
        then (Data.HashMap.Strict.insert n d depths)
    else 
        do
            (itNodes av (d + 1) graph depths)

itNodes :: [String] -> Int -> (HashMap String [String]) -> (HashMap String Int) -> (HashMap String Int)
itNodes [] _ _ _ = Data.HashMap.Strict.empty
itNodes (a:av) d graph depths = do 
    let aux = itNodes av d graph depths
    (Data.HashMap.Strict.union (bfs a d graph (Data.HashMap.Strict.insert a d depths)) aux)

main = do
    paths <- inputToGraph

    let table = bfs "YOU" 0 paths (Data.HashMap.Strict.fromList [("YOU", 0)])

    let ans = (table!"SAN") - 2
    print ans
    