package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/dhanusaputra/go-exercises/twitter/twitter"
)

func main() {
	var (
		keyFile    string
		userFile   string
		tweetID    string
		numWinners int
	)
	flag.StringVar(&keyFile, "key", ".keys.json", "consumer key and secret for twitter API")
	flag.StringVar(&userFile, "users", "users.csv", "users who have retweeted")
	flag.StringVar(&tweetID, "tweet", "1241994872187318274", "ID of tweet wish to find")
	flag.IntVar(&numWinners, "winner", 0, "number of winner")
	flag.Parse()

	key, secret, err := keys(userFile)
	if err != nil {
		panic(err)
	}
	client, err := twitter.New(key, secret)
	if err != nil {
		panic(err)
	}
	newUsernames, err := client.Retweeters(tweetID)
	if err != nil {
		panic(err)
	}
	existUsernames := existing(userFile)
	allUsernames := merge(newUsernames, existUsernames)
	err = writeUsers(userFile, allUsernames)
	if err != nil {
		panic(err)
	}
	if numWinners == 0 {
		return
	}
	existUsernames = existing(userFile)
	winners := pickWinners(existUsernames, numWinners)
	fmt.Println("The winners are :")
	for _, username := range winners {
		fmt.Printf("\t%s\n", username)
	}
}

func keys(keyFile string) (key, secret string, err error) {
	var keys struct {
		Key    string `json:"consumer_key"`
		Secret string `json:"consumer_secret"`
	}
	f, err := os.Open(".keys.json")
	if err != nil {
		return "", "", err
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	dec.Decode(&keys)
	return keys.Key, keys.Secret, nil
}

func existing(usersFile string) []string {
	f, err := os.Open(usersFile)
	if err != nil {
		return []string{}
	}
	defer f.Close()
	r := csv.NewReader(f)
	lines, err := r.ReadAll()
	users := make([]string, 0, len(lines))
	for _, line := range lines {
		users = append(users, line[0])
	}
	return users
}

func merge(a, b []string) []string {
	uniq := make(map[string]struct{}, len(a)+len(b))
	for _, user := range a {
		uniq[user] = struct{}{}
	}
	for _, user := range b {
		uniq[user] = struct{}{}
	}
	ret := make([]string, 0, len(uniq))
	for user := range uniq {
		ret = append(ret, user)
	}
	return ret
}

func writeUsers(userFile string, users []string) error {
	f, err := os.OpenFile(userFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	w := csv.NewWriter(f)
	for _, username := range users {
		if err = w.Write([]string{username}); err != nil {
			return err
		}
	}
	w.Flush()
	return w.Error()
}

func pickWinners(users []string, numWinners int) []string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	perm := r.Perm(len(users))
	winners := perm[:numWinners]
	ret := make([]string, 0, numWinners)
	for _, idx := range winners {
		ret = append(ret, users[idx])
	}
	return ret
}
