package wire

import (
	goConfig "github.com/liampulles/go-config"

	"github.com/liampulles/matchstick-video/pkg/adapter/config"
	adapterDb "github.com/liampulles/matchstick-video/pkg/adapter/db"
	"github.com/liampulles/matchstick-video/pkg/adapter/db/sql"
	"github.com/liampulles/matchstick-video/pkg/adapter/http"
	"github.com/liampulles/matchstick-video/pkg/adapter/http/json"
	"github.com/liampulles/matchstick-video/pkg/domain"
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
	"github.com/liampulles/matchstick-video/pkg/driver/db"
	"github.com/liampulles/matchstick-video/pkg/driver/http/mux"
	"github.com/liampulles/matchstick-video/pkg/usecase/inventory"
)

// CreateApp creates a runnable for the entrypoint of the
// application
func CreateApp(source goConfig.Source) domain.Runnable {
	factory, err := CreateServerFactory(source)
	if err != nil {
		return func() error {
			return err
		}
	}
	return factory.Create()
}

// CreateServerFactory injects all the dependencies needed to create
// http.ServerFactory
func CreateServerFactory(source goConfig.Source) (http.ServerFactory, error) {
	// Each "tap" below indicates a level of dependency
	configStore, err := config.NewStoreImpl(
		source,
	)
	if err != nil {
		return nil, err
	}
	errorParser := adapterDb.NewErrorParserImpl()

	// --- NEXT TAP ---
	helperService := sql.NewHelperServiceImpl(errorParser)
	databaseService, err := db.NewDatabaseServiceImpl(
		configStore,
	)
	if err != nil {
		return nil, err
	}
	inventoryItemConstructor := entity.NewInventoryItemConstructorImpl()
	muxWrapper := mux.NewWrapperImpl()

	// --- NEXT TAP ---
	inventoryRepository := sql.NewInventoryRepositoryImpl(
		databaseService,
		helperService,
		inventoryItemConstructor,
	)
	entityFactory := inventory.NewEntityFactoryImpl(
		inventoryItemConstructor,
	)
	entityModifier := inventory.NewEntityModifierImpl()
	voFactory := inventory.NewVOFactoryImpl()
	ioMapper := mux.NewIOMapperImpl(
		muxWrapper,
	)

	// --- NEXT TAP ---
	inventoryService := inventory.NewServiceImpl(
		inventoryRepository,
		entityFactory,
		entityModifier,
		voFactory,
	)
	decoderService := json.NewDecoderServiceImpl()
	encoderService := json.NewEncoderServiceImpl()
	responseFactory := http.NewResponseFactoryImpl()
	parameterConverter := http.NewParameterConverterImpl()
	handlerMapper := mux.NewHandlerMapperImpl(
		ioMapper,
	)

	// --- NEXT TAP ---
	inventoryController := http.NewInventoryControllerImpl(
		inventoryService,
		decoderService,
		encoderService,
		responseFactory,
		parameterConverter,
	)
	serverConfiguration := mux.NewServerConfigurationImpl(
		configStore,
		handlerMapper,
		muxWrapper,
	)

	// --- NEXT TAP ---
	return http.NewServerFactoryImpl(
		inventoryController,
		serverConfiguration,
	), nil
}
