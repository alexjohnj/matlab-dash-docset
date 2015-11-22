# MATLAB Documentation for Dash

This is a very simple configuration file for [Dashing][dashing-cli]
with which you generate MATLAB documentation for
[Dash][dash-app]. Since Mathworks doesn't like people redistributing
their documentation, you need to build the documentation
yourself. Here's how you do that:

1. Install [Dashing][dashing-cli].
2. Clone this repository.
3. Copy the `dashing.json` file to the root of your MATLAB
   installation's help directory. On a Mac, this is
   `/Applications/MATLAB_R2015b.app/help`.
4. Run `dashing build` in the MATLAB help folder.
5. Add the newly generated `matlab.docset` file to Dash.

If you're using a version of MATLAB other than r2015b, you'll need to
edit the `dashing.json` file and replace all instances of `r2015b`
with whatever your version number is.

[dashing-cli]: https://github.com/technosophos/dashing
[dash-app]: https://kapeli.com/dash

## Completeness

Currently, the following types of pages are indexed in Dash:

- Functions
- Objects
- Classes
- Toolboxes (Libraries in Dash)
- Methods
- Properties
- Input Arguments (Parameters in Dash)
- Output Arguments (Values in Dash)
- Guides

In terms of table of contents support, some functions get a table of
contents with their input arguments and output arguments as well as
links to different sections of the documentation. Some functions don't
get a TOC because Mathworks is anything but consistent with the
structure of its documentation.

Graphics objects get a table of contents with all of their properties
in.

## Things To Do

There's a lot to do to improve the quality of the conversion. Here's
some of the things I want to do:

- Write a post-processing script (see below).
- Improve the reliability of TOC generation for functions.
- Generate a TOC for classes.
- Generate a TOC for objects.
- Generate a TOC for methods.
- Get an "input arguments" and "output arguments" entry type added to
  Dash.

## On a Post-Processing Script

A post-processing script would go a long way in improving the quality
of the outputted documentation. While it doesn't exist yet, here's
what a hypothetical post-processing script might do:

- Remove the header and sidebar from the documentation.
- Ensure that all examples are expanded/collapsed by default (your
  choice).
- Generate TOC links for input and output arguments. This is currently
  handled by Dashing but results in the input and output arguments
  being search-able. Considering these tend to have names like `X` and
  `Y`, these results just clutter the search results. A
  post-processing script should be able to add the TOC without
  generating search results.
 - Generate section links in the TOC for guides.
