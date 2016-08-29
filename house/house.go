package house

import "strings"

// Embed embeds a noun phrase as the object of relative clause with a transitive verb.
func Embed(relPhrase, nounPhrase string) string {
	return relPhrase + " " + nounPhrase
}

// Verse generates a verse of a song.
func Verse(subject string, relPhrases []string, nounPhrase string) string {
	if len(relPhrases) == 0 {
		return Embed(subject, nounPhrase)
	}
	return Verse(subject, relPhrases[:len(relPhrases)-1], Embed(relPhrases[len(relPhrases)-1], nounPhrase))
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

// Song generates the full text of "The House That Jack Built".
func Song() string {
	var s []string
	for i := len(phrases) - 1; i >= 0; i-- {
		s = append(s, Verse("This is", phrases[i:], "Jack built."))
	}
	return strings.Join(s, "\n\n")
}
