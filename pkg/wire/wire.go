package wire

import (
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
func CreateApp() domain.Runnable {
	factory, err := CreateServerFactory()
	if err != nil {
		return func() error {
			return err
		}
	}
	return factory.Create()
}

// CreateServerFactory injects all the dependencies needed to create
// http.ServerFactory
func CreateServerFactory() (http.ServerFactory, error) {
	// Each "tap" below indicates a level of dependency
	databaseService, err := db.NewDatabaseServiceImpl()
	if err != nil {
		return nil, err
	}
	inventoryItemConstructor := entity.NewInventoryItemConstructorImpl()
	muxWrapper := mux.NewWrapperImpl()

	// --- NEXT TAP ---
	inventoryRepository := sql.NewInventoryRepositoryImpl(
		databaseService,
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
		handlerMapper,
		muxWrapper,
	)

	// --- NEXT TAP ---
	return http.NewServerFactoryImpl(
		inventoryController,
		serverConfiguration,
	), nil
}
