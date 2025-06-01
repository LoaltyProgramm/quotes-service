package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/LoaltyProgramm/quotes-service/internal/models/quotes"
	quoteservice "github.com/LoaltyProgramm/quotes-service/internal/services/quote_service"
	"github.com/LoaltyProgramm/quotes-service/internal/utils/bind"
	writejson "github.com/LoaltyProgramm/quotes-service/internal/utils/write_json"
)

const (
	ERRORCREATEQUOTE001 = "author or quote is not empty"
	ERRORCREATEQUOTE002 = "such a record already exists"
)

const (
	ERRORLISTQUOTES001 = "quotes is not found"
)

const (
	ERRORGETRANDOMQUOTE001 = "quotes is not found"
)

const (
	ERRORLISTQUOTESBYAUTHOR001 = "the data cannot be a number"
	ERRORLISTQUOTESBYAUTHOR002 = "author is not found"
)

const (
	ERRORREMOVEQUOTE001 = "the identifier cannot be a letter"
	ERRORREMOVEQUOTE002 = "no record was found for this id"
)

type Handlers struct {
	quoteService quoteservice.QuoteService
}

func NewHandlers(quoteService quoteservice.QuoteService) Handlers {
	return Handlers{
		quoteService: quoteService,
	}
}

func (h *Handlers) QuoteHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var quote quotes.Quote
		if err := bind.Bind(r, &quote); err != nil {
			writejson.WriteJson(w, map[string]string{"error": "bad request"}, http.StatusBadRequest)
			log.Printf("ERROR Method:%s, url-path:%s, error:%s", r.Method, r.URL.Path, err.Error())
			return
		}

		if err := h.quoteService.CreateQuote(quote); err != nil {
			if err.Error() == ERRORCREATEQUOTE001 {
				writejson.WriteJson(w, map[string]string{"error": ERRORCREATEQUOTE001}, http.StatusBadRequest)
				log.Printf("ERROR Method:%s, url-path:%s, error:%s", r.Method, r.URL.Path, err.Error())
				return
			} else if err.Error() == ERRORCREATEQUOTE002 {
				writejson.WriteJson(w, map[string]string{"error": ERRORCREATEQUOTE002}, http.StatusBadRequest)
				log.Printf("ERROR Method:%s, url-path:%s, error:%s", r.Method, r.URL.Path, err.Error())
				return
			} else {
				writejson.WriteJson(w, map[string]string{"error": "bad request"}, http.StatusBadRequest)
				log.Printf("ERROR Method:%s, url-path:%s, error:%s", r.Method, r.URL.Path, err.Error())
				return
			}
		}

		log.Printf("Method:%s, url-path:%s,", r.Method, r.URL.Path)
		writejson.WriteJson(w, map[string]string{"status": "created"}, http.StatusCreated)

	case http.MethodGet:
		author := r.URL.Query().Get("author")
		if author == "" {
			quotes, err := h.quoteService.ListQuotes()
			if err != nil {
				if err.Error() == ERRORLISTQUOTES001 {
					writejson.WriteJson(w, map[string]string{"error": ERRORLISTQUOTES001}, http.StatusNotFound)
					log.Printf("ERROR Method:%s, url-path:%s, error:%s", r.Method, r.URL.Path, err.Error())
					return
				} else {
					writejson.WriteJson(w, map[string]string{"error": "repeat later"}, http.StatusInternalServerError)
					log.Printf("ERROR Method:%s, url-path:%s, error:%s", r.Method, r.URL.Path, err.Error())
					return
				}
			}

			log.Printf("Method:%s, url-path:%s,", r.Method, r.URL.Path)
			writejson.WriteJson(w, quotes, http.StatusOK)
		} else {
			quotes, err := h.quoteService.ListQuotesByAuthor(author)
			if err != nil {
				if err.Error() == ERRORLISTQUOTESBYAUTHOR001 {
					writejson.WriteJson(w, map[string]string{"error": ERRORLISTQUOTESBYAUTHOR001}, http.StatusBadRequest)
					log.Printf("ERROR Method:%s, url-path:%s, error:%s", r.Method, r.URL.Path, err.Error())
					return
				} else if err.Error() == ERRORLISTQUOTESBYAUTHOR002 {
					writejson.WriteJson(w, map[string]string{"error": ERRORLISTQUOTESBYAUTHOR002}, http.StatusNotFound)
					log.Printf("ERROR Method:%s, url-path:%s, error:%s", r.Method, r.URL.Path, err.Error())
					return
				} else {
					writejson.WriteJson(w, map[string]string{"error": "repeat later"}, http.StatusInternalServerError)
					log.Printf("ERROR Method:%s, url-path:%s, error:%s", r.Method, r.URL.Path, err.Error())
					return
				}
			}

			log.Printf("Method:%s, url-path:%s", r.Method, r.URL.Path)
			writejson.WriteJson(w, quotes, http.StatusOK)
		}
	default:
		log.Printf("Method:%s, url-path:%s", r.Method, r.URL.Path)
		writejson.WriteJson(w, map[string]string{"error": fmt.Sprintf("This method %s is not supported", r.Method)}, http.StatusOK)
	}
}

func (h *Handlers) RandomQuoteHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		randomQuote, err := h.quoteService.GetQuoteRandom()
		if err != nil {
			if err.Error() == ERRORGETRANDOMQUOTE001 {
				writejson.WriteJson(w, map[string]string{"error": ERRORGETRANDOMQUOTE001}, http.StatusNotFound)
				log.Printf("ERROR Method:%s, url-path:%s, error:%s", r.Method, r.URL.Path, err.Error())
				return
			} else {
				writejson.WriteJson(w, map[string]string{"error": "repeat later"}, http.StatusInternalServerError)
				log.Printf("ERROR Method:%s, url-path:%s, error:%s", r.Method, r.URL.Path, err.Error())
				return
			}
		}

		log.Printf("Method:%s, url-path:%s", r.Method, r.URL.Path)
		writejson.WriteJson(w, randomQuote, http.StatusOK)
	default:
		log.Printf("Method:%s, url-path:%s", r.Method, r.URL.Path)
		writejson.WriteJson(w, map[string]string{"error": fmt.Sprintf("This method %s is not supported", r.Method)}, http.StatusOK)
	}
}

func (h *Handlers) DeleteQuoteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/quotes/")
	if idStr == "" {
		writejson.WriteJson(w, map[string]string{"error": "missing id"}, http.StatusBadRequest)
		return
	}

	err := h.quoteService.RemoveQuoteById(idStr)
	if err != nil {
		if err.Error() == ERRORREMOVEQUOTE001 {
			writejson.WriteJson(w, map[string]string{"error": ERRORREMOVEQUOTE001}, http.StatusBadRequest)
			log.Printf("ERROR Method:%s, url-path:%s, error:%s", r.Method, r.URL.Path, err.Error())
			return
		} else if err.Error() == ERRORREMOVEQUOTE002 {
			writejson.WriteJson(w, map[string]string{"error": ERRORREMOVEQUOTE002}, http.StatusBadRequest)
			log.Printf("ERROR Method:%s, url-path:%s, error:%s", r.Method, r.URL.Path, err.Error())
			return
		} else {
			writejson.WriteJson(w, map[string]string{"error": "repeat later"}, http.StatusInternalServerError)
			log.Printf("ERROR Method:%s, url-path:%s, error:%s", r.Method, r.URL.Path, err.Error())
			return
		}
	}

	writejson.WriteJson(w, map[string]string{"delete": idStr}, http.StatusOK)
}

func (h *Handlers) InitHandlers() {
	http.HandleFunc("/quotes", h.QuoteHandler)
	http.HandleFunc("/quotes/random", h.RandomQuoteHandler)
	http.HandleFunc("/quotes/", h.DeleteQuoteHandler)
}
