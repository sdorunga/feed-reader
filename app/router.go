package app

import (
	"github.com/kataras/muxie"

	"feed-reader/feedlist"
	"feed-reader/fetcher"

	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"strings"
)

func InitRouter(feedListStore feedlist.FeedListStore, feedFetcher fetcher.Fetcher) *muxie.Mux {
	mux := muxie.NewMux()

	api := mux.Of("/api")
	api.Use(apiMiddleware)
	api.HandleFunc("/*path", apiHandler(NotFoundHandler{}))
	api.Handle("/feeds/:id", muxie.Methods().
		HandleFunc(http.MethodGet, apiHandler(GETFeedHandler{store: feedListStore, fetcher: feedFetcher})))
	api.Handle("/feeds", muxie.Methods().
		HandleFunc(http.MethodGet, apiHandler(GETFeedListHandler{store: feedListStore})).
		HandleFunc(http.MethodPost, apiHandler(POSTFeedHandler{store: feedListStore, fetcher: feedFetcher})))

	return mux
}

func apiMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")
		w.Header().Set("Content-Type", "application/json;charset=utf8")

		for _, v := range strings.Split(contentType, ",") {
			t, _, err := mime.ParseMediaType(v)
			if err != nil {
				w.WriteHeader(http.StatusUnsupportedMediaType)
				w.Write([]byte(fmt.Sprintf(`{"error": "Media type (%s) not parseable"}`, v)))
				return
			}
			if t == "application/json" {
				next.ServeHTTP(w, r)
				return
			}
		}
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf(`{"error": "Media type (%s) not supported"}`, contentType)))
	})
}

type APIHandler interface {
	Handle(body []byte, params map[string]string, queryParams url.Values) (interface{}, error)
}

func apiHandler(handler APIHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, jsonError(err))
		}
		defer r.Body.Close()
		// Build up the path params from muxie
		params := map[string]string{}
		for _, keyPair := range muxie.GetParams(w) {
			params[keyPair.Key] = keyPair.Value
		}
		// Get query and request body params populated
		err = r.ParseForm()
		if err != nil {
			fmt.Fprintf(w, jsonError(err))
		}

		rsp, err := handler.Handle(b, params, r.Form)
		if err != nil {
			// If we specifically return an HttpError we can use its code,
			// otherwise we just default to Internal Server Error
			if httpErr, ok := err.(HttpError); ok {
				w.WriteHeader(httpErr.StatusCode())
				fmt.Fprintf(w, jsonError(err))
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, jsonError(err))
			return
		}

		jsonRsp, err := json.MarshalIndent(rsp, "", "\t")
		if err != nil {
			fmt.Fprintf(w, jsonError(err))
			return
		}

		fmt.Fprintf(w, string(jsonRsp))
	}
}

// NotFoundHandler is the default handler for any rout that is not provided. It
// returns a JSON formatted error
type NotFoundHandler struct {
}

func (handler NotFoundHandler) Handle(body []byte, params map[string]string, queryParams url.Values) (interface{}, error) {
	return nil, NotFoundError{err: errors.New("Endpoint not found")}
}

func jsonError(err error) string {
	return fmt.Sprintf("{\"error\": \"%s\"}", err.Error())
}
