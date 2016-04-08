# MATLAB Documentation for Dash

Mathworks doesn't like people redistributing their documentation, so the
wonderful [Dash][dash-app] doesn't include any MATLAB documentation out of the
box. This repository contains a configuration file for [Dashing][dashing-cli]
and a custom post-processing script (well, Go program) to generate a MATLAB
docset for Dash using your local copy of the documentation. No copyright
infringement required!

Here's what you'll need to do:

1. Install [Go][golang], [Dashing][dashing-cli] and [goquery][goquery-github].
2. Clone this repository.
3. Copy `dashing.json` and `post-processing.go` to the root of your MATLAB
   installation's help directory. On a Mac, this is at
   `/Applications/MATLAB_R2016a.app/help`.
4. Run `dashing build` from the root of the MATLAB help folder.
5. Run `go post-processing.go` from the root of the MATLAB help folder.
6. Add the newly generated `matlab.docset` file to Dash.

If you're using a version of MATLAB other than r2016a, you'll need to edit the
`dashing.json` file and replace all instances of `r2016a` with whatever your
version number is.

[dashing-cli]: https://github.com/technosophos/dashing
[dash-app]: https://kapeli.com/dash
[golang]: http://golang.org
[goquery-github]: http://github.com/PuerkitoBio/goquery

## What Gets Indexed

When you run `dashing build` the following parts of the documentation are
indexed:

- Functions
- Objects
- Classes
- Toolboxes (Libraries in Dash)
- Methods
- Properties
- Guides

When you run `post-processing.go`, it'll try and add a table of contents for

- Sections
- Input arguments
- Output variables

Sections are generated from `h2` elements in the documentation, and work pretty
reliably. Inputs and outputs are less reliable because they aren't represented
consistently throughout the documentation and so are harder to detect.

Note that Dash doesn't have a table of contents type for inputs/outputs to/from
functions so I've put them under the parameters/values types respectively. The
former makes sense, the latter, not so much.

## The Post-Processing "Script"

As already mentioned, `post-processing.go` is responsible for building the
table of contents in Dash for (some of) the documentation. In addition, the
script cleans up the documentation a bit. Notably, it removes the sidebar
(which is redundant with the TOC), the header (redundant and obscures anchor
links) and the general links at the bottom of each page.

## Things To Do

There's plenty of things to do to improve the generated documentation. The HTML
structure of the documentation is anything but consistent. As a result, there's
lots of cases where functions don't get a TOC or objects' methods aren't
detected. MATLAB has tonnes of documentation, so there's no way I'm going to
find all of these. If you run into anything that doesn't get indexed but seems
like it should, just file an issue.

There's also the issue of toolboxes. I have a couple of toolboxes and, AFAIK,
their functions, objects, guides etc. are indexed. Unfortunately, given the
aforementioned inconsistency of the documentation's HTML structure, I'm not
sure other toolboxes will be indexed. Again, if you come across any problems,
file an issue.
