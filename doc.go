// Ethiocal is the Ethiopian Calendar graphical application for desktop and
// mobile. It converts dates between the Ethiopian and Gregorian calendars and
// shows Bahire-Hasab fasting and festival dates.
//
// Run the binary, or launch the packaged app, to open the GUI:
//
//	ethiocal
//
// # CLI and HTTP server
//
// Terminal and REST access live in the companion command at ./cmd/ethiocal-cli,
// built separately so the GUI app carries no CLI dependencies.
//
//	ethiocal-cli bahir 2017
//	ethiocal-cli convert gtoe 2025 2 2
//	ethiocal-cli convert etog 2017 5 25
//	ethiocal-cli --server
//
// # Libraries
//
// Two packages are importable for use in your own Go project:
//
//   - [github.com/yinebebt/ethiocal/bahirehasab] — Ethiopian fasting and festival dates for a given year.
//   - [github.com/yinebebt/ethiocal/dateconverter] — conversion between Gregorian and Ethiopian dates.
package main
