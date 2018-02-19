package xcb

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ParseDisplay parses the display string and returns hostname, display number and screen number
func ParseDisplay(name string) (hostname string, display int, screen int, err error) {
	if len(name) == 0 {
		name = os.Getenv("DISPLAY")
	}
	if len(name) == 0 {
		return "", 0, 0, errors.New("no display name provided either in parameter or DISPLAY environmental variable")
	}

	splitName := strings.Split(name, "/")
	name = splitName[len(splitName)-1]

	splitOnColon := strings.Split(name, ":")
	if len(splitOnColon) < 2 {
		return "", 0, 0, errors.New("unknown format")
	}

	hostname = strings.Join(splitOnColon[0:len(splitOnColon)-1], ":")
	displayString := splitOnColon[len(splitOnColon)-1]

	if len(displayString) == 0 {
		return hostname, 0, 0, fmt.Errorf("display number missing")
	}

	splitDisplayString := strings.Split(displayString, ".")

	switch len(splitDisplayString) {
	case 1:
		display, err = strconv.Atoi(splitDisplayString[0])
		if err != nil {
			return "", 0, 0, fmt.Errorf("cannot convert display number: %s", err)
		}
		return hostname, display, 0, nil
	case 2:
		display, err = strconv.Atoi(splitDisplayString[0])
		if err != nil {
			return "", 0, 0, fmt.Errorf("cannot convert display number: %s", err)
		}
		screen, err = strconv.Atoi(splitDisplayString[1])
		if err != nil {
			return "", 0, 0, fmt.Errorf("cannot convert screen number: %s", err)
		}
		return hostname, display, screen, nil
	default:
		return "", 0, 0, fmt.Errorf("cannot parse display ID %q", splitDisplayString)
	}
}
