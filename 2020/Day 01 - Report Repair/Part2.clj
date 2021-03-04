(defn getInput []
    (with-open [rdr (clojure.java.io/reader "./in.txt")]
    (doall (map #(Integer/parseInt %) (line-seq rdr)))))

(def arr (getInput))
(doseq [a arr
        b arr
        c arr]
    (def sum (+ a b c))
    (when (= sum 2020)
        (println (* a b c))))