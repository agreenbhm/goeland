package cmd

const goelandURL = `<a href="https://www.github.com/slurdge/goeland">goeland</a>`

var footers = [...]string{
	`Sent with ❤️ by `,
	`Sent with 💖 by `,
	`Sent with 💙 by `,
	`Sent with 🥰 by `,
	`Brought to you 🐣 by `,
	`Sent quickly ⚡ by `,
	`Sent with a touch of 💐 by `,
	`Sent with a touch of 🌸 by `,
	`Sent with a touch of 🌼 by `,
	`In your 📧 from `,
	`Your 📰 from `,
	`Sent bravely ⚔️ by `,
	`Sent happily 😊 by `,
	`Sent smoothly ✈️ by `,
	`Sent simply 🌞 by `,
	`Sent with 🤘🏻 by `,
}

func init() {
	for i := range footers {
		footers[i] += goelandURL
	}
}
