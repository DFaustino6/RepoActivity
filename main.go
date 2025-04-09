package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

type RepoActivity struct {
	Name           string
	UniqueUsers    map[string]bool
	Commits        int
	Additions      int
	Deletions      int
	RecentActivity int
	Score          float64
}

const (
	Points_Commits        = 1.0
	Points_Users          = 1.5
	Points_Additions      = 0.5
	Points_Deletions      = 0.5
	Points_RecentActivity = 3.0
)

func readCsvFile(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func scoreRepos(commits [][]string) []*RepoActivity {
	activity := make(map[string]*RepoActivity, 0)

	for i, row := range commits {
		if i == 0 {
			continue
		}

		timestamp, _ := strconv.Atoi(row[0])
		username := row[1]
		repository := row[2]
		additions, _ := strconv.Atoi(row[4])
		deletions, _ := strconv.Atoi(row[5])

		repo, exists := activity[repository]

		if !exists {
			repo = &RepoActivity{
				Name:        repository,
				UniqueUsers: make(map[string]bool),
			}

			activity[repository] = repo
		}

		repo.Commits++
		repo.Additions += additions
		repo.Deletions += deletions
		repo.UniqueUsers[username] = true

		now := int(time.Now().Unix())
		sevenDaysAgo := now - (7 * 24 * 60 * 60)

		if timestamp >= sevenDaysAgo {
			repo.RecentActivity++
		}

	}

	reposArray := make([]*RepoActivity, 0, len(activity))

	//Score the repos activity
	for _, repo := range activity {
		repo.Score = Points_Commits*float64(repo.Commits) +
			Points_Users*float64(len(repo.UniqueUsers)) +
			Points_Additions*float64(repo.Additions) +
			Points_Deletions*float64(repo.Deletions) +
			Points_RecentActivity*float64(repo.RecentActivity)

		reposArray = append(reposArray, repo)
	}

	//Sort repos by score desc
	sort.Slice(reposArray, func(i, j int) bool {
		return reposArray[i].Score > reposArray[j].Score
	})

	return reposArray
}

func main() {
	data := readCsvFile("commits.csv")

	repos := scoreRepos(data)

	limit := min(len(repos), 10)

	//Print the 10 most active repos
	for i, repo := range repos[:limit] {
		fmt.Printf("%d. %s %v\n", i+1, repo.Name, repo.Score)
	}
}
