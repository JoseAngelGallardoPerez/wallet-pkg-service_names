package service_names

type serverName struct {
	Internal string
	Public   string
}

var (
	Accounts = serverName{
		Internal: "accounts",
		Public:   "Accounts",
	}
	Currencies = serverName{
		Internal: "currencies",
		Public:   "Currencies",
	}
	Files = serverName{
		Internal: "files",
		Public:   "Files",
	}
	Messages = serverName{
		Internal: "messages",
		Public:   "Messages",
	}
	Notifications = serverName{
		Internal: "notifications",
		Public:   "Notifications",
	}
	Permissions = serverName{
		Internal: "permissions",
		Public:   "Permissions",
	}
	Settings = serverName{
		Internal: "settings",
		Public:   "Settings",
	}
	Swagger = serverName{
		Internal: "swagger",
		Public:   "Swagger",
	}
	Users = serverName{
		Internal: "users",
		Public:   "Users",
	}
	Logs = serverName{
		Internal: "logs",
		Public:   "Logs",
	}
	News = serverName{
		Internal: "news",
		Public:   "News",
	}
	Customization = serverName{
		Internal: "customization",
		Public:   "Customization",
	}
	Newsletters = serverName{
		Internal: "newsletters",
		Public:   "Newsletters",
	}
	Verification = serverName{
		Internal: "verification",
		Public:   "Verification",
	}
	Extensions = serverName{
		Internal: "extensions",
		Public:   "Extensions",
	}
)
