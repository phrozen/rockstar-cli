package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"time"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/urfave/cli"
)

// Repo general information about repo
type Repo struct {
	DirPath  string
	FilePath string
}

func randomName() string {
	return fmt.Sprintf("%v-%v-%v", adjs[rand.Intn(len(adjs))], nouns[rand.Intn(len(nouns))], repos[rand.Intn(len(repos))])
}

func newRepo(filename string) Repo {
	dirPath := randomName()

	os.Mkdir(dirPath, 0755)
	os.Chdir(dirPath)
	exec.Command("git", "init", ".").Run()

	return Repo{DirPath: dirPath, FilePath: filename}
}

func (r *Repo) appendCommit(data string, date time.Time) {
	err := os.WriteFile(r.FilePath, []byte(data), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	os.Setenv("GIT_AUTHOR_DATE", date.Format(time.RFC3339))
	os.Setenv("GIT_COMMITTER_DATE", date.Format(time.RFC3339))

	exec.Command("git", "add", r.FilePath).Run()
	exec.Command("git", "commit", "-m", messages[rand.Intn(len(messages))]).Run()
}

func (r *Repo) createRemoteRepo() {
	_, err := exec.Command("hub", "version").Output()
	if err != nil {
		return
	}
	fmt.Println("Creating in your github account")
	exec.Command("hub", "create").Run()
	exec.Command("git", "push", "origin", "master").Run()
}

func main() {
	app := cli.NewApp()
	app.Name = "rockstar-cli"
	app.Usage = "It make you a rockstar in less than 10 seconds"
	app.Version = "0.0.1"
	app.Author = "Luis Ezcurdia"
	app.Email = "ing.ezcurdia@gmail.com"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "days, d",
			Value: "500",
			Usage: "days of activity",
		},
		cli.StringFlag{
			Name:  "code, c",
			Value: "writeln('Go is Awesome!!!')",
			Usage: "code base",
		},
		cli.StringFlag{
			Name:  "filename, f",
			Value: "main.go",
			Usage: "output file",
		},
		cli.BoolFlag{
			Name:  "keep, k",
			Usage: "do not create a new repo",
		},
	}
	app.Action = func(c *cli.Context) {
		days, _ := strconv.Atoi(c.String("days"))
		if days > 0 {
			days *= -1
		}

		repo := Repo{
			DirPath:  ".",
			FilePath: c.String("filename"),
		}
		if !c.Bool("keep") {
			repo = newRepo(c.String("filename"))
		}

		for i := days; i < 0; i++ {
			d := time.Now().Add(time.Duration(i*24) * time.Hour)
			if (d.Weekday() == time.Sunday || d.Weekday() == time.Saturday) && i%3 == 0 {
				continue
			}
			for j := 0; j < rand.Intn(10); j++ {
				authorDate := time.Date(d.Year(), d.Month(), d.Day(), int(rand.NormFloat64()*3.0+12.0), rand.Intn(59), rand.Intn(59), 0, d.Location())
				uid, err := uuid.NewV5(uuid.NamespaceURL, []byte(time.Now().Format(time.RFC3339Nano)))
				commitData := uid.String()
				if err != nil {
					continue
				}
				repo.appendCommit(commitData, authorDate)
			}
			fmt.Print(".")
		}
		repo.appendCommit(c.String("code"), time.Now())
		os.Setenv("GIT_AUTHOR_DATE", "")
		os.Setenv("GIT_COMMITTER_DATE", "")
		repo.createRemoteRepo()

		fmt.Printf("\nProyect created on: %v", repo.DirPath)
		fmt.Println("\nNow you are a goddamn rockstar!")
	}
	app.Run(os.Args)
}
