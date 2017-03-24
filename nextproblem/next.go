package nextproblem

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)
const ProjectEulerProblemSite = "https://projecteuler.net/problem=%d"

func downloadProblemPage(n int) string {
	resp, err := http.Get(fmt.Sprintf(ProjectEulerProblemSite, n))
	if err != nil {
		return err.Error()
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func extractProblemText(problemPage string) string {

	problemPage = strings.Replace(problemPage, "\n", "\\n", -1)
	ret := ""
	startTitle := `<h2>`
	endTitle := `</h2>`
	startProb := `<div class="problem_content" role="problem">`
	endProb := `</div><br />`
	r, _ := regexp.Compile(fmt.Sprintf(".*%s(.*)%s.*%s(.*)%s.*", startTitle, endTitle, startProb, endProb))
	match := r.FindStringSubmatch(problemPage)

	for i, m := range match {
		if i != 0 {
			m = strings.Replace(m, "\\n", "\n", -1)
			ret += fmt.Sprintf("%s\n", m)
		}
	}
	return ret
}

func Run() {
	probNum := 5
	fmt.Printf("%s\n\n", fmt.Sprintf(ProjectEulerProblemSite, probNum))
	fmt.Printf("%s", extractProblemText(downloadProblemPage(probNum)))
}
