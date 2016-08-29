package house

import "strings"

func Embed(relPhrase, nounPhrase string) string {
	return relPhrase + " " + nounPhrase
}

func Verse(subject string, relPhrases []string, nounPhrase string) string {
	if len(relPhrases) == 0 {
		subject += Embed("", nounPhrase)
	} else {
		subject += " " + Embed(relPhrases[0], Verse("", relPhrases[1:], nounPhrase))
	}
	return strings.TrimSpace(subject)
}

var phrases = []string{
	"the horse and the hound and the horn\nthat belonged to",
	"the farmer sowing his corn\nthat kept",
	"the rooster that crowed in the morn\nthat woke",
	"the priest all shaven and shorn\nthat married",
	"the man all tattered and torn\nthat kissed",
	"the maiden all forlorn\nthat milked",
	"the cow with the crumpled horn\nthat tossed",
	"the dog\nthat worried",
	"the cat\nthat killed",
	"the rat\nthat ate",
	"the malt\nthat lay in",
	"the house that",
}

func Song() string {
	var s []string
	for i := len(phrases) - 1; i >= 0; i-- {
		s = append(s, Verse("This is", phrases[i:], "Jack built."))
	}
	return strings.Join(s, "\n\n")
}
