# Publishing Package to Go

To publish a package to Go, you need to create a new tag in your Git repository. Follow these steps:

1. Open your terminal and navigate to the root directory of your Go module.
2. Create a new tag using the following command, replacing `{tag_version}` with your desired version number (e.g., v1.0.0):
    ```bash
    git tag {tag_version}
    ```
3. Push the tag to your remote repository with the command:
    ```bash
    git push origin {tag_version}
    ```
4. Push to Go by running:
    ```bash
    GOPROXY=proxy.golang.org go list -m github.com/ivanleekk/ltadatamall@{tag_version}
    ```
