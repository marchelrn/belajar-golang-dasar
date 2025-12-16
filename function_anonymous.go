package main

func main() {
	blackListed := isBlacklisted
	// anonymous function
	blackListName := func(name string) bool {
		nameBlacklisted := []string{"budi", "joko", "marchel", "ani"}
		for _, blacklistedName := range nameBlacklisted {
			if name == blacklistedName {
				return true
			}
		}
		return false
	}
	name := "joko"
	println(blackListed(name, blackListName))
}

type Blacklist func(string) bool

func isBlacklisted(name string, blacklist Blacklist) string {
	if blacklist(name) == true {
		return "You are blacklisted"
	}
	return "Welcome, " + name
}
