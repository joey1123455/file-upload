package config

type Config struct {
	DBUri    string `mapstructure:"MONGO_URI"`
	CloudKey string `mapstructure:"CLOUDINARY"`
}
