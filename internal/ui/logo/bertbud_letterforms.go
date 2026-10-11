package logo

import "github.com/MakeNowJust/heredoc"

// The BertBud wordmark needs B, T and D, which upstream Crush doesn't draw.
// They live in their own file so upstream merges into letterforms.go stay
// clean.

// LetterB renders the letter B in a stylized way. It takes an integer that
// determines how many cells to stretch the letter. If the stretch is less than
// 1, it defaults to no stretching.
func LetterB(stretch bool) string {
	// Here's what we're making:
	//
	// █▀▀▀▄
	// █▀▀▀▄
	// ▀▀▀▀

	left := heredoc.Doc(`
		█
		█
		▀
	`)
	center := heredoc.Doc(`
		▀
		▀
		▀
	`)
	right := heredoc.Doc(`
		▄
		▄

	`)
	return joinLetterform(
		left,
		stretchLetterformPart(center, letterformProps{
			stretch:    stretch,
			width:      3,
			minStretch: 7,
			maxStretch: 12,
		}),
		right,
	)
}

// LetterD renders the letter D in a stylized way. It takes an integer that
// determines how many cells to stretch the letter. If the stretch is less than
// 1, it defaults to no stretching.
func LetterD(stretch bool) string {
	// Here's what we're making:
	//
	// █▀▀▀▄
	// █   █
	// ▀▀▀▀

	left := heredoc.Doc(`
		█
		█
		▀
	`)
	center := heredoc.Doc(`
		▀

		▀
	`)
	right := heredoc.Doc(`
		▄
		█

	`)
	return joinLetterform(
		left,
		stretchLetterformPart(center, letterformProps{
			stretch:    stretch,
			width:      3,
			minStretch: 7,
			maxStretch: 12,
		}),
		right,
	)
}

// LetterT renders the letter T in a stylized way. It takes an integer that
// determines how many cells to stretch the letter. If the stretch is less than
// 1, it defaults to no stretching.
func LetterT(stretch bool) string {
	// Here's what we're making:
	//
	// ▀▀█▀▀
	//   █
	//   ▀

	arm := heredoc.Doc(`
		▀


	`)
	stem := heredoc.Doc(`
		█
		█
		▀
	`)
	stretchedArm := stretchLetterformPart(arm, letterformProps{
		stretch:    stretch,
		width:      2,
		minStretch: 4,
		maxStretch: 7,
	})
	return joinLetterform(
		stretchedArm,
		stem,
		stretchedArm,
	)
}
