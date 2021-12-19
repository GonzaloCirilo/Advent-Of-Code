(defn getInput []
    (doall (map #(Integer/parseInt %) (line-seq (java.io.BufferedReader. *in*))))) ;; reading lines from stdin

(def arr (getInput))
(println (reduce 1 (distinct (for 
        [a arr
         b arr
         c arr
         :let [sum (+ a b c)]
         :when (= sum 2020)] (* a b c)))))