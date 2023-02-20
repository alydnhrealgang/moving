// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"github.com/alydnhrealgang/moving/assets"
	"github.com/alydnhrealgang/moving/cache"
	"github.com/alydnhrealgang/moving/common/utils"
	"github.com/alydnhrealgang/moving/items"
	"github.com/alydnhrealgang/moving/server/models"
	"github.com/alydnhrealgang/moving/store"
	"github.com/rs/cors"
	"github.com/samber/lo"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/alydnhrealgang/moving/server/restapi/operations"
)

//go:generate swagger generate server --target ..\..\server --name Moving --spec ..\swagger.yaml --principal interface{}

func configureFlags(api *operations.MovingAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.MovingAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.BinConsumer = runtime.ByteStreamConsumer()
	api.JSONConsumer = runtime.JSONConsumer()
	api.MultipartformConsumer = runtime.DiscardConsumer
	api.TxtConsumer = runtime.TextConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// operations.UploadAssetMaxParseMemory = 32 << 20

	api.DownloadAssetHandler = operations.DownloadAssetHandlerFunc(func(params operations.DownloadAssetParams) middleware.Responder {
		if utils.EmptyOrWhiteSpace(params.Path) {
			return operations.NewDownloadAssetBadRequest().WithPayload("Path cannot be null")
		}

		if utils.EmptyOrWhiteSpace(params.Name) {
			return operations.NewDownloadAssetBadRequest().WithPayload("Name cannot be null")
		}

		a, err := assetManager.Query(params.Path, params.Name)
		if nil != err {
			return operations.NewDownloadAssetBadRequest().WithPayload(err.Error())
		}

		if utils.EmptyArray(a) {
			return operations.NewDownloadAssetNotFound()
		}

		if len(a) > 1 {
			return operations.NewDownloadAssetBadRequest().WithPayload(fmt.Sprintf("%d contents found.", len(a)))
		}

		path := fmt.Sprintf("%s/%s", params.Path, params.Name)
		return middleware.ResponderFunc(func(writer http.ResponseWriter, producer runtime.Producer) {
			writer.Header().Set("Content-Type", a[0].ContentType())
			writer.WriteHeader(200)
			err = assetContentStore.DownloadTo(path, a[0].Size(), writer)
			if nil != err {
				api.Logger("Write %s to response error: %s", path, err.Error())
			}
		})
	})
	api.GetAssetHandler = operations.GetAssetHandlerFunc(func(params operations.GetAssetParams) middleware.Responder {
		if utils.EmptyOrWhiteSpace(params.Path) {
			return operations.NewGetAssetBadRequest().WithPayload("Path cannot be null")
		}

		if strings.Compare(params.Name, "*") == 0 {
			params.Name = utils.EmptyString
		}

		data, err := assetManager.Query(params.Path, params.Name)
		if nil != err {
			return operations.NewGetAssetBadRequest().WithPayload(err.Error())
		}

		if utils.EmptyArray(data) {
			return operations.NewGetAssetNotFound()
		}
		return operations.NewGetAssetOK().WithPayload(
			lo.Map(data, func(item *assets.Asset, index int) *models.AssetData {
				return item.ToModel()
			}))
	})
	api.DeleteAssetHandler = operations.DeleteAssetHandlerFunc(func(params operations.DeleteAssetParams) middleware.Responder {
		if utils.EmptyOrWhiteSpace(params.Path) {
			return operations.NewGetAssetBadRequest().WithPayload("Path cannot be null")
		}

		if strings.Compare(params.Name, "*") == 0 {
			data, err := assetManager.Query(params.Path, utils.EmptyString)
			if nil != err {
				return operations.NewDeleteAssetBadRequest().WithPayload(err.Error())
			}
			for _, asset := range data {
				err = assetManager.Delete(asset.Path(), asset.Name())
				if nil != err {
					return operations.NewDeleteAssetBadRequest().WithPayload(err.Error())
				}
			}
			return operations.NewDeleteAssetOK()
		}

		err := assetManager.Delete(params.Path, params.Name)
		if nil != err {
			return operations.NewDeleteAssetBadRequest().WithPayload(err.Error())
		}
		return operations.NewDeleteAssetOK()
	})
	api.UploadAssetHandler = operations.UploadAssetHandlerFunc(func(params operations.UploadAssetParams) middleware.Responder {
		if utils.EmptyOrWhiteSpace(params.Path) {
			return operations.NewUploadAssetBadRequest().WithPayload("Path cannot be null")
		}

		if utils.EmptyOrWhiteSpace(params.Name) {
			return operations.NewUploadAssetBadRequest().WithPayload("Name cannot be null")
		}

		if utils.EmptyOrWhiteSpace(params.ContentType) {
			return operations.NewUploadAssetBadRequest().WithPayload("ContentType cannot be null")
		}

		path := fmt.Sprintf("%s/%s", params.Path, params.Name)
		size, err := assetContentStore.UploadFrom(path, params.ContentType, params.File)
		if nil != err {
			return operations.NewUploadAssetBadRequest().WithPayload(err.Error())
		}
		asset, err := assetManager.Save(params.Path, params.Name, params.ContentType, size)
		if nil != err {
			return operations.NewUploadAssetBadRequest().WithPayload(err.Error())
		}
		return operations.NewUploadAssetOK().WithPayload(&models.AssetDataResponse{
			Code:     0,
			Data:     asset.ToModel(),
			Messages: "OK",
		})
	})

	api.SaveItemHandler = operations.SaveItemHandlerFunc(func(params operations.SaveItemParams) middleware.Responder {
		if utils.EmptyOrWhiteSpace(params.Body.Code) {
			return operations.NewSaveItemBadRequest().WithPayload("code cannot be null")
		}
		if utils.EmptyOrWhiteSpace(params.Body.Type) {
			return operations.NewSaveItemBadRequest().WithPayload("type cannot be null")
		}

		item := items.NewItem(
			params.Body.Code,
			params.Body.BoxCode,
			params.Body.Type,
			params.Body.Name,
			params.Body.Description,
			params.Body.Tags,
			params.Body.Props,
			int(params.Body.Count),
		)
		err := itemManager.SaveItem(item)
		if nil != err {
			return operations.NewSaveItemBadRequest().WithPayload(err.Error())
		}
		return operations.NewSaveItemOK().WithPayload(item.ID())
	})

	api.GetItemByCodeHandler = operations.GetItemByCodeHandlerFunc(
		func(params operations.GetItemByCodeParams) middleware.Responder {

			if utils.EmptyOrWhiteSpace(params.Code) {
				return operations.NewGetItemByCodeBadRequest().WithPayload("code cannot be null")
			}

			if params.ChildrenOnly {
				children, itemExist, err := itemManager.GetChildren(params.Code)
				if nil != err {
					return operations.NewGetItemByCodeBadRequest().WithPayload(err.Error())
				}
				if !itemExist {
					return operations.NewGetItemByCodeNotFound().WithPayload("item not found")
				}

				return operations.NewGetItemByCodeOK().WithPayload(lo.Map(children, itemToModel))
			}

			item, err := itemManager.GetItem(params.Code)
			if nil != err {
				return operations.NewGetItemByCodeBadRequest().WithPayload(err.Error())
			}

			if nil == item {
				return operations.NewGetItemByCodeNotFound().WithPayload("item not found")
			}

			return operations.NewGetItemByCodeOK().WithPayload([]*models.ItemData{itemToModel(item, 0)})
		})

	api.MoveItemsHandler = operations.MoveItemsHandlerFunc(
		func(params operations.MoveItemsParams) middleware.Responder {
			if utils.EmptyArray(params.Body.Codes) {
				return operations.NewMoveItemsBadRequest().WithPayload("codes cannot be null or empty")
			}
			if utils.EmptyOrWhiteSpace(params.Body.To) {
				return operations.NewMoveItemsBadRequest().WithPayload("to cannot be null")
			}

			moved, codesNotFound, err := itemManager.Move(params.Body.Codes, params.Body.To)
			if nil != err {
				return operations.NewMoveItemsBadRequest().WithPayload(err.Error())
			}

			return operations.NewMoveItemsOK().WithPayload(&operations.MoveItemsOKBody{
				CodesNotFound: codesNotFound,
				Moved:         lo.Map(moved, itemToModel),
			})
		})

	api.DeleteItemByCodeHandler = operations.DeleteItemByCodeHandlerFunc(
		func(params operations.DeleteItemByCodeParams) middleware.Responder {
			if utils.EmptyOrWhiteSpace(params.Code) {
				return operations.NewDeleteItemByCodeBadRequest().WithPayload("code cannot be null")
			}

			err := itemManager.DeleteItem(params.Code)
			if nil != err {
				return operations.NewDeleteItemByCodeBadRequest().WithPayload(err.Error())
			}

			return operations.NewDeleteItemByCodeOK()
		})

	api.QueryItemsHandler = operations.QueryItemsHandlerFunc(
		func(params operations.QueryItemsParams) middleware.Responder {
			if utils.EmptyOrWhiteSpace(params.Type) {
				return operations.NewQueryItemsBadRequest().WithPayload("type cannot be null")
			}
			if utils.EmptyOrWhiteSpace(params.Keyword) {
				return operations.NewQueryItemsBadRequest().WithPayload("keyword cannot be null")
			}

			if params.StartIndex < 0 {
				params.StartIndex = 0
			}

			if params.FetchSize < 1 {
				params.FetchSize = 1
			}

			tagName := utils.EmptyString
			if nil != params.TagName {
				tagName = *params.TagName
			}

			result, err := itemManager.QueryItems(
				params.Type, tagName, params.Keyword, params.StartIndex, params.FetchSize)

			if nil != err {
				return operations.NewQueryItemsBadRequest().WithPayload(err.Error())
			}

			return operations.NewQueryItemsOK().WithPayload(lo.Map(result, itemToModel))
		})

	api.SuggestTextsHandler = operations.SuggestTextsHandlerFunc(
		func(params operations.SuggestTextsParams) middleware.Responder {
			if utils.EmptyOrWhiteSpace(params.Name) {
				return operations.NewSuggestTextsBadRequest().WithPayload("name cannot be null")
			}
			if params.Text == "*" {
				params.Text = ""
			}

			texts, err := badgerDB.GetSuggestionTexts(params.Name, params.Text)
			if nil != err {
				return operations.NewSuggestTextsBadRequest().WithPayload(err.Error())
			}

			return operations.NewSuggestTextsOK().WithPayload(texts)
		})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {
		err := badgerDB.Shutdown()
		if nil != err {
			panic(err)
		}
	}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
	memory := cache.CreateMemory()
	badger, err := store.NewBadger(filepath.Join("data", "db"))
	if nil != err {
		panic(err)
	}
	assetContentStore = store.CreateFile(filepath.Join("data", "assets"))
	assetManager = assets.Create(badger, memory)
	assetStore = badger
	itemManager = items.NewManager(badger)
	badgerDB = badger
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	options := cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}
	return cors.New(options).Handler(handler)
}

var (
	badgerDB          *store.Badger
	assetContentStore assets.ContentStore
	assetManager      *assets.Manager
	assetStore        assets.AssetStore
	itemManager       *items.Manager
)

func itemToModel(item *items.Item, _ int) *models.ItemData {
	return &models.ItemData{
		BoxCode:     item.ParentCode(),
		Code:        item.Code(),
		Count:       int64(item.Quantity()),
		Description: item.Description(),
		Name:        item.Name(),
		Props:       item.Props(),
		ServerID:    item.ID(),
		Tags:        item.Tags(),
		Type:        item.Type(),
	}
}
