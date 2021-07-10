package main

import (
	"fmt"
	"net/http"

	"gopl.io/ch7/eval"
)

func parseAndCheck(s string) (eval.Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}
	expr, err := eval.Parse(s)
	if err != nil {
		return nil, err
	}

	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}

	for v := range vars {
		if v != "x" && v != "y" && v != "r" {
			return nil, fmt.Errorf("undefined variable: %s", v)
		}
	}
	return expr, nil
}

func plot(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	expr, err := parseAndCheck(r.Form.Get("expr"))
	if err != nil{
		http.Error(w, "bad expr"+ err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	suface
}

func  mian(){
	
}