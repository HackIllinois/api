
### Commands

* `mkdocs new [dir-name]` - Create a new project.
* `mkdocs serve` - Start the live-reloading docs server.
* `mkdocs build` - Build the documentation site.
* `mkdocs help` - Print this help message.

### Project Layout

    mkdocs.yml                  # The configuration file.
    Makefile                    # Makefile to install all dependencies
    License.md                  # What our project is licensed under

    docs/
        index.md                # The documentation homepage.

        getting-started/        # Directory containing the developer guide, and other resources to set up and run the API locally
            developer-guide.md  
        services/               # Directory containing documntation for each service.
        gateway/                # Documentation for the gateway and middleware
    ...                         # Other markdown pages, images and other files.