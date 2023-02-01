// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"github.com/alydnhrealgang/moving/assets"
	"github.com/alydnhrealgang/moving/cache"
	"github.com/alydnhrealgang/moving/common/utils"
	"github.com/alydnhrealgang/moving/server/models"
	"github.com/alydnhrealgang/moving/store"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/rs/cors"
	"github.com/samber/lo"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/alydnhrealgang/moving/server/restapi/operations"
)

//go:generate swagger generate server --target ..\..\server --name MovingAPI --spec ..\swagger.yaml --principal interface{}

func configureFlags(api *operations.MovingAPIAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.MovingAPIAPI) http.Handler {
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

	if api.DownloadAssetHandler == nil {
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
	}
	if api.GetAssetHandler == nil {
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
	}
	if api.UploadAssetHandler == nil {
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
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {
		err := assetStore.Shutdown()
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
	assetContentStore assets.ContentStore
	assetManager      *assets.Manager
	assetStore        assets.AssetStore
)
