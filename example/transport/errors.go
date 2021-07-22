// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import "github.com/rs/zerolog"

type withErrorCode interface {
	Code() int
}

type strError string

func (e strError) Error() string {
	return string(e)
}

func ExitOnError(log zerolog.Logger, err error, msg string) {
	if err != nil {
		log.Panic().Err(err).Msg(msg)
	}
}
