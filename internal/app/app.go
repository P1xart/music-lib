package app

import "log/slog"

func Run() {
	log := slog.New(slog.Default().Handler())

    log.Info("Starting app")
}
