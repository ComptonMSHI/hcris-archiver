# Medicare Cost Report Data Archiver

This is an application that will download the HCRIS Medicare cost report data files for the 2552-96 and 2552-10 form data for hospitals.  This application does two things based on configuration in `config.yaml`:

- Downloads the data files in `zip` format, and stores them within a directory organized by the date of download.

- Extracts the `csv` data files into an output directory, also organized by the date of download.

This application is useful because it maintains a running archive to be able to pull files on a quarterly basis from CMS.  Optionally, this application can be configured to extract from an already stored archive date, and not pull directly from HCRIS.

## IMPORTANT

Please be a good citizen. 

1. It is highly discouraged to repeatedly run this application and download the full data set from HCRIS. You can set a date to pull from your existing archive in the config.yaml file, particularly if you are testing and developing.

2. It is also unnecessary to run the archiver on an interval shorter than 90 days.  The release dates are available from CMS, and you might lag beyond that day since data release could be delayed for technical reasons at times.

A lot of good can come from this data, and we need to make sure it remains available in a manner that allows us to archive it for research.

## How to Run

Required files:

- A platform specific binary from the `dist` directory.
- A configured `config.yaml` 

These two files must be placed in a directory together.

The config file is completely documented, and distributions are available that should run on Windows, macOS and Linux.

## Development

The application is written in Go.  You will need to install this on your computer to make changes to the application beyond the configuration file.  It uses a package called [hcris-tools](https://github.com/ComptonMSHI/hcris-tools) that I also maintain, and holds much of the code that does the work.

If you wish to develop this application (or package) further, or make customizations, the source is provided along with script to build the binary for the supported operating systems.  If you make improvements, please contribute via pull request.