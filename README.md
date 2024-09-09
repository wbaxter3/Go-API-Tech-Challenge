![CapTech Banner](resources/images/CaptechLogo.png)

# python-api-tech-challenge: Introduction

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Prerequisites](#prerequisites)
- [Development Environment Setup](#development-environment-setup)
- [Assignment Work & Submission](#assignment-work--submission)
- [Next Part](#if-you-made-it-this-far-congratulations-you-are-now-ready-to-proceed-with-the-tech-challenge)


## Overview

Welcome to the Python API Tech Challenge, presented by CapTech Consulting! The Challenge consists of 2 parts. This guide for Part 1 of the Challenge provides instructions for setting up your development environment in order to begin the main part of the challenge in Part 2. **Be sure to read the README for each part before beginning!**

## Prerequisites

This Tech Challenge introduces building web APIs with Python via a common use case: the creation of a REST API server. The REST API server will expose endpoints that allow users to perform Create, Read, Update, and Delete (CRUD) operations of information related to Courses, Professors, and Students of a fictional educational institution.

This tech challenge can be completed in multiple different ways including using the standard library or using a 3rd party framework such as Flask or FastAPI. The choice of how you complete this challenge is up to you as long as you complete all of the requirements laid out in Step 2!

To complete the challenge, the following are **REQUIRED**:

- A basic understanding of programming with Python
- A local development environment with the following software installed:
  - **Python** version 3.12+
  - **Git** (version control system)
  - **VSCode or your favorite Integrated Development Environment (IDE)**
  - **DB Browser for SQLite**
  - **Access to the Python API Tech Challenge code repository on GitHub**
  
## Development Environment Setup

### Install Python

To install Python, see the official [Python
documentation](https://wiki.python.org/moin/BeginnersGuide/Download)
for your operating system. Please note that this Tech Challenge was developed using Python 3.12. Other versions may work but we cannot guarantee that.

### Install Git

To install Git, see the installation instructions from [Git's
documentation](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git) for your operating system.
s
### IDEs

You may use your preferred IDE to complete this challenge. We will be using VSCode which can be installed [here](https://code.visualstudio.com/download).

If you are using VSCode, we also recommend installing the following extensions:
- [Python](https://marketplace.visualstudio.com/items?itemName=ms-python.python) (Python language support for VSCode)
- [Pylance](https://marketplace.visualstudio.com/items?itemName=ms-python.vscode-pylance) (Python linter and language support)
- [IntelliCode](https://marketplace.visualstudio.com/items?itemName=VisualStudioExptTeam.vscodeintellicode) (code autocomplete)
- [Black](https://marketplace.visualstudio.com/items?itemName=ms-python.black-formatter) (Python opinionated formatter)
- [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) (REST client for testing endpoints)

Lastly, if you need help getting used to writing Python in VSCode, see this help article [here](https://code.visualstudio.com/docs/languages/python).

### Install SQLite and DB Browser for SQLite

This tech challenge uses SQLite as a practice database. Installation instructions can be found [here](https://www.tutorialspoint.com/sqlite/sqlite_installation.htm).

To make viewing this database easier, we recommend using DB Browser for SQLite which can be installed [here](https://sqlitebrowser.org/dl/).

### Access the Tech Challenge

In order to obtain the source code and instructions for each part of the Tech Challenge, you will need access to the Challenge's code repository on GitHub. See the following instructions to request access and clone the repo:

#### Create an Account and Setup Authentication

1. **Create a new Github account** with a new username and password [here](https://github.com/join?source=header-home).
1. **Link your new GitHub account to CapTech** by clicking [here](http://capte.ch/github) after logging into GitHub with your new account. Once there you should see a pop-up on the top of the page with a button to authenticate your account. You can follow the steps to authenticate.
1. In the future, if your browser logs you in automatically, you may not see the CapTech tech challenges. You will need to click on the single sign-on link and authenticate through CapTech to see the repositories. This link is located on the top of your home screen.

## Assignment Work & Submission

[GitHub classroom](https://classroom.github.com/a/xMxWfNub) should already have generated a copy of the tech challenge for you to work on named `captechconsulting/python-api-tech-challenge-{your GitHub username}`. You will be working from this repository for the duration of the challenge.
> [!WARNING]  
> **GitHub classroom is not yet configured for this project. To proceed, please clone this repo and then push it to your own GitHub account.**

### Working branch

Once you have cloned your working repository to your local machine, create a branch to house work for each part.

1. Navigate to the cloned repository on your local machine
1. Check out the relevant branch with the starter code for the part you are working on. Using `api_tech_challenge` as an example:

    ```bash
    git checkout api_tech_challenge
    ```

1. Create a new branch for your development work. For example:

    ```bash
    git checkout -b develop/api_tech_challenge
    ```

1. Push your new local branch to the remote repository and set your local branch to track the new remote branch:

    ```bash
    git push -u origin develop/api_tech_challenge
    ```

After you have completed your work for each part and committed & pushed your changes up to the remote repository, you
are ready to open a pull request for submission.

### Submission

1. Navigate to your GitHub repository
1. Open a pull request from your working branch to the starting branch (for example: `develop/api_tech_challenge`&rarr;`api_tech_challenge`).
1. **Be sure to assign a reviewer so they will receive a notification that your solution is ready for review.**

  > If you do not have a reviewer yet, or need assistance, please make a post in the Python CoP channel [Tech Challenge Collaboration](https://teams.microsoft.com/l/channel/19%3Ab8ed52b50ec44d4e9f03ca4e60a5df43%40thread.tacv2/Tech%20Challenge%20Collaboration?groupId=713ca228-a89f-4f1f-99ea-e684d961144b&tenantId=).


## If you made it this far, Congratulations! You are now ready to proceed with the Tech Challenge!

[Go to Python API Tech Challenge: Assignment](../../tree/api_tech_challenge)
