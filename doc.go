// Ethiocal provides utilities for converting dates between the Ethiopian and
// Gregorian calendars and fetching important Ethiopian religious dates.
//
// This module offers two importable packages:
//
//   - [github.com/yinebebt/ethiocal/bahirehasab] — Ethiopian fasting and festival dates for a given year.
//   - [github.com/yinebebt/ethiocal/dateconverter] — conversion between Gregorian and Ethiopian dates.
//
// # GUI (default)
//
// Run without arguments to launch the desktop application.
//
//	ethiocal
//
// # CLI
//
// Use subcommands for terminal and scripting workflows.
//
//	ethiocal bahir 2017
//	ethiocal convert gtoe 2025 2 2
//	ethiocal convert etog 2017 5 25
//
// # HTTP Server
//
// Start the REST API with the --server flag.
//
//	ethiocal --server
//
// To integrate into your own Go project, import the bahirehasab and
// dateconverter packages directly.
package main
