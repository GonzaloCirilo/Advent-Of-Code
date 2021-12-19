(defn getInput []
    (doall (map #(Integer/parseInt %) (line-seq (java.io.BufferedReader. *in*))))) ;; reading lines from stdin

(def arr (getInput))
(println (reduce 1 (distinct (for 
        [a arr
         b arr
         :let [sum (+ a b)]
         :when (= sum 2020)] (* a b)))))