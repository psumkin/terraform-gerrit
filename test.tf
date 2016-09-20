provider "gerrit" {
    auth = "digest"
    url = "http://127.0.1.1:8080"
    user = "tbot"
    password = "tbot"
}

resource "gerrit_project" "project_0" {
    description = "Project 0 Description"
    name = "project_0"
}

resource "gerrit_project" "project_1" {
    description = "Project 1 Description"
    name = "Project 1"
}
