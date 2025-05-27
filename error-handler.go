package messaging

import logging "github.com/payly-solucoes-de-pagamentos/golang-logging"

func failOnError(logger *logging.Logger, err error, message string) {
	if err == nil {
		return
	}
	logger.Standard.Panic().AnErr(message, err)
	panic(err)
}
