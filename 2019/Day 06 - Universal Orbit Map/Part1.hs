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
                        
inputToGraph :: IO(HashMap String  [String])
inputToGraph = do
    line <- getLine
    if Data.List.null line
        then return (Data.HashMap.Strict.empty)
    else 
        do 
            let s = split line ')'
            aux <- inputToGraph
            return (Data.HashMap.Strict.insert (s!!0) (makeArray (s!!0) (s!!1) aux) aux)

-- BFS Params: depth, current node, graph            
getOrbitCount :: Int -> String -> HashMap String [String] -> Int
getOrbitCount d n graph = do
    let dests = Data.HashMap.Strict.lookup n graph
    if isNothing dests
        then d
    else 
        do
            let sum = sumOrbitCount (fromMaybe [] dests) (d + 1) graph
            (d + sum)

sumOrbitCount :: [String] -> Int ->  HashMap String [String] -> Int
sumOrbitCount [] _ _ = 0
sumOrbitCount (x:xs) d g = (getOrbitCount d x g) + (sumOrbitCount xs d g)

main = do
    paths <- inputToGraph
    let us = Data.HashMap.Strict.keys paths
    let vs =  concat ( Data.HashMap.Strict.elems paths )
    let roots = us \\ vs

    -- starting from roots
    let ans = sumOrbitCount roots 0 paths
    print ans
    