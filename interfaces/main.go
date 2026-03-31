package main

import "fmt"

// The interface — a contract with one method
type Checker interface {
	Check() string
}

// HTTPChecker satisfies Checker — it has a Check() method
type HTTPChecker struct{ URL string }

func (h HTTPChecker) Check() string {
	return "HTTP check on " + h.URL
}

// DBChecker also satisfies Checker
type DBChecker struct{ Host string }

func (d DBChecker) Check() string {
	return "DB check on " + d.Host
}

// This function works with ANY Checker — HTTP, DB, or anything else
func RunCheck(c Checker) {
	fmt.Println(c.Check())
}

func main() {
	RunCheck(HTTPChecker{URL: "http://api:8080/health"})
	RunCheck(DBChecker{Host: "postgres:5432"})
}
