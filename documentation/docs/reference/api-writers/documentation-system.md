# Documentation System

We use [MkDocs](https://www.mkdocs.org/) for our documentation, and host at [HackIllinois API Docs](https://docs.api.hackillinois.org).
MkDocs is a Python module generates a static website, and/or a live-reloading server, from our Markdown-based documentation. It relies on a configuration file called `mkdocs.yml`, that is located the `documentation`. This configuration file can be used to specify the theme as well as menu structure of the documentation.

## Commands
Here are the commands that can be used for interfacing with MkDocs.
However, we have abstracted the ones useful to us through Makefile targets, which handle installing dependencies, and building a
release documentation website.

* `mkdocs new [dir-name]` - Create a new project.
* `mkdocs serve` - Start the live-reloading docs server.
* `mkdocs build` - Build the documentation site.
* `mkdocs help` - Print this help message.

## Project Layout

```py
mkdocs.yml                  # The configuration file.
Makefile                    # Makefile to install all dependencies
License.md                  # What our project is licensed under

docs/
├── index.md                # The documentation homepage
│
├── getting-started/        # Directory containing the developer guide, and other resources to set up and run the API locally
│   └── developer-guide.md  
├── services/               # Directory containing documentation for each service
├── gateway/                # Documentation for the gateway and middleware
└── api-writers/            # Includes information useful to API writers / developers
...                         # Other markdown pages, images and other files
```
