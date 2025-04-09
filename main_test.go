package main

import (
	"testing"
)

func TestScoreCalculation(t *testing.T) {
	commits := [][]string{
		{"timestamp", "username", "repository", "files", "additions", "deletions"},
		{"1615484900", "user1", "repo1", "2", "98", "50"},
		{"1615484901", "user2", "repo1", "1", "184", "50"},
		{"1744203732", "user1", "repo2", "1", "26", "5"},
	}
	expectedScores := []float64{196, 21}
	expectedOrder := []string{"repo1", "repo2"}
	repos := scoreRepos(commits)

	for i := range expectedOrder {
		//Expected Order
		if repos[i].Name != expectedOrder[i] {
			t.Errorf("Unexpected repo order: got %s expected %s", repos[i].Name, expectedOrder[i])
		}

		//Expected Score
		if expectedScores[i] != repos[i].Score {
			t.Errorf("Unexpected %s score: got %v expected %v", repos[i].Name, repos[i].Score, expectedScores[i])
		}
	}

}
