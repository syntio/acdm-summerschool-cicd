# summerschool-cicd
A repository with a simple go program and a ci pipeline.

## Rundown of the repo

### main
Here you will find a simple GO program, with a functions package.

### .github/workflows
In this directory you place all your workflow files. In this example you will find a ci.yaml which contains a workflow which triggers when a pull request is made with a head branch being 'master'.  
Jobs that are in this example are (jobs within the workflow are running concurrently, while steps within a job run sequentially):

##### Unit tests
First step in this job is to set up golang. This step uses a premade action making it easier for us to install Go.  
Second step is to clone our repository to the GitHub Actions VM where our job is running.  
Third step is to position ourselves in the directory which contains our unit tests (inside the functions directory) and to run them.

##### Linter step
First two steps are the same as in the previous job. Again we want to install Go and to clone the repository.  
Last step in this jobs is to run the golangci-lint. There is a premade Github action which is used in this example.

##### Editorconfig step
Here we won't be using Go since we are only checking the consistency of coding styles.  
Firstly, we clone the repository. Once we have our repository mounted, we run the editorconfig with the help of the multiline command.  
What editorconfig checks (which rules of code styling) are located in the root of the repository in a file called .editorconfig. For the documentations check [this link](https://editorconfig.org/).