// Ethiocal provides utilities for converting dates between the Ethiopian and
// Gregorian calendars and fetching important Ethiopian religious dates.
//
// This module offers two importable packages:
//
//   - [bahirehasab] — Ethiopian fasting and festival dates for a given year.
//   - [dateconverter] — conversion between Gregorian and Ethiopian dates.
//
// The ethiocal binary can be used in three ways:
//
//  1. GUI (default):
//     Run without arguments to launch the desktop application.
//
//     ethiocal
//
//  2. CLI:
//     Use subcommands for terminal and scripting workflows.
//
//     ethiocal bahir 2017
//     ethiocal convert gtoe 2025 2 2
//     ethiocal convert etog 2017 5 25
//
//  3. HTTP Server:
//     Start the REST API with the --server flag.
//
//     ethiocal --server
//
// To integrate into your own Go project, import the bahirehasab and
// dateconverter packages directly.
package main
