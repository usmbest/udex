package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/citypayorg/udex/udex-backend/interfaces"
	"github.com/citypayorg/udex/udex-backend/services"
	"github.com/citypayorg/udex/udex-backend/types"
	"github.com/citypayorg/udex/udex-backend/utils/httputils"
	"github.com/gorilla/mux"
)

type pairEndpoint struct {
	pairService  interfaces.PairService
	tokenService interfaces.TokenService
}

// ServePairResource sets up the routing of pair endpoints and the corresponding handlers.
func ServePairResource(
	r *mux.Router,
	p interfaces.PairService,
	t interfaces.TokenService,
) {
	e := &pairEndpoint{p, t}
	r.HandleFunc("/pairs/create", e.HandleCreatePairs).Methods("POST")
	r.HandleFunc("/pair/create", e.HandleCreatePair).Methods("POST")
	r.HandleFunc("/pairs", e.HandleGetPairs).Methods("GET")
	r.HandleFunc("/pair", e.HandleGetPair).Methods("GET")
	r.HandleFunc("/pairs/data", e.HandleGetPairData).Methods("GET")
}

func (e *pairEndpoint) HandleCreatePairs(w http.ResponseWriter, r *http.Request) {
	token := types.Token{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&token)
	if err != nil {
		logger.Info(err)
		httputils.WriteError(w, http.StatusBadRequest, "Invalid Payload")
		return
	}

	if token.Asset == "" && token.Symbol != "" {
		t, err := e.tokenService.GetBySymbol(token.Symbol)
		if err != nil {
			logger.Info(err)
			httputils.WriteError(w, http.StatusBadRequest, "Bad symbol")
			return
		}
		if t == nil {
			logger.Info(err)
			httputils.WriteError(w, http.StatusBadRequest, "symbol not found")
			return
		}
		token = *t
	}

	defer r.Body.Close()

	pairs, err := e.pairService.CreatePairs(token.Asset)
	if err != nil {
		switch err {
		case services.ErrPairExists:
			httputils.WriteError(w, http.StatusBadRequest, "Pair exists")
			return
		case services.ErrBaseTokenNotFound:
			httputils.WriteError(w, http.StatusBadRequest, "Base token not found")
			return
		case services.ErrQuoteTokenNotFound:
			httputils.WriteError(w, http.StatusBadRequest, "Quote token not found")
			return
		case services.ErrQuoteTokenInvalid:
			httputils.WriteError(w, http.StatusBadRequest, "Quote token invalid (token is not registered as quote)")
			return
		case services.ErrNoAsset:
			httputils.WriteError(w, http.StatusBadRequest, "Asset not found")
			return
		default:
			logger.Error(err)
			httputils.WriteError(w, http.StatusInternalServerError, "Internal server error")
			return
		}
	}

	if len(pairs) == 0 {
		httputils.WriteError(w, http.StatusBadRequest, "Pairs already exist")
		return
	}

	httputils.WriteJSON(w, http.StatusCreated, pairs)
}

func (e *pairEndpoint) HandleCreatePair(w http.ResponseWriter, r *http.Request) {
	p := &types.Pair{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(p)
	if err != nil {
		logger.Info(err)
		httputils.WriteError(w, http.StatusBadRequest, "Invalid payload")
		return
	}

	defer r.Body.Close()

	err = p.ValidateAssets()
	if err != nil {
		httputils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = e.pairService.Create(p)
	if err != nil {
		switch err {
		case services.ErrPairExists:
			httputils.WriteError(w, http.StatusBadRequest, "Pair exists")
			return
		case services.ErrBaseTokenNotFound:
			httputils.WriteError(w, http.StatusBadRequest, "Base token not found")
			return
		case services.ErrQuoteTokenNotFound:
			httputils.WriteError(w, http.StatusBadRequest, "Quote token not found")
			return
		case services.ErrQuoteTokenInvalid:
			httputils.WriteError(w, http.StatusBadRequest, "Quote token invalid (token is not registered as quote")
			return
		case services.ErrNoAsset:
			httputils.WriteError(w, http.StatusBadRequest, "Asset not found")
			return
		default:
			logger.Error(err)
			httputils.WriteError(w, http.StatusInternalServerError, "")
			return
		}
	}

	httputils.WriteJSON(w, http.StatusCreated, p)
}

func (e *pairEndpoint) HandleGetPairs(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()

	var res []types.Pair
	var err error

	switch v.Get("listed") {
	case "":
		res, err = e.pairService.GetAll()
	case "true":
		res, err = e.pairService.GetListedPairs()
	case "false":
		res, err = e.pairService.GetUnlistedPairs()
	}

	if err != nil {
		logger.Error(err)
		httputils.WriteError(w, http.StatusInternalServerError, "")
		return
	}

	if res == nil {
		httputils.WriteJSON(w, http.StatusOK, []types.Pair{})
		return
	}

	httputils.WriteJSON(w, http.StatusOK, res)
}

func (e *pairEndpoint) HandleGetPair(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	baseToken := v.Get("baseToken")
	quoteToken := v.Get("quoteToken")

	if baseToken == "" {
		httputils.WriteError(w, http.StatusBadRequest, "baseToken Parameter missing")
		return
	}

	if quoteToken == "" {
		httputils.WriteError(w, http.StatusBadRequest, "quoteToken Parameter missing")
		return
	}

	if !isValidAsset(baseToken) {
		httputils.WriteError(w, http.StatusBadRequest, "Invalid Base Token Asset")
		return
	}

	if !isValidAsset(quoteToken) {
		httputils.WriteError(w, http.StatusBadRequest, "Invalid Quote Token Asset")
		return
	}

	baseAsset := baseToken
	quoteAsset := quoteToken
	res, err := e.pairService.GetByAsset(baseAsset, quoteAsset)
	if err != nil {
		logger.Error(err)
		httputils.WriteError(w, http.StatusInternalServerError, "")
		return
	}

	if res == nil {
		httputils.WriteJSON(w, http.StatusOK, []types.Pair{})
		return
	}

	httputils.WriteJSON(w, http.StatusOK, res)
}

func (e *pairEndpoint) HandleGetPairData(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	baseToken := v.Get("baseToken")
	quoteToken := v.Get("quoteToken")
	exact := v.Get("exact")
	simple := v.Get("simple")

	if simple == "true" && exact == "true" {
		httputils.WriteError(w, http.StatusBadRequest, "'simple' and 'exact' param can not both be true")
		return
	}

	//Return simplified version of token pair data
	if baseToken == "" && quoteToken == "" && simple == "true" {
		res, err := e.pairService.GetAllSimplifiedTokenPairData()
		if err != nil {
			logger.Error(err)
			httputils.WriteError(w, http.StatusInternalServerError, "")
			return
		}

		if res == nil {
			httputils.WriteJSON(w, http.StatusOK, []types.Pair{})
			return
		}

		httputils.WriteJSON(w, http.StatusOK, res)
		return
	}

	//Return formal version of token pair data
	if baseToken == "" && quoteToken == "" && exact == "true" {
		res, err := e.pairService.GetAllExactTokenPairData()
		if err != nil {
			logger.Error(err)
			httputils.WriteError(w, http.StatusInternalServerError, "")
			return
		}

		if res == nil {
			httputils.WriteJSON(w, http.StatusOK, []types.Pair{})
			return
		}

		httputils.WriteJSON(w, http.StatusOK, res)
		return
	}

	//Return the simplified version of token pair data
	if baseToken == "" && quoteToken == "" {
		res, err := e.pairService.GetAllTokenPairData()
		if err != nil {
			logger.Error(err)
			httputils.WriteError(w, http.StatusInternalServerError, "")
			return
		}

		if res == nil {
			httputils.WriteJSON(w, http.StatusOK, []types.Pair{})
			return
		}

		httputils.WriteJSON(w, http.StatusOK, res)
		return
	}

	if quoteToken == "" {
		httputils.WriteError(w, http.StatusBadRequest, "quoteToken Parameter missing")
		return
	}

	if baseToken == "" {
		httputils.WriteError(w, http.StatusBadRequest, "baseToken Parameter missing")
		return
	}

	if !isValidAsset(baseToken) {
		httputils.WriteError(w, http.StatusBadRequest, "Invalid Base Token Asset")
		return
	}

	if !isValidAsset(quoteToken) {
		httputils.WriteError(w, http.StatusBadRequest, "Invalid Quote Token Asset")
		return
	}

	baseAsset := baseToken
	quoteAsset := quoteToken

	res, err := e.pairService.GetTokenPairData(baseAsset, quoteAsset)
	if err != nil {
		logger.Error(err)
		httputils.WriteError(w, http.StatusInternalServerError, "")
		return
	}

	if res == nil {
		httputils.WriteJSON(w, http.StatusOK, []types.Pair{})
		return
	}

	httputils.WriteJSON(w, http.StatusOK, res)
}
