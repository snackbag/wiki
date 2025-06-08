# wiki

This project is meant as a global wiki for all of SnackBag's projects.\
https://wiki.snackbag.net/

## Building

To build, run GNU make: "`make`". It will produce a binary `wiki`.

You cam also build manually: `go mod tidy` and `go build`. 

For deploying to the server, MAKE SURE THAT `dev_mode` IN THE config.json` IS SET TO FALSE! Otherwise, the program will not cache pages but reload them on each request.

## Specifications

This wiki runs a [Compass](https://github.com/snackbag/compass) server in the backend and should be built with Go
1.22.7. All pages can be found within the `pages` folder -- contributions are welcome!

## Contributing

> How do you contribute?

This question can easily be solved! Well, kinda. Just read the text below, and you'll hopefully have all the information
you need. If not, just open an issue or join our [Discord server](https://discord.gg/7uYhxN7cFj), and we'll be happy to
help.

Rule of thumb: after you've forked the repo, create a new branch for your changes

### Modify wiki content

You can find all pages inside the `pages` folder. From there, you can select the file you need and just start writing.
All text should be written in Markdown, but there are some custom elements. If you want to add a note, type `NOTE BEGIN`
and below type the content of the note. You can end the note with `NOTE END`. To view an example of this behaviour,
check out the [TT20 setup page](/pages/tt20/setup.md)

### New wiki content

You may not just add new projects. If you want to create a new wiki for an entirely new project that doesn't have its
own wiki here yet, open a ticket.

Otherwise, you can go to each project and create new markdown files. The file should only contain alphabetical
characters, numbers and dashes. The page name will be displayed as each word being capitalized and all dashes being
replaced as spaces. You cannot create nested folders (e.g. `/tt20/extra/hello-world.md`). After your page is finished,
open the `pages/struct.json` file and add your newly created file to the `page_structure` value. Test your changes by
running `go run .` (after you've done `go mod tidy` at least once)

### Code

make it and pr it -- stick to Go conventions and use the minimum of extra libraries as possible. If you're changing the
layout: try to avoid as much JavaScript as possible. This wiki should be accessible to anybody, even those who have
JavaScript disabled on their browsers. Vanilla CSS & HTML for the win.
