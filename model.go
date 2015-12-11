package main

import (
	"net/http"
)

type Connection struct {
	Username    string
	AccessToken string
	BaseUrl     string
}

type JenkinsApi struct {
	connection *Connection
	client     *http.Client
}

type Job struct {
	Id                string `json:"id"`
	Number            int `json:"number"`
	Result            string `json:"result"`
	Description       string `json:"description"`
	DisplayName       string `json:"displayName"`
	FullDisplayName   string `json:"fullDisplayName"`
	Duration          int64 `json:"duration"`
	EstimatedDuration int64 `json:"estimatedDuration"`
	QueueId           int `json:"queueId"`
	Timestamp         int64 `json:"timestamp"`
	Url               string `json:"url"`
	Building          bool `json:"building"`
	Artifacts         []Artifact `json:"artifacts"`
	Actions           []Action `json:"actions"`
}

type Artifact struct {
	DisplayPath  string `json:"displayPath"`
	FileName     string `json:"fileName"`
	RelativePath string `json:"relativePath"`
}

type Action struct {
	Parameters []Parameter `json:"parameters"`
	Causes     []Cause `json:"causes"`
	TestResult
}

type Cause struct {
	ShortDescription string `json:"shortDescription"`
	User
	UpstreamJob
}

type User struct {
	UserId           string `json:"userId"`
	UserName         string `json:"userName"`
}

type UpstreamJob struct {
	UpstreamBuild    int `json:"upstreamBuild"`
	UpstreamProject  string `json:"upstreamProject"`
	UpstreamUrl      string `json:"upstreamUrl"`
}

type TestResult struct {
	FailCount  int `json:"failCount"`
	SkipCount  int `json:"skipCount"`
	TotalCount int `json:"totalCount"`
	UrlName    string `json:"urlName"`
}

type Parameter struct {
	Name  string `json:"name"`
	Value interface{} `json:"value"`
}
