// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"github.com/petergtz/pegomock"
	"reflect"

	jobs "github.com/runatlantis/atlantis/server/jobs"
)

func AnyJobsProjectCommandOutputHandler() jobs.ProjectCommandOutputHandler {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(jobs.ProjectCommandOutputHandler))(nil)).Elem()))
	var nullValue jobs.ProjectCommandOutputHandler
	return nullValue
}

func EqJobsProjectCommandOutputHandler(value jobs.ProjectCommandOutputHandler) jobs.ProjectCommandOutputHandler {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue jobs.ProjectCommandOutputHandler
	return nullValue
}

func NotEqJobsProjectCommandOutputHandler(value jobs.ProjectCommandOutputHandler) jobs.ProjectCommandOutputHandler {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue jobs.ProjectCommandOutputHandler
	return nullValue
}

func JobsProjectCommandOutputHandlerThat(matcher pegomock.ArgumentMatcher) jobs.ProjectCommandOutputHandler {
	pegomock.RegisterMatcher(matcher)
	var nullValue jobs.ProjectCommandOutputHandler
	return nullValue
}