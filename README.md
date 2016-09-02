# terraform-gerrit
gerrit plugin for terraform


## Prepare Gerrit

* Start container

    ```
    docker run --name gerrit_1 -d -p 127.0.1.1:8080:8080 -p 29418:29418 -e  WEBURL=http://127.0.1.1:8080 openfrontier/gerrit

    docker ps -a
    ```
* Gerrit WebUI: SignIn with Launchpad account - got admin rights

    Then set username and add ssh key. Optional set ~/.ssh/config

    ```
    Host                127.0.1.1
        User            gerrituser
        IdentityFile    ~/.ssh/gerrituser.rsa
    ```

* SSH: Create bot account in the 'Non-Interactive Users' group

    ```
    ssh -v -p 29418 127.0.1.1 gerrit create-account --group \'Non-Interactive Users\' --http-password tbot tbot

    ## curl --digest --user username:password http://localhost:8080/a/path/to/api/
    curl -v --digest -u tbot:tbot http://127.0.1.1:8080/a/accounts/self
    ```

* SSH: Create bot group and include bot account

    ```
    ssh -v -p 29418 127.0.1.1 gerrit create-group --member tbot tbot
    ```

* Gerrit WebUI: Check members in bot group

* Gerrit WebUI: Add global capabilities to bot group for Create/Delete project
    in All-Projects access form

    http://127.0.1.1:8080/#/admin/projects/All-Projects,access


## Compile the terraform plugin and run

    ```
    go build -o terraform-provider-gerrit
    terraform plan
    TF_LOG=DEBUG terraform apply
    ```

## Useful resources

 * https://gerrit-review.googlesource.com/Documentation/rest-api.html
 * https://gerrit-review.googlesource.com/Documentation/cmd-index.html
