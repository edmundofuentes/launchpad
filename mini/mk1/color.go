package mk1

// Reference: https://sourceforge.net/p/ourorgan/discussion/962124/thread/a1d3cf0c/bb64/attachment/LaunchpadMiniColours.pdf

type Color struct {
	R int
	G int
}

const (
	Off = iota

	RedHigh
	RedMid
	RedLow

	GreenHigh
	GreenMid
	GreenLow

	YellowHigh
	YellowMid
	YellowLow

	YellowGreen

	OrangeRed
	OrangeRed2

	OrangeHigh
	OrangeHigh2
	OrangeMid
)

// return the correct Green and Red values to build a color
func ColorGR(color int) (g, r int){
	switch color {
	case Off:
		return 0, 0

	case RedHigh:
		return 0, 3
	case RedMid:
		return 0, 2
	case RedLow:
		return 0, 1

	case GreenHigh:
		return 3, 0
	case GreenMid:
		return 2, 0
	case GreenLow:
		return 1, 0

	case YellowHigh:
		return 3, 2
	case YellowMid:
		return 2, 1
	case YellowLow:
		return 1, 1

	case YellowGreen:
		return 3, 1

	case OrangeRed:
		return 1, 3
	case OrangeRed2:
		return 1, 2

	case OrangeHigh:
		return 3, 3
	case OrangeHigh2:
		return 2, 3
	case OrangeMid:
		return 2, 2


	default:
		// if the color is unknown, default to "off" state
		return 0, 0
	}

}


func (l *Launchpad) DebugColors() {
	l.LightColor(0, 0, RedHigh)
	l.LightColor(1, 0, RedMid)
	l.LightColor(2, 0, RedLow)

	l.LightColor(0, 1, GreenHigh)
	l.LightColor(1, 1, GreenMid)
	l.LightColor(2, 1, GreenLow)

	l.LightColor(0, 2, YellowHigh)
	l.LightColor(1, 2, YellowMid)
	l.LightColor(2, 2, YellowLow)

	l.LightColor(0, 3, YellowGreen)

	l.LightColor(0, 4, OrangeRed)
	l.LightColor(1, 4, OrangeRed2)

	l.LightColor(0, 5, OrangeHigh)
	l.LightColor(1, 5, OrangeHigh2)
	l.LightColor(2, 5, OrangeMid)
}