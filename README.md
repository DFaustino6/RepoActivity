# Repository Activity Scoring

This challenge implements an algorithm that scores and ranks repositories based on their commit activity. The goal is to identify the **top 10 most active repositories** using a points scoring system derived from commits information present in a csv file.

## Scoring Algorithm

Each repository is scored using the following factors:

| Metric            | Points | Description |
|-------------------|--------|-------------|
| Commits           | 1.0    | Total number of commits |
| Unique Users      | 1.5    | Number of distinct commit authors |
| Additions         | 0.5    | Total lines of code added |
| Deletions         | 0.5    | Total lines of code deleted |
| Recent Activity   | 3.0    | Number of commits in the last 7 days |

**Final Score Formula:**
```
Score = (1.0 × Commits) + (1.5 × UniqueUsers) + (0.5 × Additions) + (0.5 × Deletions) + (3.0 × RecentActivity)
```
This formula takes into account overall activity volume, how recent the work is, meaning that it is still being maintained, and how many different people are contributing, rewarding collaboration within the project. The decision of ignoring the number of files changed was made to avoid giving an unfair advantage to larger repositories.

**Top 10 Repos:**

| Rank | Repository | Value        |
|------|------------|--------------|
| 1    | **repo476** | 1,819,695.0  |
| 2    | **repo260** | 571,376.0    |
| 3    | **repo920** | 328,319.5    |
| 4    | **repo795** | 282,388.5    |
| 5    | **repo161** | 206,756.5    |
| 6    | **repo1143**| 194,123.5    |
| 7    | **repo518** | 174,305.5    |
| 8    | **repo1185**| 151,079.0    |
| 9    | **repo1243**| 140,036.5    |
| 10   | **repo250** | 119,454.0    |


## Input Format

The input file `commits.csv` should have the following columns:

| Column     | Description                             |
|------------|-----------------------------------------|
| `timestamp`| Unix timestamp of the commit            |
| `username` | GitHub username of the commit author    |
| `repository` | Repository name the commit was pushed to |
| `files`    | Number of files changed (not used in scoring) |
| `additions`| Number of lines added                   |
| `deletions`| Number of lines deleted                 |


## Running the Program

### Run

This algorithm was created using Go 1.23.0. 
> **Note:** The algorithm is expecting a file "commits.csv" in the root directory

```bash
go run main.go
```

### Test
```bash
go test
```