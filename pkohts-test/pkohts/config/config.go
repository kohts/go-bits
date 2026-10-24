package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Workflow struct {
	Name     string `mapstructure:"name"`
	Filter   string `mapstructure:"filter"`
	Handler  string `mapstructure:"handler"`
	Schedule string `mapstructure:"schedule"`
}

type SmartNotes struct {
	Path               string          `mapstructure:"path"`
	ExcludeDirectories map[string]bool `mapstructure:"exclude_directories"`

	AggregatedOutputPath string `mapstructure:"aggregated_output_path"`
	TodoListOutputPath   string `mapstructure:"todo_list_output_path"`
}

type Config struct {
	Bot struct {
		Name         string `mapstructure:"name"`
		Token        string `mapstructure:"token"`
		RefreshToken string `mapstructure:"refresh_token"`
		SlackChannel string `mapstructure:"slack_channel"`
	} `mapstructure:"bot"`
	Jira struct {
		BaseURL   string `mapstructure:"base_url"`
		AuthToken string `mapstructure:"auth_token"`
		Username  string `mapstructure:"username"`
	} `mapstructure:"jira"`
	YouTrack struct {
		BaseURL   string `mapstructure:"base_url"`
		AuthToken string `mapstructure:"auth_token"`
	} `mapstructure:"youtrack"`
	SmartNotes   SmartNotes                       `mapstructure:"smartnotes"`
	Workflows    []Workflow                       `mapstructure:"workflows"`
	AIAssistants map[string]OpenAIAssistantConfig `mapstructure:"ai_assistants"`
	Assistants   map[string]AssistantConfig       `mapstructure:"assistants"`
}

type AssistantConfig struct {
	ID string `mapstructure:"id"`
}

func (ac *AssistantConfig) AssistantID() string {
	return ac.ID
}

type OpenAIAssistantParameters struct {
	TopP                       *float32 `mapstructure:"top_p"`
	Temperature                *float32 `mapstructure:"temperature"`
	InstructionsFilename       string   `mapstructure:"instructions_filename"`
	InstructionsBackupFilename string   `mapstructure:"instructions_backup_filename"`
	ToolsFilename              string   `mapstructure:"tools_filename"`
	FeedbackFolder             string   `mapstructure:"feedback_folder"`
}

type OpenAIAssistantConfig struct {
	ID                string                    `mapstructure:"id"`
	OpenAIAssistantID string                    `mapstructure:"openai_assistant_id"`
	APIType           string                    `mapstructure:"api_type"`
	APIBaseURL        string                    `mapstructure:"api_base_url"`
	APIVersion        string                    `mapstructure:"api_version"`
	Parameters        OpenAIAssistantParameters `mapstructure:"parameters"`
}

func LoadConfig(configPath string) (config *Config, err error) {
	v := viper.New()

	// Set configuration file and type
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")

	// Read in the YAML configuration
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	// Decode the configuration into the Config struct
	var c Config
	if err := v.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return &c, nil
}
