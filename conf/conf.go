package conf

const FileServerAddress = "http://10.0.13.134:3927"
const TestResultsDirFmt = `/opt/gitlab-data/gitlab-test/{{.Project}}/test-result/{{.TestStages}}/{{.Branch}}/{{.FileType}}`
