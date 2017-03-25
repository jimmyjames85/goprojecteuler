package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/jaytaylor/html2text"
	"github.com/jimmyjames85/eflag"
)

const ProjectEulerProblemSite = "https://projecteuler.net/problem=%d"

type settings struct {
	ProblemNum int `flag:"p,problem" desc:"which problem to download"`
	//Dir        string `flag:"d,dir" desc:"which problem to download"`
}

func main() {
	s := &settings{
		ProblemNum: 1,
		//Dir:        wd,
	}
	eflag.StructVar(s)
	flag.Usage = eflag.POSIXStyle
	flag.Parse()

	problemURL := problemUrl(s.ProblemNum)
	title, problemText := extractProblemText(downloadProblemPage(s.ProblemNum))
	if title == "" || problemText == "" {
		errAndQuit(fmt.Errorf("unable to retrieve or parse: %s\n", problemURL))
	}

	wd, err := os.Getwd()
	mustBeNil(err)

	dirName := fmt.Sprintf("prob%d", s.ProblemNum)
	dirPath := filepath.Join(wd, dirName)
	os.Mkdir(dirPath, 0766)
	mustBeNil(os.Chdir(dirPath))

	fileName := fmt.Sprintf("%s.go", strings.ToLower(strings.Replace(title, " ", "", -1)))
	fileName = strings.Replace(fileName, ",", "_", -1)

	if _, err = os.Stat(fileName); !os.IsNotExist(err) {
		errAndQuit(fmt.Errorf("%s already exists\n", fileName))
	}

	f, err := os.Create(fileName)
	mustBeNil(err)

	_, err = f.Write([]byte(generateGoFileContent(problemURL, title, problemText)))
	mustBeNil(err)

}

func downloadProblemPage(n int) string {
	resp, err := http.Get(problemUrl(n))
	if err != nil {
		return err.Error()
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func extractProblemText(problemPage string) (string, string) {

	problemPage = strings.Replace(problemPage, "\n", "\\n", -1)
	startTitle := `<h2>`
	endTitle := `</h2>`
	startProb := `<div class="problem_content" role="problem">`
	endProb := `</div><br />`
	r, _ := regexp.Compile(fmt.Sprintf(".*%s(.*)%s.*%s(.*)%s.*", startTitle, endTitle, startProb, endProb))
	match := r.FindStringSubmatch(problemPage)

	var title, problemHTML string
	if len(match) == 3 {
		title = strings.Replace(match[1], "\\n", "\n", -1)
		problemHTML = strings.Replace(match[2], "\\n", "\n", -1)
	}

	problemText, err := html2text.FromString(problemHTML)
	mustBeNil(err)

	return title, problemText
}

func generateGoFileContent(problemURL, title, problemText string) string {
	hr := strings.Repeat("-", len(title))
	ret := fmt.Sprintf("package main\n\n// %s\n//\n// %s\n// %s\n//\n", problemURL, title, hr)
	lines := strings.Split(problemText, "\n")
	for _, line := range lines {
		ret += fmt.Sprintf("// %s\n", line)
	}
	ret += fmt.Sprintf("\nfunc main() {\n\n\tprintln(`TODO`)\n\n}\n")
	return ret
}

func mustBeNil(err error) {
	if err == nil {
		return
	}
	errAndQuit(err)
}

func errAndQuit(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	os.Exit(-1)
}

func problemUrl(n int) string {
	return fmt.Sprintf(ProjectEulerProblemSite, n)
}
