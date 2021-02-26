(defn getInput []
    (with-open [rdr (clojure.java.io/reader "./in.txt")] ;;put input in in.txt file
    (doall (map #(Integer/parseInt %) (line-seq rdr)))))

(def arr (getInput))
(doseq [a arr
        b arr]
    (def sum (+ a b))
    (when (= sum 2020)
        (println (* a b))))