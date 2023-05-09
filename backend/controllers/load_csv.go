package controllers

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"

	model "github.com/kshitijjagtap/quiz_usingreact/models"
)

func answer_abs(fileName string) ([]model.Answer, error) {
	if fObj, err := os.Open(fileName); err == nil {
		csvR := csv.NewReader(fObj)
		if cLines, err := csvR.ReadAll(); err == nil {
			return parseProblem(cLines), nil
		} else {
			return nil, fmt.Errorf("error in reading csv")
		}
	} else {
		return nil, fmt.Errorf("error in the opening ")
	}
}

func Answer_puller() []model.Answer {
	fName := flag.String("f", "quiz.csv", "path of csv file")
	flag.Parse()
	answ, err := answer_abs(*fName)
	if err != nil {
		fmt.Printf("error opening the csv file")
	}
	return answ
}

func parseProblem(lines [][]string) []model.Answer {
	r := make([]model.Answer, len(lines))
	for i := 0; i < len(lines); i++ {
		r[i] = model.Answer{Id: lines[i][0], Ans: lines[i][1]}
	}
	return r
}
