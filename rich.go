package rich

import (
	"github.com/frustra/bbcode"
	"github.com/muesli/termenv"
	"golang.org/x/image/colornames"
	"image/color"
	"strings"
)

type StyleTag struct {
	tag   string
	style termenv.Style
}

func NewStyleTag(tag string, base termenv.Style) StyleTag {
	return StyleTag{
		tag:   tag,
		style: tagToStyle(tag, base),
	}
}

func tagToStyle(tag string, base termenv.Style) termenv.Style {
	style := base
	profile := termenv.ColorProfile()
	setBg := false
	setColor := func(c termenv.Color) termenv.Style {
		if setBg {
			return style.Background(c)
		}
		return style.Foreground(c)
	}

	parts := strings.Split(tag, " ")
	for _, part := range parts {
		switch part {
		case "italic":
			style = style.Italic()
		case "bold":
			style = style.Bold()
		case "underline":
			style = style.Underline()
		case "overline":
			style = style.Overline()
		case "crossout":
			style = style.CrossOut()
		case "blink":
			style = style.Blink()
		case "faint":
			style = style.Faint()
		case "reverse":
			style = style.Reverse()
		case "on":
			setBg = true
		default:
			if namedColor := colorByName(part); namedColor != nil {
				style = setColor(profile.FromColor(namedColor))
			} else {
				style = setColor(profile.Color(part))
			}
		}
	}

	return style
}

func colorByName(name string) color.Color {
	namedColor, exists := colornames.Map[name]
	if exists {
		return namedColor
	}
	return nil
}

func Stylize(s string) string {

	styleStack := make([]StyleTag, 1)

	builder := strings.Builder{}

	format := func(s string) {
		style := styleStack[len(styleStack)-1]
		builder.WriteString(style.style.Styled(s))
	}

	setStyle := func(tag string) {
		styleStack = append(styleStack, NewStyleTag(tag, styleStack[len(styleStack)-1].style))
	}

	popStyle := func(tag string) {
		style := styleStack[len(styleStack)-1]
		if style.tag != tag {
			panic("Style tag doesn't match!")
		}
		styleStack = styleStack[:len(styleStack)-1]
	}

	for token := range bbcode.Lex(s) {
		switch token.ID {
		case bbcode.TEXT:
			format(token.Value.(string))
		case bbcode.OPENING_TAG:
			tag := token.Value.(bbcode.BBOpeningTag)
			setStyle(tag.Raw[1 : len(tag.Raw)-1])

		case bbcode.CLOSING_TAG:
			tag := token.Value.(bbcode.BBClosingTag)
			popStyle(tag.Raw[2 : len(tag.Raw)-1])
		}
	}

	return builder.String()
}
