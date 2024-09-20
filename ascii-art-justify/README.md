## Ascii-art-justify

![Go-Logo_Blue-removebg-preview(2)](https://github.com/makebelief/makebelief/assets/166484145/ad53f422-f338-4dd7-9ef1-ab772aa1fbb5)
### Overview

This project is an ASCII art program implemented in Go. It takes a string as input and outputs a graphic representation of the string using ASCII characters. The output resembles the input string but with ASCII characters. Additionally the output is printed to the terminal with a specified alignment
### Content

-[Features](#features)

-[Banner Formats](#banner-formats)

-[How to Install](#how-to-install)

-[Usage](#usage)

-[Example Outputs](#example-outputs)

-[Allowed Packages](#allowed-packages)

-[How to contribute](#how-to-contribute--👷)

### Features

- Handles input with numbers, letters, spaces and newlines.
- Supports various graphical templates using ASCII characters.
- prints art with specified alignment
- Written in Go with adherence to good practices.
- Includes test files for unit testing.￼ 


### Banner Formats

The ASCII banner files included in the project have specific graphical representations formatted in ASCII. Each character has a height of 8 lines, and characters are separated by a newline (\n) character.

The following banner formats are included:
- `shadow`
- `standard`
- `thinkertoy`

### How to Install
Ascii-art requires [go](https://go.dev/dl/)  v 1.22.2 to run

Once go is installed  clone into this repository to do this 

- Open terminal and run
``` sh
git clone https://github.com/skanenje/ascii-art-justify
cd ascii-art-justify
```

### Usage
Once you are in the directory,

To run the program, use the following command, on the terminal command line:


```bash
Usage: go run . [OPTION] [STRING]

EX: go run . --align=<position> <substring to be colored> "something"
```

For example:

```bash
go run . "--align=right" "Hello" standard
```
To print ascii art in left allignment
```bash
go run . "--align=left" "<your string>" thinkertoy
```
To print ascii art in justify allignment
```bash
go run . "--align=justify" "<substring to be colored>" "<your string>" shadow
```

### Allowed Packages

Only standard Go packages are used in this project.





## HOW TO CONTRIBUTE ? 👷 

**1.** Fork [this](https://github.com/skanenje/ascii-art-justify) repository.

**2.** Clone the forked repository.

```terminal
git clone https://github.com/skanenje/ascii-art-justify
```

**3.** Navigate to the project directory.

```terminal
cd ascii-art-justify
```

**4.**  MAKE A NEW FOLDER WITH YOUR PROJECT NAME INSIDE 
<br>

**5.**  Also Add a README file in your project folder which consists of Description/screenshots about your project !
          
 
<br>

**6.** Create a new branch.

```terminal
git checkout -b <your_branch_name>
```

**7.** Add & Commit your changes.

```terminal
  git add .
  git commit -m "<your_commit_message>"
```

**7.** Push your local branch to the remote repository.

```terminal
git push -u origin <your_branch_name>
```

**8.** Create a Pull Request!

**Congratulations!** Sit and relax till we review your PR, you've made your contribution to (https://github.com/skanenje/ascii-art-color) project

### Issues and Contributions

If you encounter any issues or have suggestions for improvements, feel free to submit an issue or propose a change.

## Enjoy ASCII Art Web in Go!

Feel free to explore the codebase, run the program with different inputs, and contribute to enhancing the project. Happy coding experience!
### Contributors

<table>
<tr>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://www.linkedin.com/in/swabri-musa-565350291?lipi=urn%3Ali%3Apage%3Ad_flagship3_profile_view_base_contact_details%3Buf0Ls4oWR2O2WLUMO5sIBg%3D%3D>
            <img src=https://learn.zone01kisumu.ke/git/avatars/bc7899a0aac2630a0a9b50bf330437a7?size=870 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=Swabri/>
            <br />
            <sub style="font-size:14px"><b>skanenje</b></sub>
        </a>
    </td>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://www.linkedin.com/in/rodgers-kaunda>
            <img src=https://learn.zone01kisumu.ke/git/avatars/aa19095145ab1ad43695e3cd3f7f3a5b?size=870 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=Rodgers/>
            <br />
            <sub style="font-size:14px"><b>krodgers</b></sub>
        </a>
    </td>
</tr>
</table>
