package tgbotapi

import "time"

// MessageReaction is a reaction to a message
type MessageReaction struct {
	// Type of the reaction
	// can be "emoji" or "custom_emoji"
	Type string `json:"type"`
	// Emoji Reaction emoji.
	// Currently, it can be one of "👍", "👎", "❤", "🔥", "🥰", "👏", "😁", "🤔",
	// "🤯", "😱", "🤬", "😢", "🎉", "🤩", "🤮", "💩", "🙏", "👌", "🕊", "🤡", "🥱",
	// "🥴", "😍", "🐳", "❤‍🔥", "🌚", "🌭", "💯", "🤣", "⚡", "🍌", "🏆", "💔", "🤨",
	// "😐", "🍓", "🍾", "💋", "🖕", "😈", "😴", "😭", "🤓", "👻", "👨‍💻", "👀", "🎃",
	// "🙈", "😇", "😨", "🤝", "✍", "🤗", "🫡", "🎅", "🎄", "☃", "💅", "🤪", "🗿", "🆒",
	// "💘", "🙉", "🦄", "😘", "💊", "🙊", "😎", "👾", "🤷‍♂", "🤷", "🤷‍♀", "😡"
	Emoji string `json:"emoji,omitempty"`
	// CustomEmoji Custom emoji identifier
	CustomEmoji string `json:"custom_emoji,omitempty"`
}

// MessageReactionUpdated represents a change of a reaction on a message performed by a user
type MessageReactionUpdated struct {
	// Chat The chat containing the message the user reacted to
	Chat Chat `json:"chat"`
	// MessageID Unique identifier of the message inside the chat
	MessageID int `json:"message_id"`
	// User The user that changed the reaction, if the user isn't anonymous
	//
	// Optional
	User *User `json:"user,omitempty"`
	// ActorChat The chat on behalf of which the reaction was changed, if the user is anonymous
	//
	// Optional
	ActorChat *Chat `json:"actor_chat,omitempty"`
	// Date of the change in Unix time
	Date int `json:"date"`
	// OldReaction Previous list of reaction types that were set by the user
	OldReactions []MessageReaction `json:"old_reaction"`
	// NewReaction New list of reaction types that were set by the user
	NewReactions []MessageReaction `json:"new_reaction"`
}

// Time converts the MessageReactionUpdated timestamp into a Time.
func (m *MessageReactionUpdated) Time() time.Time {
	return time.Unix(int64(m.Date), 0)
}

// ReactionCount Represents a reaction added to a message along with the number of times it was added.
type ReactionCount struct {
	MessageReaction MessageReaction `json:"type"`
	TotalCount      int             `json:"total_count"`
}

// MessageReactionCountUpdated represents reaction changes on a message with anonymous reactions
type MessageReactionCountUpdated struct {
	// Chat The chat containing the message
	Chat Chat `json:"chat"`
	// MessageID Unique message identifier inside the chat
	MessageID int `json:"message_id"`
	// Date of the change in Unix time
	Date int `json:"date"`
	// Reactions List of reactions that are present on the message
	Reactions []ReactionCount `json:"reactions"`
}

type SetReactionConfig struct {
	BaseChat
	MessageID int               `json:"message_id"`
	Reaction  []MessageReaction `json:"reaction"`
	IsBig     bool              `json:"is_big"`
}

func (config SetReactionConfig) method() string {
	return "setMessageReaction"
}

func (config SetReactionConfig) params() (Params, error) {
	params := make(Params)

	params.AddBool("is_big", config.IsBig)
	params.AddNonZero("message_id", config.MessageID)
	err := params.AddInterface("reaction", config.Reaction)
	if err != nil {
		return params, err
	}

	p1, _ := config.BaseChat.params()
	params.Merge(p1)

	return params, nil
}

// SetEmojiReaction creates a new forum topic with name
// Currently, it can be one of "👍", "👎", "❤", "🔥", "🥰", "👏", "😁", "🤔",
// "🤯", "😱", "🤬", "😢", "🎉", "🤩", "🤮", "💩", "🙏", "👌", "🕊", "🤡", "🥱",
// "🥴", "😍", "🐳", "❤‍🔥", "🌚", "🌭", "💯", "🤣", "⚡", "🍌", "🏆", "💔", "🤨",
// "😐", "🍓", "🍾", "💋", "🖕", "😈", "😴", "😭", "🤓", "👻", "👨‍💻", "👀", "🎃",
// "🙈", "😇", "😨", "🤝", "✍", "🤗", "🫡", "🎅", "🎄", "☃", "💅", "🤪", "🗿", "🆒",
// "💘", "🙉", "🦄", "😘", "💊", "🙊", "😎", "👾", "🤷‍♂", "🤷", "🤷‍♀", "😡"
func (bot *BotAPI) SetEmojiReaction(chatID int64, messageID int, reaction string, isBig bool) (Message, error) {
	return bot.Send(SetReactionConfig{
		BaseChat: BaseChat{
			ChatID: chatID,
		},
		MessageID: messageID,
		Reaction: []MessageReaction{
			{
				Type:  "emoji",
				Emoji: reaction,
			},
		},
		IsBig: isBig,
	})
}
