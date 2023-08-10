package slashcommand

import (
	"discordbot/service"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

var slashCommand *slashcommand

func init() {
	slashCommand = &slashcommand{
		messageService:         service.GetBackMessageService(),
		commands:               []*discordgo.ApplicationCommand{},
		commandHandleFuncMap:   make(map[string]func(*discordgo.Session, *discordgo.InteractionCreate)),
		componentHandleFuncMap: make(map[string]func(*discordgo.Session, *discordgo.InteractionCreate)),
	}
}

type slashcommand struct {
	messageService         service.BackMessageService
	commands               []*discordgo.ApplicationCommand
	registeredCommands     []*discordgo.ApplicationCommand
	commandHandleFuncMap   map[string]func(*discordgo.Session, *discordgo.InteractionCreate)
	componentHandleFuncMap map[string]func(*discordgo.Session, *discordgo.InteractionCreate)
}

func AddSlashCommand(s *discordgo.Session) {
	registSlashCommond(s)
}

func registSlashCommond(s *discordgo.Session) {
	setMessageCommand()
	deleteMessageComand()
	allKey()

	slashCommand.registeredCommands = make([]*discordgo.ApplicationCommand, len(slashCommand.commands))

	for i, v := range slashCommand.commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)

		if err != nil {
			log.Err(err).Msgf("Create command %v error", v.Name)
			continue
		}

		slashCommand.registeredCommands[i] = cmd
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			if h, ok := slashCommand.commandHandleFuncMap[i.ApplicationCommandData().Name]; ok {
				h(s, i)
			}
		case discordgo.InteractionMessageComponent:
			if h, ok := slashCommand.componentHandleFuncMap[i.MessageComponentData().CustomID]; ok {
				h(s, i)
			}
		}
	})
}

type commandHandleFunc func(context)

type componentHandleFunc func(context)

type slashCommandRegistry struct {
	command             *discordgo.ApplicationCommand
	commandHandleFunc   commandHandleFunc
	componentId         string
	componentHandleFunc componentHandleFunc
}

type context struct {
	session             *discordgo.Session
	interactionCreate   *discordgo.InteractionCreate
	commandOptionArgMap map[string]string
	componentArgs       []string
}

func (sc *slashcommand) rCommand(command slashCommandRegistry) {
	if command.commandHandleFunc != nil {
		sc.commands = append(sc.commands, command.command)
		sc.commandHandleFuncMap[command.command.Name] = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			context := context{
				session:             s,
				interactionCreate:   i,
				commandOptionArgMap: handleArg(i.ApplicationCommandData().Options),
			}

			command.commandHandleFunc(context)
		}
	}

	if command.componentId != "" && command.componentHandleFunc != nil {
		sc.componentHandleFuncMap[command.componentId] = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			context := context{
				session:           s,
				interactionCreate: i,
				componentArgs:     i.MessageComponentData().Values,
			}

			command.componentHandleFunc(context)
		}
	}

}

func handleArg(options []*discordgo.ApplicationCommandInteractionDataOption) map[string]string {
	args := make(map[string]string, len(options))

	for _, option := range options {
		args[option.Name] = option.StringValue()
	}

	return args
}

func DeleteSlashCommand(s *discordgo.Session) {
	for _, cmd := range slashCommand.registeredCommands {
		err := s.ApplicationCommandDelete(s.State.User.ID, "", cmd.ID)

		if err != nil {
			log.Err(err).Msg("DeleteCommand Error")
		}
	}
}
